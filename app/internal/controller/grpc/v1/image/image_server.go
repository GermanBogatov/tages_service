package image

import (
	"context"
	"fmt"
	pb_prod_images "github.com/GermanBogatov/tages_contracts/gen/go/tages_service/images/v1"
	"github.com/GermanBogatov/tages_service/app/pkg/logging"
)

func (s *Server) CreateImage(ctx context.Context, req *pb_prod_images.CreateImageRequest) (*pb_prod_images.CreateImageResponse, error) {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"1": "1",
	})
	logger.Info("create image from image_server in controller")
	fmt.Println("create image from image_server in controller")
	err := s.policy.Create(ctx)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (s *Server) ViewAllImage(ctx context.Context, req *pb_prod_images.ViewAllImageRequest) (*pb_prod_images.ViewAllImageResponse, error) {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"1": "1",
	})
	logger.Info("VIEW image from image_server in controller")
	fmt.Println("VIEW image from image_server in controller")
	err := s.policy.All(ctx)
	if err != nil {
		return nil, err
	}
	return nil, err
}
