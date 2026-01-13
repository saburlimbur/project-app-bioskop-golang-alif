package usecase

import (
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/pkg/utils"

	"go.uber.org/zap"
)

type Service struct {
	UserService   UserServiceCase
	CinemaService CinemasServiceCase
}

func NewService(
	repo repository.Repository,
	log *zap.Logger,
	conf utils.Configuration,
) Service {
	return Service{
		UserService:   NewUserServiceCase(repo, log, conf),
		CinemaService: NewCinemaServiceCase(repo, log, conf),
	}
}
