package repository

import (
	"context"
	"github.com/yutify/architecture-pattern-2/domain/model"
)

type ChargeRepository interface {
	CreateChargeLog(ctx context.Context, detail *model.ChargeDetail) error
}
