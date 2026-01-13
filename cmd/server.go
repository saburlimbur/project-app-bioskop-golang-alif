package cmd

import (
	"alfdwirhmn/bioskop/pkg/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func ApiServer(
	config utils.Configuration,
	logger *zap.Logger,
	h *chi.Mux,
) {
	fmt.Println("Server running on port 3000")
	if err := http.ListenAndServe(":3000", h); err != nil {
		logger.Fatal("can't run service", zap.Error(err))
	}
}
