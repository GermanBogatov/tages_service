package image

import (
	pb_prod_images "github.com/GermanBogatov/tages_contracts/gen/go/tages_service/images/v1"
	"github.com/GermanBogatov/tages_service/app/internal/domain/image/policy"
)

type Server struct {
	policy *policy.ImagePolicy
	pb_prod_images.UnimplementedTagesServiceServer
}

func NewServer(
	policy *policy.ImagePolicy,
	srv pb_prod_images.UnimplementedTagesServiceServer,
) *Server {
	return &Server{
		policy:                          policy,
		UnimplementedTagesServiceServer: srv,
	}
}
