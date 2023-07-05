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

	"permission-open/internal/pkg/dal/model"
)

func newThirdNoticeGroup(db *gorm.DB, opts ...gen.DOOption) thirdNoticeGroup {
	_thirdNoticeGroup := thirdNoticeGroup{}

	_thirdNoticeGroup.thirdNoticeGroupDo.UseDB(db, opts...)
	_thirdNoticeGroup.thirdNoticeGroupDo.UseModel(&model.ThirdNoticeGroup{})

	tableName := _thirdNoticeGroup.thirdNoticeGroupDo.TableName()
	_thirdNoticeGroup.ALL = field.NewAsterisk(tableName)
	_thirdNoticeGroup.ID = field.NewInt64(tableName, "id")
	_thirdNoticeGroup.GroupID = field.NewString(tableName, "group_id")
	_thirdNoticeGroup.Name = field.NewString(tableName, "name")
	_thirdNoticeGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_thirdNoticeGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_thirdNoticeGroup.DeletedAt = field.NewField(tableName, "deleted_at")

	_thirdNoticeGroup.fillFieldMap()

	return _thirdNoticeGroup
}

type thirdNoticeGroup struct {
	thirdNoticeGroupDo thirdNoticeGroupDo

	ALL       field.Asterisk
	ID        field.Int64  // 主键id
	GroupID   field.String // 通知组id
	Name      field.String // 通知组名称
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (t thirdNoticeGroup) Table(newTableName string) *thirdNoticeGroup {
	t.thirdNoticeGroupDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t thirdNoticeGroup) As(alias string) *thirdNoticeGroup {
	t.thirdNoticeGroupDo.DO = *(t.thirdNoticeGroupDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *thirdNoticeGroup) updateTableName(table string) *thirdNoticeGroup {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.GroupID = field.NewString(table, "group_id")
	t.Name = field.NewString(table, "name")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *thirdNoticeGroup) WithContext(ctx context.Context) *thirdNoticeGroupDo {
	return t.thirdNoticeGroupDo.WithContext(ctx)
}

func (t thirdNoticeGroup) TableName() string { return t.thirdNoticeGroupDo.TableName() }

func (t thirdNoticeGroup) Alias() string { return t.thirdNoticeGroupDo.Alias() }

func (t *thirdNoticeGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *thirdNoticeGroup) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["group_id"] = t.GroupID
	t.fieldMap["name"] = t.Name
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t thirdNoticeGroup) clone(db *gorm.DB) thirdNoticeGroup {
	t.thirdNoticeGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t thirdNoticeGroup) replaceDB(db *gorm.DB) thirdNoticeGroup {
	t.thirdNoticeGroupDo.ReplaceDB(db)
	return t
}

type thirdNoticeGroupDo struct{ gen.DO }

func (t thirdNoticeGroupDo) Debug() *thirdNoticeGroupDo {
	return t.withDO(t.DO.Debug())
}

func (t thirdNoticeGroupDo) WithContext(ctx context.Context) *thirdNoticeGroupDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t thirdNoticeGroupDo) ReadDB() *thirdNoticeGroupDo {
	return t.Clauses(dbresolver.Read)
}

func (t thirdNoticeGroupDo) WriteDB() *thirdNoticeGroupDo {
	return t.Clauses(dbresolver.Write)
}

func (t thirdNoticeGroupDo) Session(config *gorm.Session) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Session(config))
}

func (t thirdNoticeGroupDo) Clauses(conds ...clause.Expression) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t thirdNoticeGroupDo) Returning(value interface{}, columns ...string) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t thirdNoticeGroupDo) Not(conds ...gen.Condition) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t thirdNoticeGroupDo) Or(conds ...gen.Condition) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t thirdNoticeGroupDo) Select(conds ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t thirdNoticeGroupDo) Where(conds ...gen.Condition) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t thirdNoticeGroupDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *thirdNoticeGroupDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t thirdNoticeGroupDo) Order(conds ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t thirdNoticeGroupDo) Distinct(cols ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t thirdNoticeGroupDo) Omit(cols ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t thirdNoticeGroupDo) Join(table schema.Tabler, on ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t thirdNoticeGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t thirdNoticeGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t thirdNoticeGroupDo) Group(cols ...field.Expr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t thirdNoticeGroupDo) Having(conds ...gen.Condition) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t thirdNoticeGroupDo) Limit(limit int) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t thirdNoticeGroupDo) Offset(offset int) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t thirdNoticeGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t thirdNoticeGroupDo) Unscoped() *thirdNoticeGroupDo {
	return t.withDO(t.DO.Unscoped())
}

func (t thirdNoticeGroupDo) Create(values ...*model.ThirdNoticeGroup) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t thirdNoticeGroupDo) CreateInBatches(values []*model.ThirdNoticeGroup, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t thirdNoticeGroupDo) Save(values ...*model.ThirdNoticeGroup) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t thirdNoticeGroupDo) First() (*model.ThirdNoticeGroup, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ThirdNoticeGroup), nil
	}
}

func (t thirdNoticeGroupDo) Take() (*model.ThirdNoticeGroup, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ThirdNoticeGroup), nil
	}
}

func (t thirdNoticeGroupDo) Last() (*model.ThirdNoticeGroup, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ThirdNoticeGroup), nil
	}
}

func (t thirdNoticeGroupDo) Find() ([]*model.ThirdNoticeGroup, error) {
	result, err := t.DO.Find()
	return result.([]*model.ThirdNoticeGroup), err
}

func (t thirdNoticeGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ThirdNoticeGroup, err error) {
	buf := make([]*model.ThirdNoticeGroup, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t thirdNoticeGroupDo) FindInBatches(result *[]*model.ThirdNoticeGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t thirdNoticeGroupDo) Attrs(attrs ...field.AssignExpr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t thirdNoticeGroupDo) Assign(attrs ...field.AssignExpr) *thirdNoticeGroupDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t thirdNoticeGroupDo) Joins(fields ...field.RelationField) *thirdNoticeGroupDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t thirdNoticeGroupDo) Preload(fields ...field.RelationField) *thirdNoticeGroupDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t thirdNoticeGroupDo) FirstOrInit() (*model.ThirdNoticeGroup, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ThirdNoticeGroup), nil
	}
}

func (t thirdNoticeGroupDo) FirstOrCreate() (*model.ThirdNoticeGroup, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ThirdNoticeGroup), nil
	}
}

func (t thirdNoticeGroupDo) FindByPage(offset int, limit int) (result []*model.ThirdNoticeGroup, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t thirdNoticeGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t thirdNoticeGroupDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t thirdNoticeGroupDo) Delete(models ...*model.ThirdNoticeGroup) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *thirdNoticeGroupDo) withDO(do gen.Dao) *thirdNoticeGroupDo {
	t.DO = *do.(*gen.DO)
	return t
}
