package service

import (
	"context"
	"github.com/GermanBogatov/tages_service/app/pkg/errors"
)

type repository interface {
	All(context.Context) error
	Create(context.Context) error
}

type ImageService struct {
	repository repository
}

func NewImageService(repository repository) *ImageService {
	return &ImageService{repository: repository}
}

func (s *ImageService) All(ctx context.Context) error {
	err := s.repository.All(ctx)
	if err != nil {
		return errors.Wrap(err, "repository.All")
	}
	return nil
}

func (s *ImageService) Create(ctx context.Context) error {

	err := s.repository.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}
