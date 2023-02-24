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
	productionFields              = builder.RawFieldNames(&Production{}, true)
	productionRows                = strings.Join(productionFields, ",")
	productionRowsExpectAutoSet   = strings.Join(stringx.Remove(productionFields, "id", "create_time", "update_time", "create_at", "update_at"), ",")
	productionRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(productionFields, "id", "create_time", "update_time", "create_at", "update_at"))

	cacheProductionIdFmt = "cache:%s:production:id:"
)

type (
	productionModel interface {
		Columns(ctx context.Context) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error

		Insert(ctx context.Context, data *Production) (int64, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Production) (int64, error)

		Update(ctx context.Context, data *Production) error
		TransUpdate(ctx context.Context, session sqlx.Session, data *Production) error

		FindOne(ctx context.Context, id int64) (*Production, error)
		TransFindOne(ctx context.Context, session sqlx.Session, id int64) (*Production, error)

		Delete(ctx context.Context, id int64) error
	}

	defaultProductionModel struct {
		sqlc.CachedConn
		table string
	}

	Production struct {
		Id        int64  `db:"id" json:"id"`
		Category  string `db:"category" json:"category"`
		Name      string `db:"name" json:"name"`
		Status    int64  `db:"status" json:"status"`
		CreatedAt int64  `db:"created_at" json:"created_at"`
		UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	}
)

func newProductionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductionModel {
	return &defaultProductionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"production"`,
	}
}

func (m *defaultProductionModel) tableName(schema string) string {
	fmt.Println(schema)
	return fmt.Sprintf("%s.%s", schema, m.table)
}

func (q *Production) SetCreatedAtAndUpdatedAt() {
	if q == nil {
		return
	}
	q.CreatedAt = time.Now().Unix()
	q.SetUpdatedAt()
}

func (q *Production) SetUpdatedAt() {
	if q == nil {
		return
	}
	q.UpdatedAt = time.Now().Unix()
}

func (m *defaultProductionModel) Columns(ctx context.Context) squirrel.SelectBuilder {
	fmt.Println(ctxSup.GetSchema(ctx))
	return squirrel.Select(productionRows).From(m.tableName(ctxSup.GetSchema(ctx))).PlaceholderFormat(squirrel.Dollar)
}

func (m *defaultProductionModel) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.CachedConn.TransactCtx(
		ctx,
		func(ctx context.Context, session sqlx.Session) error {
			return fn(ctx, session)
		},
	)
}

func (m *defaultProductionModel) Insert(ctx context.Context, data *Production) (int64, error) {
	schema := ctxSup.GetSchema(ctx)

	var id int64 = 0
	data.SetCreatedAtAndUpdatedAt()

	productionIdKey := fmt.Sprintf("%s%v", cacheProductionIdFmt, data.Id)

	productionIdKey = dbSup.SchInit(productionIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5) RETURNING id", m.tableName(schema), productionRowsExpectAutoSet)

		return nil, conn.QueryRowCtx(ctx, &id, query, data.Category, data.Name, data.Status, data.CreatedAt, data.UpdatedAt)
	}, productionIdKey)

	return id, errors.Wrap(err, "failed to insert Production")
}

func (m *defaultProductionModel) TransInsert(ctx context.Context, session sqlx.Session, data *Production) (int64, error) {
	schema := ctxSup.GetSchema(ctx)

	var id int64 = 0
	data.SetCreatedAtAndUpdatedAt()

	productionIdKey := fmt.Sprintf("%s%v", cacheProductionIdFmt, data.Id)

	productionIdKey = dbSup.SchInit(productionIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5) RETURNING id", m.tableName(schema), productionRowsExpectAutoSet)

		return nil, session.QueryRowCtx(ctx, &id, query, data.Category, data.Name, data.Status, data.CreatedAt, data.UpdatedAt)
	}, productionIdKey)

	return id, errors.Wrap(err, "failed to insert Production")
}

func (m *defaultProductionModel) Update(ctx context.Context, data *Production) error {
	schema := ctxSup.GetSchema(ctx)

	data.SetUpdatedAt()

	productionIdKey := fmt.Sprintf("%s%v", cacheProductionIdFmt, data.Id)

	productionIdKey = dbSup.SchInit(productionIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.tableName(schema), productionRowsWithPlaceHolder)

		return conn.ExecCtx(ctx, query, data.Id, data.Category, data.Name, data.Status, data.CreatedAt, data.UpdatedAt)
	}, productionIdKey)

	return errors.Wrap(err, "failed to update Production")
}

func (m *defaultProductionModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Production) error {
	schema := ctxSup.GetSchema(ctx)

	data.SetUpdatedAt()

	productionIdKey := fmt.Sprintf("%s%v", cacheProductionIdFmt, data.Id)

	productionIdKey = dbSup.SchInit(productionIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.tableName(schema), productionRowsWithPlaceHolder)

		return session.ExecCtx(ctx, query, data.Id, data.Category, data.Name, data.Status, data.CreatedAt, data.UpdatedAt)
	}, productionIdKey)

	return errors.Wrap(err, "failed to update Production")
}

func (m *defaultProductionModel) FindOne(ctx context.Context, id int64) (*Production, error) {
	schema := ctxSup.GetSchema(ctx)

	var resp Production

	toSQL, args, err := m.Columns(ctx).Where(
		squirrel.Eq{"id": id},
	).Limit(1).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "toSQL construction error")
	}

	productionIdKey := fmt.Sprintf("%s%v", cacheProductionIdFmt, id)
	productionIdKey = dbSup.SchInit(productionIdKey, schema)

	err = m.QueryRowCtx(ctx, &resp, productionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, v, toSQL, args...)
	})
	switch err {
	case nil:
		return &resp, nil
	case ErrNotFound:
		return nil, errors.Wrap(err, fmt.Sprintf("Production not found by id: %v", id))
	default:
		return nil, errors.Wrap(err, "failed to find Production")
	}
}

func (m *defaultProductionModel) TransFindOne(ctx context.Context, session sqlx.Session, id int64) (*Production, error) {
	var resp Production

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
		return nil, errors.Wrap(err, fmt.Sprintf("Production not found by id: %v", id))
	default:
		return nil, errors.Wrap(err, "failed to find Production")
	}
}

func (m *defaultProductionModel) Delete(ctx context.Context, id int64) error {
	schema := ctxSup.GetSchema(ctx)

	productionIdKey := fmt.Sprintf("%s%v", cacheProductionIdFmt, id)

	productionIdKey = dbSup.SchInit(productionIdKey, schema)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.tableName(schema))
		return conn.ExecCtx(ctx, query, id)
	}, productionIdKey)

	return errors.Wrap(err, "failed to delete Production")
}

func (m *defaultProductionModel) byPrimaryKey(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	toSQL, args, err := m.Columns(ctx).Where(
		squirrel.Eq{"id": primary},
	).Limit(1).ToSql()
	if err != nil {
		return errors.Wrap(err, "toSQL construction error")
	}

	return conn.QueryRowCtx(ctx, v, toSQL, args...)
}
