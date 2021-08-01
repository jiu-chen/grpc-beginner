package services

import (
	context "context"
)

type ProdService struct {
}

func (ps *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch in.ProdArea {
	case ProdArea_A:
		stock = 31
	case ProdArea_B:
		stock = 32
	case ProdArea_C:
		stock = 33
	default:
		stock = -1
	}

	return &ProdResponse{ProdStock: stock}, nil
}

func (ps *ProdService) GetProdStocks(ctx context.Context, in *QuerySize) (*ProdResponseList, error) {
	ProdRes := []*ProdResponse{
		{ProdStock: 1},
		{ProdStock: 2},
		{ProdStock: 3},
		{ProdStock: 4},
	}
	return &ProdResponseList{Prodres: ProdRes}, nil
}

func (ps *ProdService) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: 1, ProdName: "phone", ProdPrice: 3999}, nil
}
