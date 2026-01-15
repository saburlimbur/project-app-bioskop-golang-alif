package adaptor

import (
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/internal/usecase"
	"alfdwirhmn/bioskop/pkg/middleware"
	"alfdwirhmn/bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type PaymentAdaptorHandler struct {
	Payment   usecase.PaymentServiceCase
	Logger    *zap.Logger
	Validator *validator.Validate
	Config    utils.Configuration
}

func NewPaymentdaptorHandler(payment usecase.PaymentServiceCase, logg *zap.Logger, validator *validator.Validate, conf utils.Configuration,
) PaymentAdaptorHandler {
	return PaymentAdaptorHandler{
		Payment:   payment,
		Logger:    logg,
		Validator: validator,
		Config:    conf,
	}
}

func (h *PaymentAdaptorHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := r.Context().Value(middleware.ContextUserID).(int)
	if !ok {
		utils.JSONError(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	var req dto.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "validation failed", err)
		return
	}

	payment, err := h.Payment.CreatePayment(ctx, userID, req)
	if err != nil {
		h.Logger.Error("failed to create payment", zap.Error(err))
		utils.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONSuccess(w, http.StatusCreated, "Payment created succesfully", payment)
}

func (h *PaymentAdaptorHandler) ListPaymentMethods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	methods, err := h.Payment.ListPaymentMethods(ctx)
	if err != nil {
		h.Logger.Error("failed to list payment methods", zap.Error(err))
		utils.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONSuccess(
		w,
		http.StatusOK,
		"successfully get payment methods",
		dto.ToPaymentMethodResponses(methods),
	)
}
