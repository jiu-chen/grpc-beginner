package services

import (
	context "context"
)

type ProdService struct {
}

func (ps *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 13}, nil
}
