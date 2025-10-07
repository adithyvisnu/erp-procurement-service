package postgre

import (
	"context"
	"log"

	"github.com/adithyvisnu/erp-procurement-service/internal/model"
	"gorm.io/gorm"
)

func Migrate(ctx context.Context, db *gorm.DB) {
	if err := db.WithContext(ctx).AutoMigrate(
		&model.Supplier{},
		&model.Item{},
		&model.PurchaseRequest{},
		&model.PurchaseRequestItem{},
	); err != nil {
		log.Fatalf("migrate: %v", err)
	}
}
