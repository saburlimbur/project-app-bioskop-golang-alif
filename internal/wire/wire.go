package wire

import (
	"alfdwirhmn/bioskop/internal/adaptor"
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/internal/usecase"
	"alfdwirhmn/bioskop/pkg/database"
	appMiddleware "alfdwirhmn/bioskop/pkg/middleware"
	"alfdwirhmn/bioskop/pkg/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Response struct {
	Message string `json:"message"`
}

func SetupRouter(
	db *pgxpool.Pool,
	logger *zap.Logger,
	config utils.Configuration,
) *chi.Mux {
	r := chi.NewRouter()

	// global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// container
	repo := repository.NewRepository(db, logger)
	txMgr := database.NewTxManager(db)
	validate := validator.New()
	service := usecase.NewService(repo, txMgr, logger, config)

	userHandler := adaptor.NewUserAdaptorHandler(
		service.UserService,
		logger,
		config,
	)

	cinemaHandler := adaptor.NewCinemadaptorHandler(
		service.CinemaService,
		logger,
		config,
	)

	bookingHandler := adaptor.NewBookingAdaptorHandler(
		service.BookingService,
		validate,
		logger,
		config,
	)

	paymentHandler := adaptor.NewPaymentdaptorHandler(
		service.PaymentService,
		logger,
		validate,
		config,
	)

	// health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.JSONSuccess(w, http.StatusOK, "OK", Response{
			Message: "Bioskop Cinema Booking API Running",
		})
	})

	// router
	r.Route("/api/v1", func(api chi.Router) {

		api.Route("/users", func(r chi.Router) {
			r.Post("/register", userHandler.Register)
			r.Post("/login", userHandler.Login)
			r.Post("/logout", userHandler.Logout)
			r.Post("/verify-email", userHandler.VerifyEmail)
		})

		api.Route("/user", func(r chi.Router) {
			r.With(appMiddleware.AuthMiddleware(repo)).
				Get("/bookings", bookingHandler.GetUserBookingHistory)
		})

		api.Route("/cinemas", func(r chi.Router) {
			r.With(appMiddleware.AuthMiddleware(repo)).Get("/", cinemaHandler.Lists)
			r.With(appMiddleware.AuthMiddleware(repo)).Get("/{id}", cinemaHandler.DetailById)
			r.With(appMiddleware.AuthMiddleware(repo)).Get("/{id}/seats", cinemaHandler.SeatAvailability)
		})

		api.Route("/booking", func(r chi.Router) {
			r.With(appMiddleware.AuthMiddleware(repo)).
				Post("/", bookingHandler.CreateBooking)
			r.With(appMiddleware.AuthMiddleware(repo)).
				Get("/{id}", bookingHandler.GetBookingByID)
		})

		api.Route("/payment", func(r chi.Router) {
			r.With(appMiddleware.AuthMiddleware(repo)).
				Post("/", paymentHandler.CreatePayment)
		})

		api.Route("/payment-methods", func(r chi.Router) {
			r.With(appMiddleware.AuthMiddleware(repo)).
				Get("/", paymentHandler.ListPaymentMethods)
		})
	})

	return r
}
