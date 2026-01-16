package adaptor

import (
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/internal/usecase"
	"alfdwirhmn/bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type UserAdaptorHandler struct {
	User   usecase.UserServiceCase
	Logger *zap.Logger
	Config utils.Configuration
}

func NewUserAdaptorHandler(
	user usecase.UserServiceCase,
	logg *zap.Logger,
	conf utils.Configuration,
) UserAdaptorHandler {
	return UserAdaptorHandler{
		User:   user,
		Logger: logg,
		Config: conf,
	}
}

func (h *UserAdaptorHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	// validate
	if validationErrors, err := utils.ValidateErrors(req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "validation failed", validationErrors)
		return
	}

	ctx := r.Context()

	user, err := h.User.Register(ctx, req)
	if err != nil {
		h.Logger.Error("failed to register user", zap.Error(err))
		utils.JSONError(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// mapping
	resp := dto.ToUserResponse(user)

	utils.JSONSuccess(w, http.StatusCreated, "Register successfully, Please verify your email", resp)
}

func (h *UserAdaptorHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	// validate
	if validationErrors, err := utils.ValidateErrors(req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "validation failed", validationErrors)
		return
	}

	ctx := r.Context()

	// IP & device info
	ip := r.RemoteAddr
	deviceInfo := r.UserAgent()

	resp, err := h.User.Login(ctx, req, ip, deviceInfo)
	if err != nil {
		h.Logger.Warn("login failed",
			zap.String("identifier", req.Identifier),
			zap.Error(err),
		)
		utils.JSONError(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	utils.JSONSuccess(w, http.StatusOK, "Login successfully", resp)
}

func (h *UserAdaptorHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.JSONError(w, http.StatusUnauthorized, "missing authorization header", nil)
		return
	}

	const prefix = "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		utils.JSONError(w, http.StatusUnauthorized, "invalid authorization format", nil)
		return
	}

	token := authHeader[len(prefix):]

	err := h.User.Logout(r.Context(), token)
	if err != nil {
		h.Logger.Warn("logout failed", zap.Error(err))
		utils.JSONError(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	utils.JSONSuccess(w, http.StatusOK, "Logout successfully", nil)
}

func (h *UserAdaptorHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyEmailRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	if validationErrors, err := utils.ValidateErrors(req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "validation failed", validationErrors)
		return
	}

	user, err := h.User.VerifyEmail(r.Context(), req)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONSuccess(w, http.StatusOK, "Email verified successfully", user)
}
