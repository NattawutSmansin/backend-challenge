package handler

import (
	domain "backend-challenge/module/thirdChallenge/domain"
	proroBuff "backend-challenge/module/thirdChallenge/proto"
	"context"
)

type Handler struct {
	Usecase domain.UseCase
	proroBuff.UnimplementedPieFireDireServicesServer
}

func NewHandler(Usecase domain.UseCase) *Handler {
	return &Handler{
		Usecase: Usecase,
	}
}

func (h *Handler) GetPieFireDire(ctx context.Context, request *proroBuff.Request) (*proroBuff.Response, error) {
	fetchBeefData, err := h.Usecase.GetPieFireDire()
	if err != nil {
		return nil, err
	}

	return &proroBuff.Response{Beef: fetchBeefData}, nil
}
