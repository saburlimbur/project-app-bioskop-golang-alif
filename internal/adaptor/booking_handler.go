package adaptor

import (
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/internal/usecase"
	"alfdwirhmn/bioskop/pkg/middleware"
	"alfdwirhmn/bioskop/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BookingAdapterHandler struct {
	Booking   usecase.BookingServiceCase
	Logger    *zap.Logger
	Validator *validator.Validate
	Config    utils.Configuration
}

func NewBookingAdaptorHandler(booking usecase.BookingServiceCase, validator *validator.Validate, logg *zap.Logger, conf utils.Configuration) BookingAdapterHandler {
	return BookingAdapterHandler{
		Booking:   booking,
		Validator: validator,
		Logger:    logg,
		Config:    conf,
	}
}

func (h *BookingAdapterHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// user from context
	userID, ok := r.Context().Value(middleware.ContextUserID).(int)
	if !ok {
		utils.JSONError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	var req dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Validation failed", err)
		return
	}

	booking, err := h.Booking.CreateBooking(ctx, userID, req)
	if err != nil {
		h.Logger.Error("failed to create booking", zap.Error(err))
		utils.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONSuccess(w, http.StatusCreated, "Booking created successfully", booking)
}

func (h *BookingAdapterHandler) GetBookingByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	// auth
	userID, ok := ctx.Value(middleware.ContextUserID).(int)
	if !ok {
		utils.JSONError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	// path param
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		utils.JSONError(w, http.StatusBadRequest, "booking id is required", nil)
		return
	}

	bookingID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid booking id", nil)
		return
	}

	booking, err := h.Booking.FindByID(ctx, bookingID)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "booking not found", nil)
		return
	}

	if booking.UserID != userID {
		utils.JSONError(w, http.StatusForbidden, "access denied", nil)
		return
	}

	// mapping
	resp := dto.ToBookingResponse(booking)

	utils.JSONSuccess(
		w,
		http.StatusOK,
		"successfully get booking detail",
		resp,
	)
}
