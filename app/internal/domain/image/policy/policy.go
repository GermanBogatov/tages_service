package policy

import (
	"context"
	"github.com/GermanBogatov/tages_service/app/internal/domain/image/service"
	"github.com/GermanBogatov/tages_service/app/pkg/errors"
)

type ImagePolicy struct {
	imageService *service.ImageService
}

func NewImagePolicy(imageService *service.ImageService) *ImagePolicy {
	return &ImagePolicy{imageService: imageService}
}

func (p *ImagePolicy) All(ctx context.Context) error {
	err := p.imageService.All(ctx)
	if err != nil {
		return errors.Wrap(err, "productService.All")
	}

	return nil
}

func (p *ImagePolicy) Create(ctx context.Context) error {
	return p.imageService.Create(ctx)
}
