package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CustomerModel = (*customCustomerModel)(nil)

type (
	// CustomerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerModel.
	CustomerModel interface {
		customerModel
		FindByEmail(ctx context.Context, email string) (*Customer, error)
	}

	customCustomerModel struct {
		*defaultCustomerModel
	}
)

// NewCustomerModel returns a model for the database table.
func NewCustomerModel(conn sqlx.SqlConn, c cache.CacheConf) CustomerModel {
	return &customCustomerModel{
		defaultCustomerModel: newCustomerModel(conn, c),
	}
}

// ByEmail
func (m *defaultCustomerModel) FindByEmail(ctx context.Context, email string) (*Customer, error) {
	trytryCustomerIdKey := fmt.Sprintf("%s%v", cacheTrytryCustomerIdPrefix, email)
	var resp Customer
	err := m.QueryRowCtx(ctx, &resp, trytryCustomerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", customerRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, email)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
