package repository

import (
	"context"

	"gorm.io/gorm"
)

type TransactionRepositoryI interface {
	BeginTransaction(ctx context.Context) *gorm.DB
	RollBackTransaction(ctx context.Context, trx *gorm.DB) *gorm.DB
	CommitTransaction(ctx context.Context, trx *gorm.DB) *gorm.DB
}

type TransactionRepository struct {
	Conn *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryI {
	return TransactionRepository{Conn: db}
}

func (re TransactionRepository) BeginTransaction(ctx context.Context) *gorm.DB {
	return re.Conn.WithContext(ctx).Begin()
}
func (re TransactionRepository) RollBackTransaction(ctx context.Context, trx *gorm.DB) *gorm.DB {
	return trx.Rollback()
}
func (re TransactionRepository) CommitTransaction(ctx context.Context, trx *gorm.DB) *gorm.DB {
	return trx.Commit()
}
