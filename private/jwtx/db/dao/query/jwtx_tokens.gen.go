// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/5-say/zero-services/private/jwtx/db/dao/model"
)

func newJwtxToken(db *gorm.DB, opts ...gen.DOOption) jwtxToken {
	_jwtxToken := jwtxToken{}

	_jwtxToken.jwtxTokenDo.UseDB(db, opts...)
	_jwtxToken.jwtxTokenDo.UseModel(&model.JwtxToken{})

	tableName := _jwtxToken.jwtxTokenDo.TableName()
	_jwtxToken.ALL = field.NewAsterisk(tableName)
	_jwtxToken.ID = field.NewUint64(tableName, "id")
	_jwtxToken.TokenGroup = field.NewString(tableName, "token_group")
	_jwtxToken.AccountID = field.NewUint64(tableName, "account_id")
	_jwtxToken.MakeTokenIP = field.NewString(tableName, "make_token_ip")
	_jwtxToken.CreatedAt = field.NewTime(tableName, "created_at")
	_jwtxToken.LastRefreshAt = field.NewTime(tableName, "last_refresh_at")
	_jwtxToken.FinalRefreshAt = field.NewTime(tableName, "final_refresh_at")
	_jwtxToken.ExpirationAt = field.NewTime(tableName, "expiration_at")

	_jwtxToken.fillFieldMap()

	return _jwtxToken
}

type jwtxToken struct {
	jwtxTokenDo

	ALL            field.Asterisk
	ID             field.Uint64 // token ID
	TokenGroup     field.String // token 分组
	AccountID      field.Uint64 // 账户 ID
	MakeTokenIP    field.String // 首次请求生成 token 的 IP 地址
	CreatedAt      field.Time   // 创建时间
	LastRefreshAt  field.Time   // 上次的刷新时间
	FinalRefreshAt field.Time   // 最后的刷新时间
	ExpirationAt   field.Time   // 过期时间

	fieldMap map[string]field.Expr
}

func (j jwtxToken) Table(newTableName string) *jwtxToken {
	j.jwtxTokenDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jwtxToken) As(alias string) *jwtxToken {
	j.jwtxTokenDo.DO = *(j.jwtxTokenDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jwtxToken) updateTableName(table string) *jwtxToken {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewUint64(table, "id")
	j.TokenGroup = field.NewString(table, "token_group")
	j.AccountID = field.NewUint64(table, "account_id")
	j.MakeTokenIP = field.NewString(table, "make_token_ip")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.LastRefreshAt = field.NewTime(table, "last_refresh_at")
	j.FinalRefreshAt = field.NewTime(table, "final_refresh_at")
	j.ExpirationAt = field.NewTime(table, "expiration_at")

	j.fillFieldMap()

	return j
}

func (j *jwtxToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jwtxToken) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 8)
	j.fieldMap["id"] = j.ID
	j.fieldMap["token_group"] = j.TokenGroup
	j.fieldMap["account_id"] = j.AccountID
	j.fieldMap["make_token_ip"] = j.MakeTokenIP
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["last_refresh_at"] = j.LastRefreshAt
	j.fieldMap["final_refresh_at"] = j.FinalRefreshAt
	j.fieldMap["expiration_at"] = j.ExpirationAt
}

func (j jwtxToken) clone(db *gorm.DB) jwtxToken {
	j.jwtxTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jwtxToken) replaceDB(db *gorm.DB) jwtxToken {
	j.jwtxTokenDo.ReplaceDB(db)
	return j
}

type jwtxTokenDo struct{ gen.DO }

type IJwtxTokenDo interface {
	gen.SubQuery
	Debug() IJwtxTokenDo
	WithContext(ctx context.Context) IJwtxTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJwtxTokenDo
	WriteDB() IJwtxTokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJwtxTokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJwtxTokenDo
	Not(conds ...gen.Condition) IJwtxTokenDo
	Or(conds ...gen.Condition) IJwtxTokenDo
	Select(conds ...field.Expr) IJwtxTokenDo
	Where(conds ...gen.Condition) IJwtxTokenDo
	Order(conds ...field.Expr) IJwtxTokenDo
	Distinct(cols ...field.Expr) IJwtxTokenDo
	Omit(cols ...field.Expr) IJwtxTokenDo
	Join(table schema.Tabler, on ...field.Expr) IJwtxTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJwtxTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJwtxTokenDo
	Group(cols ...field.Expr) IJwtxTokenDo
	Having(conds ...gen.Condition) IJwtxTokenDo
	Limit(limit int) IJwtxTokenDo
	Offset(offset int) IJwtxTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJwtxTokenDo
	Unscoped() IJwtxTokenDo
	Create(values ...*model.JwtxToken) error
	CreateInBatches(values []*model.JwtxToken, batchSize int) error
	Save(values ...*model.JwtxToken) error
	First() (*model.JwtxToken, error)
	Take() (*model.JwtxToken, error)
	Last() (*model.JwtxToken, error)
	Find() ([]*model.JwtxToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JwtxToken, err error)
	FindInBatches(result *[]*model.JwtxToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.JwtxToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJwtxTokenDo
	Assign(attrs ...field.AssignExpr) IJwtxTokenDo
	Joins(fields ...field.RelationField) IJwtxTokenDo
	Preload(fields ...field.RelationField) IJwtxTokenDo
	FirstOrInit() (*model.JwtxToken, error)
	FirstOrCreate() (*model.JwtxToken, error)
	FindByPage(offset int, limit int) (result []*model.JwtxToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJwtxTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jwtxTokenDo) Debug() IJwtxTokenDo {
	return j.withDO(j.DO.Debug())
}

func (j jwtxTokenDo) WithContext(ctx context.Context) IJwtxTokenDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jwtxTokenDo) ReadDB() IJwtxTokenDo {
	return j.Clauses(dbresolver.Read)
}

func (j jwtxTokenDo) WriteDB() IJwtxTokenDo {
	return j.Clauses(dbresolver.Write)
}

func (j jwtxTokenDo) Session(config *gorm.Session) IJwtxTokenDo {
	return j.withDO(j.DO.Session(config))
}

func (j jwtxTokenDo) Clauses(conds ...clause.Expression) IJwtxTokenDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jwtxTokenDo) Returning(value interface{}, columns ...string) IJwtxTokenDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jwtxTokenDo) Not(conds ...gen.Condition) IJwtxTokenDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jwtxTokenDo) Or(conds ...gen.Condition) IJwtxTokenDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jwtxTokenDo) Select(conds ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jwtxTokenDo) Where(conds ...gen.Condition) IJwtxTokenDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jwtxTokenDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IJwtxTokenDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j jwtxTokenDo) Order(conds ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jwtxTokenDo) Distinct(cols ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jwtxTokenDo) Omit(cols ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jwtxTokenDo) Join(table schema.Tabler, on ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jwtxTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jwtxTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jwtxTokenDo) Group(cols ...field.Expr) IJwtxTokenDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jwtxTokenDo) Having(conds ...gen.Condition) IJwtxTokenDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jwtxTokenDo) Limit(limit int) IJwtxTokenDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jwtxTokenDo) Offset(offset int) IJwtxTokenDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jwtxTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJwtxTokenDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jwtxTokenDo) Unscoped() IJwtxTokenDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jwtxTokenDo) Create(values ...*model.JwtxToken) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jwtxTokenDo) CreateInBatches(values []*model.JwtxToken, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jwtxTokenDo) Save(values ...*model.JwtxToken) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jwtxTokenDo) First() (*model.JwtxToken, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtxToken), nil
	}
}

func (j jwtxTokenDo) Take() (*model.JwtxToken, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtxToken), nil
	}
}

func (j jwtxTokenDo) Last() (*model.JwtxToken, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtxToken), nil
	}
}

func (j jwtxTokenDo) Find() ([]*model.JwtxToken, error) {
	result, err := j.DO.Find()
	return result.([]*model.JwtxToken), err
}

func (j jwtxTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JwtxToken, err error) {
	buf := make([]*model.JwtxToken, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jwtxTokenDo) FindInBatches(result *[]*model.JwtxToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jwtxTokenDo) Attrs(attrs ...field.AssignExpr) IJwtxTokenDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jwtxTokenDo) Assign(attrs ...field.AssignExpr) IJwtxTokenDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jwtxTokenDo) Joins(fields ...field.RelationField) IJwtxTokenDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jwtxTokenDo) Preload(fields ...field.RelationField) IJwtxTokenDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jwtxTokenDo) FirstOrInit() (*model.JwtxToken, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtxToken), nil
	}
}

func (j jwtxTokenDo) FirstOrCreate() (*model.JwtxToken, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JwtxToken), nil
	}
}

func (j jwtxTokenDo) FindByPage(offset int, limit int) (result []*model.JwtxToken, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jwtxTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jwtxTokenDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jwtxTokenDo) Delete(models ...*model.JwtxToken) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jwtxTokenDo) withDO(do gen.Dao) *jwtxTokenDo {
	j.DO = *do.(*gen.DO)
	return j
}
