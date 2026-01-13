package adaptor

import (
	"alfdwirhmn/bioskop/internal/dto"
	"alfdwirhmn/bioskop/internal/usecase"
	"alfdwirhmn/bioskop/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CinemaAdaptorHandler struct {
	Cinema usecase.CinemasServiceCase
	Logger *zap.Logger
	Config utils.Configuration
}

func NewCinemadaptorHandler(cinema usecase.CinemasServiceCase, logg *zap.Logger, conf utils.Configuration,
) CinemaAdaptorHandler {
	return CinemaAdaptorHandler{
		Cinema: cinema,
		Logger: logg,
		Config: conf,
	}
}

func (h *CinemaAdaptorHandler) Lists(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		utils.JSONError(w, http.StatusBadRequest, "Invalid page", nil)
		return
	}

	limit, err := strconv.Atoi(h.Config.Limit)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Invalid limit config", nil)
		return
	}

	cinema, pagination, err := h.Cinema.FindAll(r.Context(), page, limit)
	if err != nil {
		h.Logger.Error("failed get cinemas", zap.Error(err))
		utils.JSONError(w, http.StatusInternalServerError, "failed", nil)
		return
	}

	respPagination := &utils.PaginationResponse{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalRows:  pagination.TotalRows,
		TotalPages: pagination.TotalPages,
	}

	utils.JSONWithPagination(w, http.StatusOK, "success get data", cinema, respPagination)
}

func (h *CinemaAdaptorHandler) DetailById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	cinema, err := h.Cinema.FindById(r.Context(), id)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "cinema not found", nil)
		return
	}

	utils.JSONSuccess(w, http.StatusOK, "success", dto.ToCinemaResponse(cinema))
}

func (h *CinemaAdaptorHandler) SeatAvailability(w http.ResponseWriter, r *http.Request) {
	cinemaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid cinema id", nil)
		return
	}

	date := r.URL.Query().Get("date")
	time := r.URL.Query().Get("time")

	if date == "" || time == "" {
		utils.JSONError(w, http.StatusBadRequest, "date and time are required", nil)
		return
	}

	data, err := h.Cinema.SeatAvailability(r.Context(), cinemaID, date, time)
	if err != nil {
		h.Logger.Error("failed get seat availability", zap.Error(err))
		utils.JSONError(w, http.StatusInternalServerError, "failed get seats", nil)
		return
	}

	utils.JSONSuccess(w, http.StatusOK, "success get seat availability", data)
}
