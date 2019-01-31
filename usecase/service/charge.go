package service

import (
	"context"
	"fmt"
	"github.com/omise/omise-go/operations"
	"github.com/yutify/architecture-pattern-2/domain/model"
	"github.com/yutify/architecture-pattern-2/usecase/repository"
)

type ChargeService interface {
	ChargeMoney(ctx context.Context, chargeInfo *operations.CreateCharge, userID string) error
}

type chargeService struct {
	ChargeRepo repository.ChargeRepository
	UserRepo   repository.UserRepository
}

func NewChargeService(cr repository.ChargeRepository, ur repository.UserRepository) ChargeService {
	return &chargeService{
		ChargeRepo: cr,
		UserRepo:   ur,
	}
}

func (cr *chargeService) ChargeMoney(ctx context.Context, chargeInfo *operations.CreateCharge, userID string) error {
	cd := &model.ChargeDetail{
		UserID:   userID,
		Amount:   chargeInfo.Amount,
		Currency: chargeInfo.Currency,
	}
	ok, err := cr.UserRepo.UserExists(ctx, userID)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Errorf("user does not exist")
	}
	if err := cr.ChargeRepo.CreateChargeLog(ctx, cd); err != nil {
		return err
	}
	return nil
}
