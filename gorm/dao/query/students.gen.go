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

	"gorm/model"
)

func newStudents(db *gorm.DB, opts ...gen.DOOption) students {
	_students := students{}

	_students.studentsDo.UseDB(db, opts...)
	_students.studentsDo.UseModel(&model.Students{})

	tableName := _students.studentsDo.TableName()
	_students.ALL = field.NewAsterisk(tableName)
	_students.ID = field.NewUint(tableName, "id")
	_students.CreatedAt = field.NewTime(tableName, "created_at")
	_students.UpdatedAt = field.NewTime(tableName, "updated_at")
	_students.DeletedAt = field.NewField(tableName, "deleted_at")
	_students.Name = field.NewString(tableName, "name")
	_students.Age = field.NewUint8(tableName, "age")
	_students.Course = field.NewField(tableName, "course")

	_students.fillFieldMap()

	return _students
}

type students struct {
	studentsDo studentsDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Age       field.Uint8
	Course    field.Field

	fieldMap map[string]field.Expr
}

func (s students) Table(newTableName string) *students {
	s.studentsDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s students) As(alias string) *students {
	s.studentsDo.DO = *(s.studentsDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *students) updateTableName(table string) *students {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Name = field.NewString(table, "name")
	s.Age = field.NewUint8(table, "age")
	s.Course = field.NewField(table, "course")

	s.fillFieldMap()

	return s
}

func (s *students) WithContext(ctx context.Context) IStudentsDo { return s.studentsDo.WithContext(ctx) }

func (s students) TableName() string { return s.studentsDo.TableName() }

func (s students) Alias() string { return s.studentsDo.Alias() }

func (s students) Columns(cols ...field.Expr) gen.Columns { return s.studentsDo.Columns(cols...) }

func (s *students) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *students) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["age"] = s.Age
	s.fieldMap["course"] = s.Course
}

func (s students) clone(db *gorm.DB) students {
	s.studentsDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s students) replaceDB(db *gorm.DB) students {
	s.studentsDo.ReplaceDB(db)
	return s
}

type studentsDo struct{ gen.DO }

type IStudentsDo interface {
	gen.SubQuery
	Debug() IStudentsDo
	WithContext(ctx context.Context) IStudentsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStudentsDo
	WriteDB() IStudentsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStudentsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStudentsDo
	Not(conds ...gen.Condition) IStudentsDo
	Or(conds ...gen.Condition) IStudentsDo
	Select(conds ...field.Expr) IStudentsDo
	Where(conds ...gen.Condition) IStudentsDo
	Order(conds ...field.Expr) IStudentsDo
	Distinct(cols ...field.Expr) IStudentsDo
	Omit(cols ...field.Expr) IStudentsDo
	Join(table schema.Tabler, on ...field.Expr) IStudentsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStudentsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStudentsDo
	Group(cols ...field.Expr) IStudentsDo
	Having(conds ...gen.Condition) IStudentsDo
	Limit(limit int) IStudentsDo
	Offset(offset int) IStudentsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentsDo
	Unscoped() IStudentsDo
	Create(values ...*model.Students) error
	CreateInBatches(values []*model.Students, batchSize int) error
	Save(values ...*model.Students) error
	First() (*model.Students, error)
	Take() (*model.Students, error)
	Last() (*model.Students, error)
	Find() ([]*model.Students, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Students, err error)
	FindInBatches(result *[]*model.Students, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Students) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStudentsDo
	Assign(attrs ...field.AssignExpr) IStudentsDo
	Joins(fields ...field.RelationField) IStudentsDo
	Preload(fields ...field.RelationField) IStudentsDo
	FirstOrInit() (*model.Students, error)
	FirstOrCreate() (*model.Students, error)
	FindByPage(offset int, limit int) (result []*model.Students, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStudentsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s studentsDo) Debug() IStudentsDo {
	return s.withDO(s.DO.Debug())
}

func (s studentsDo) WithContext(ctx context.Context) IStudentsDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studentsDo) ReadDB() IStudentsDo {
	return s.Clauses(dbresolver.Read)
}

func (s studentsDo) WriteDB() IStudentsDo {
	return s.Clauses(dbresolver.Write)
}

func (s studentsDo) Session(config *gorm.Session) IStudentsDo {
	return s.withDO(s.DO.Session(config))
}

func (s studentsDo) Clauses(conds ...clause.Expression) IStudentsDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studentsDo) Returning(value interface{}, columns ...string) IStudentsDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studentsDo) Not(conds ...gen.Condition) IStudentsDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s studentsDo) Or(conds ...gen.Condition) IStudentsDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s studentsDo) Select(conds ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s studentsDo) Where(conds ...gen.Condition) IStudentsDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s studentsDo) Order(conds ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s studentsDo) Distinct(cols ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studentsDo) Omit(cols ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studentsDo) Join(table schema.Tabler, on ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studentsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studentsDo) RightJoin(table schema.Tabler, on ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studentsDo) Group(cols ...field.Expr) IStudentsDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s studentsDo) Having(conds ...gen.Condition) IStudentsDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s studentsDo) Limit(limit int) IStudentsDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s studentsDo) Offset(offset int) IStudentsDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s studentsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentsDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studentsDo) Unscoped() IStudentsDo {
	return s.withDO(s.DO.Unscoped())
}

func (s studentsDo) Create(values ...*model.Students) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studentsDo) CreateInBatches(values []*model.Students, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studentsDo) Save(values ...*model.Students) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studentsDo) First() (*model.Students, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Students), nil
	}
}

func (s studentsDo) Take() (*model.Students, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Students), nil
	}
}

func (s studentsDo) Last() (*model.Students, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Students), nil
	}
}

func (s studentsDo) Find() ([]*model.Students, error) {
	result, err := s.DO.Find()
	return result.([]*model.Students), err
}

func (s studentsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Students, err error) {
	buf := make([]*model.Students, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studentsDo) FindInBatches(result *[]*model.Students, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studentsDo) Attrs(attrs ...field.AssignExpr) IStudentsDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studentsDo) Assign(attrs ...field.AssignExpr) IStudentsDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studentsDo) Joins(fields ...field.RelationField) IStudentsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studentsDo) Preload(fields ...field.RelationField) IStudentsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studentsDo) FirstOrInit() (*model.Students, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Students), nil
	}
}

func (s studentsDo) FirstOrCreate() (*model.Students, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Students), nil
	}
}

func (s studentsDo) FindByPage(offset int, limit int) (result []*model.Students, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s studentsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studentsDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studentsDo) Delete(models ...*model.Students) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studentsDo) withDO(do gen.Dao) *studentsDo {
	s.DO = *do.(*gen.DO)
	return s
}