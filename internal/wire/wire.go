package wire

import (
	"alfdwirhmn/bioskop/internal/adaptor"
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/internal/usecase"
	appMiddleware "alfdwirhmn/bioskop/pkg/middleware"
	"alfdwirhmn/bioskop/pkg/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

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
	service := usecase.NewService(repo, logger, config)

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
		})

		api.Route("/cinemas", func(r chi.Router) {
			r.With(appMiddleware.AuthMiddleware(repo)).Get("/", cinemaHandler.Lists)
			r.With(appMiddleware.AuthMiddleware(repo)).Get("/{id}", cinemaHandler.DetailById)
			r.With(appMiddleware.AuthMiddleware(repo)).Get("/{id}/seats", cinemaHandler.SeatAvailability)
		})
	})

	return r
}
