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

func newRolePermission(db *gorm.DB, opts ...gen.DOOption) rolePermission {
	_rolePermission := rolePermission{}

	_rolePermission.rolePermissionDo.UseDB(db, opts...)
	_rolePermission.rolePermissionDo.UseModel(&model.RolePermission{})

	tableName := _rolePermission.rolePermissionDo.TableName()
	_rolePermission.ALL = field.NewAsterisk(tableName)
	_rolePermission.ID = field.NewInt64(tableName, "id")
	_rolePermission.RoleID = field.NewString(tableName, "role_id")
	_rolePermission.PermissionID = field.NewInt64(tableName, "permission_id")
	_rolePermission.CreatedAt = field.NewTime(tableName, "created_at")
	_rolePermission.UpdatedAt = field.NewTime(tableName, "updated_at")
	_rolePermission.DeletedAt = field.NewField(tableName, "deleted_at")

	_rolePermission.fillFieldMap()

	return _rolePermission
}

type rolePermission struct {
	rolePermissionDo rolePermissionDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键id
	RoleID       field.String // 角色id
	PermissionID field.Int64  // 权限id
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (r rolePermission) Table(newTableName string) *rolePermission {
	r.rolePermissionDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rolePermission) As(alias string) *rolePermission {
	r.rolePermissionDo.DO = *(r.rolePermissionDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rolePermission) updateTableName(table string) *rolePermission {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.RoleID = field.NewString(table, "role_id")
	r.PermissionID = field.NewInt64(table, "permission_id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *rolePermission) WithContext(ctx context.Context) *rolePermissionDo {
	return r.rolePermissionDo.WithContext(ctx)
}

func (r rolePermission) TableName() string { return r.rolePermissionDo.TableName() }

func (r rolePermission) Alias() string { return r.rolePermissionDo.Alias() }

func (r *rolePermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rolePermission) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["id"] = r.ID
	r.fieldMap["role_id"] = r.RoleID
	r.fieldMap["permission_id"] = r.PermissionID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r rolePermission) clone(db *gorm.DB) rolePermission {
	r.rolePermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r rolePermission) replaceDB(db *gorm.DB) rolePermission {
	r.rolePermissionDo.ReplaceDB(db)
	return r
}

type rolePermissionDo struct{ gen.DO }

func (r rolePermissionDo) Debug() *rolePermissionDo {
	return r.withDO(r.DO.Debug())
}

func (r rolePermissionDo) WithContext(ctx context.Context) *rolePermissionDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rolePermissionDo) ReadDB() *rolePermissionDo {
	return r.Clauses(dbresolver.Read)
}

func (r rolePermissionDo) WriteDB() *rolePermissionDo {
	return r.Clauses(dbresolver.Write)
}

func (r rolePermissionDo) Session(config *gorm.Session) *rolePermissionDo {
	return r.withDO(r.DO.Session(config))
}

func (r rolePermissionDo) Clauses(conds ...clause.Expression) *rolePermissionDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rolePermissionDo) Returning(value interface{}, columns ...string) *rolePermissionDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rolePermissionDo) Not(conds ...gen.Condition) *rolePermissionDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rolePermissionDo) Or(conds ...gen.Condition) *rolePermissionDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rolePermissionDo) Select(conds ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rolePermissionDo) Where(conds ...gen.Condition) *rolePermissionDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rolePermissionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *rolePermissionDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r rolePermissionDo) Order(conds ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rolePermissionDo) Distinct(cols ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rolePermissionDo) Omit(cols ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rolePermissionDo) Join(table schema.Tabler, on ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rolePermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rolePermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rolePermissionDo) Group(cols ...field.Expr) *rolePermissionDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rolePermissionDo) Having(conds ...gen.Condition) *rolePermissionDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rolePermissionDo) Limit(limit int) *rolePermissionDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rolePermissionDo) Offset(offset int) *rolePermissionDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rolePermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *rolePermissionDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rolePermissionDo) Unscoped() *rolePermissionDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rolePermissionDo) Create(values ...*model.RolePermission) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rolePermissionDo) CreateInBatches(values []*model.RolePermission, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rolePermissionDo) Save(values ...*model.RolePermission) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rolePermissionDo) First() (*model.RolePermission, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermission), nil
	}
}

func (r rolePermissionDo) Take() (*model.RolePermission, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermission), nil
	}
}

func (r rolePermissionDo) Last() (*model.RolePermission, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermission), nil
	}
}

func (r rolePermissionDo) Find() ([]*model.RolePermission, error) {
	result, err := r.DO.Find()
	return result.([]*model.RolePermission), err
}

func (r rolePermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RolePermission, err error) {
	buf := make([]*model.RolePermission, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rolePermissionDo) FindInBatches(result *[]*model.RolePermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rolePermissionDo) Attrs(attrs ...field.AssignExpr) *rolePermissionDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rolePermissionDo) Assign(attrs ...field.AssignExpr) *rolePermissionDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rolePermissionDo) Joins(fields ...field.RelationField) *rolePermissionDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rolePermissionDo) Preload(fields ...field.RelationField) *rolePermissionDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rolePermissionDo) FirstOrInit() (*model.RolePermission, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermission), nil
	}
}

func (r rolePermissionDo) FirstOrCreate() (*model.RolePermission, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RolePermission), nil
	}
}

func (r rolePermissionDo) FindByPage(offset int, limit int) (result []*model.RolePermission, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r rolePermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rolePermissionDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rolePermissionDo) Delete(models ...*model.RolePermission) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rolePermissionDo) withDO(do gen.Dao) *rolePermissionDo {
	r.DO = *do.(*gen.DO)
	return r
}
