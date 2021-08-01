package services

import context "context"

type OrderService struct {
}

func (os *OrderService) NewOrder(context.Context, *OrderMain) (*OrderResponse, error) {
	return &OrderResponse{Status: "ok", Message: "success"}, nil
}
