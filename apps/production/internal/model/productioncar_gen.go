// Generated code, DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"

	ctxSup "github.com/v3nooonn/trytry/pkg/utils/contextx"
	dbSup "github.com/v3nooonn/trytry/pkg/utils/dbx"
)

var (
	productionCarFields              = builder.RawFieldNames(&ProductionCar{}, true)
	productionCarRows                = strings.Join(productionCarFields, ",")
	productionCarRowsExpectAutoSet   = strings.Join(stringx.Remove(productionCarFields, "id", "create_time", "update_time", "create_at", "update_at"), ",")
	productionCarRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(productionCarFields, "id", "create_time", "update_time", "create_at", "update_at"))

	cacheProductionCarIdFmt = "cache:%s:production_car:id:"
)

type (
	productionCarModel interface {
		Columns(ctx context.Context) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error

		Insert(ctx context.Context, data *ProductionCar) (int64, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *ProductionCar) (int64, error)

		Update(ctx context.Context, data *ProductionCar) error
		TransUpdate(ctx context.Context, session sqlx.Session, data *ProductionCar) error

		FindOne(ctx context.Context, id int64) (*ProductionCar, error)
		TransFindOne(ctx context.Context, session sqlx.Session, id int64) (*ProductionCar, error)

		Delete(ctx context.Context, id int64) error
	}

	defaultProductionCarModel struct {
		sqlc.CachedConn
		table string
	}

	ProductionCar struct {
		Id        int64  `db:"id" json:"id"`
		ProdId    int64  `db:"prod_id" json:"prod_id"`
		Type      string `db:"type" json:"type"`
		Series    string `db:"series" json:"series"`
		Name      string `db:"name" json:"name"`
		CreatedAt int64  `db:"created_at" json:"created_at"`
		UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	}
)

func newProductionCarModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductionCarModel {
	return &defaultProductionCarModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"production_car"`,
	}
}

func (m *defaultProductionCarModel) tableName(schema string) string {
	return fmt.Sprintf("%s.%s", schema, m.table)
}

func (q *ProductionCar) SetCreatedAtAndUpdatedAt() {
	if q == nil {
		return
	}
	q.CreatedAt = time.Now().Unix()
	q.SetUpdatedAt()
}

func (q *ProductionCar) SetUpdatedAt() {
	if q == nil {
		return
	}
	q.UpdatedAt = time.Now().Unix()
}

func (m *defaultProductionCarModel) Columns(ctx context.Context) squirrel.SelectBuilder {
	return squirrel.Select(productionCarRows).From(m.tableName(ctxSup.GetSchema(ctx))).PlaceholderFormat(squirrel.Dollar)
}

func (m *defaultProductionCarModel) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.CachedConn.TransactCtx(
		ctx,
		func(ctx context.Context, session sqlx.Session) error {
			return fn(ctx, session)
		},
	)
}

func (m *defaultProductionCarModel) Insert(ctx context.Context, data *ProductionCar) (int64, error) {
	schema := ctxSup.GetSchema(ctx)

	var id int64 = 0
	data.SetCreatedAtAndUpdatedAt()

	productionCarIdKey := fmt.Sprintf("%s%v", cacheProductionCarIdFmt, data.Id)

	productionCarIdKey = dbSup.SchInit(productionCarIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6) RETURNING id", m.tableName(schema), productionCarRowsExpectAutoSet)

		return nil, conn.QueryRowCtx(ctx, &id, query, data.ProdId, data.Type, data.Series, data.Name, data.CreatedAt, data.UpdatedAt)
	}, productionCarIdKey)

	return id, errors.Wrap(err, "failed to insert ProductionCar")
}

func (m *defaultProductionCarModel) TransInsert(ctx context.Context, session sqlx.Session, data *ProductionCar) (int64, error) {
	schema := ctxSup.GetSchema(ctx)

	var id int64 = 0
	data.SetCreatedAtAndUpdatedAt()

	productionCarIdKey := fmt.Sprintf("%s%v", cacheProductionCarIdFmt, data.Id)

	productionCarIdKey = dbSup.SchInit(productionCarIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6) RETURNING id", m.tableName(schema), productionCarRowsExpectAutoSet)

		return nil, session.QueryRowCtx(ctx, &id, query, data.ProdId, data.Type, data.Series, data.Name, data.CreatedAt, data.UpdatedAt)
	}, productionCarIdKey)

	return id, errors.Wrap(err, "failed to insert ProductionCar")
}

func (m *defaultProductionCarModel) Update(ctx context.Context, data *ProductionCar) error {
	schema := ctxSup.GetSchema(ctx)

	data.SetUpdatedAt()

	productionCarIdKey := fmt.Sprintf("%s%v", cacheProductionCarIdFmt, data.Id)

	productionCarIdKey = dbSup.SchInit(productionCarIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.tableName(schema), productionCarRowsWithPlaceHolder)

		return conn.ExecCtx(ctx, query, data.Id, data.ProdId, data.Type, data.Series, data.Name, data.CreatedAt, data.UpdatedAt)
	}, productionCarIdKey)

	return errors.Wrap(err, "failed to update ProductionCar")
}

func (m *defaultProductionCarModel) TransUpdate(ctx context.Context, session sqlx.Session, data *ProductionCar) error {
	schema := ctxSup.GetSchema(ctx)

	data.SetUpdatedAt()

	productionCarIdKey := fmt.Sprintf("%s%v", cacheProductionCarIdFmt, data.Id)

	productionCarIdKey = dbSup.SchInit(productionCarIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.tableName(schema), productionCarRowsWithPlaceHolder)

		return session.ExecCtx(ctx, query, data.Id, data.ProdId, data.Type, data.Series, data.Name, data.CreatedAt, data.UpdatedAt)
	}, productionCarIdKey)

	return errors.Wrap(err, "failed to update ProductionCar")
}

func (m *defaultProductionCarModel) FindOne(ctx context.Context, id int64) (*ProductionCar, error) {
	schema := ctxSup.GetSchema(ctx)

	var resp ProductionCar

	toSQL, args, err := m.Columns(ctx).Where(
		squirrel.Eq{"id": id},
	).Limit(1).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "toSQL construction error")
	}

	productionCarIdKey := fmt.Sprintf("%s%v", cacheProductionCarIdFmt, id)
	productionCarIdKey = dbSup.SchInit(productionCarIdKey, schema)

	err = m.QueryRowCtx(ctx, &resp, productionCarIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, v, toSQL, args...)
	})
	switch err {
	case nil:
		return &resp, nil
	case ErrNotFound:
		return nil, errors.Wrap(err, fmt.Sprintf("ProductionCar not found by id: %v", id))
	default:
		return nil, errors.Wrap(err, "failed to find ProductionCar")
	}
}

func (m *defaultProductionCarModel) TransFindOne(ctx context.Context, session sqlx.Session, id int64) (*ProductionCar, error) {
	var resp ProductionCar

	toSQL, args, err := m.Columns(ctx).Where(
		squirrel.Eq{"id": id},
	).Limit(1).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "toSQL construction error")
	}

	err = session.QueryRowCtx(ctx, &resp, toSQL, args...)
	switch err {
	case nil:
		return &resp, nil
	case ErrNotFound:
		return nil, errors.Wrap(err, fmt.Sprintf("ProductionCar not found by id: %v", id))
	default:
		return nil, errors.Wrap(err, "failed to find ProductionCar")
	}
}

func (m *defaultProductionCarModel) Delete(ctx context.Context, id int64) error {
	schema := ctxSup.GetSchema(ctx)

	productionCarIdKey := fmt.Sprintf("%s%v", cacheProductionCarIdFmt, id)

	productionCarIdKey = dbSup.SchInit(productionCarIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.tableName(schema))
		return conn.ExecCtx(ctx, query, id)
	}, productionCarIdKey)

	return errors.Wrap(err, "failed to delete ProductionCar")
}

func (m *defaultProductionCarModel) byPrimaryKey(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	toSQL, args, err := m.Columns(ctx).Where(
		squirrel.Eq{"id": primary},
	).Limit(1).ToSql()
	if err != nil {
		return errors.Wrap(err, "toSQL construction error")
	}

	return conn.QueryRowCtx(ctx, v, toSQL, args...)
}
