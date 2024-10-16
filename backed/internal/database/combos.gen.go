// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package database

import (
	"backed/app/models"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCombo(db *gorm.DB, opts ...gen.DOOption) combo {
	_combo := combo{}

	_combo.comboDo.UseDB(db, opts...)
	_combo.comboDo.UseModel(&models.Combo{})

	tableName := _combo.comboDo.TableName()
	_combo.ALL = field.NewAsterisk(tableName)
	_combo.ID = field.NewUint(tableName, "id")
	_combo.CreatedAt = field.NewTime(tableName, "created_at")
	_combo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_combo.DeletedAt = field.NewField(tableName, "deleted_at")
	_combo.Name = field.NewString(tableName, "name")
	_combo.Price = field.NewInt(tableName, "price")
	_combo.StartYear = field.NewInt(tableName, "start_year")
	_combo.EndYear = field.NewInt(tableName, "end_year")
	_combo.MonthFee = field.NewInt(tableName, "month_fee")
	_combo.HighMedicalCoverage = field.NewBool(tableName, "high_medical_coverage")
	_combo.RefundPeriod = field.NewInt(tableName, "refund_period")
	_combo.Description = comboHasManyDescription{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Description", "models.ComboDescription"),
	}

	_combo.fillFieldMap()

	return _combo
}

type combo struct {
	comboDo

	ALL                 field.Asterisk
	ID                  field.Uint
	CreatedAt           field.Time
	UpdatedAt           field.Time
	DeletedAt           field.Field
	Name                field.String
	Price               field.Int
	StartYear           field.Int
	EndYear             field.Int
	MonthFee            field.Int
	HighMedicalCoverage field.Bool
	RefundPeriod        field.Int
	Description         comboHasManyDescription

	fieldMap map[string]field.Expr
}

func (c combo) Table(newTableName string) *combo {
	c.comboDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c combo) As(alias string) *combo {
	c.comboDo.DO = *(c.comboDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *combo) updateTableName(table string) *combo {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Name = field.NewString(table, "name")
	c.Price = field.NewInt(table, "price")
	c.StartYear = field.NewInt(table, "start_year")
	c.EndYear = field.NewInt(table, "end_year")
	c.MonthFee = field.NewInt(table, "month_fee")
	c.HighMedicalCoverage = field.NewBool(table, "high_medical_coverage")
	c.RefundPeriod = field.NewInt(table, "refund_period")

	c.fillFieldMap()

	return c
}

func (c *combo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *combo) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["name"] = c.Name
	c.fieldMap["price"] = c.Price
	c.fieldMap["start_year"] = c.StartYear
	c.fieldMap["end_year"] = c.EndYear
	c.fieldMap["month_fee"] = c.MonthFee
	c.fieldMap["high_medical_coverage"] = c.HighMedicalCoverage
	c.fieldMap["refund_period"] = c.RefundPeriod

}

func (c combo) clone(db *gorm.DB) combo {
	c.comboDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c combo) replaceDB(db *gorm.DB) combo {
	c.comboDo.ReplaceDB(db)
	return c
}

type comboHasManyDescription struct {
	db *gorm.DB

	field.RelationField
}

func (a comboHasManyDescription) Where(conds ...field.Expr) *comboHasManyDescription {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a comboHasManyDescription) WithContext(ctx context.Context) *comboHasManyDescription {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a comboHasManyDescription) Session(session *gorm.Session) *comboHasManyDescription {
	a.db = a.db.Session(session)
	return &a
}

func (a comboHasManyDescription) Model(m *models.Combo) *comboHasManyDescriptionTx {
	return &comboHasManyDescriptionTx{a.db.Model(m).Association(a.Name())}
}

type comboHasManyDescriptionTx struct{ tx *gorm.Association }

func (a comboHasManyDescriptionTx) Find() (result []*models.ComboDescription, err error) {
	return result, a.tx.Find(&result)
}

func (a comboHasManyDescriptionTx) Append(values ...*models.ComboDescription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a comboHasManyDescriptionTx) Replace(values ...*models.ComboDescription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a comboHasManyDescriptionTx) Delete(values ...*models.ComboDescription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a comboHasManyDescriptionTx) Clear() error {
	return a.tx.Clear()
}

func (a comboHasManyDescriptionTx) Count() int64 {
	return a.tx.Count()
}

type comboDo struct{ gen.DO }

type IComboDo interface {
	gen.SubQuery
	Debug() IComboDo
	WithContext(ctx context.Context) IComboDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IComboDo
	WriteDB() IComboDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IComboDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IComboDo
	Not(conds ...gen.Condition) IComboDo
	Or(conds ...gen.Condition) IComboDo
	Select(conds ...field.Expr) IComboDo
	Where(conds ...gen.Condition) IComboDo
	Order(conds ...field.Expr) IComboDo
	Distinct(cols ...field.Expr) IComboDo
	Omit(cols ...field.Expr) IComboDo
	Join(table schema.Tabler, on ...field.Expr) IComboDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IComboDo
	RightJoin(table schema.Tabler, on ...field.Expr) IComboDo
	Group(cols ...field.Expr) IComboDo
	Having(conds ...gen.Condition) IComboDo
	Limit(limit int) IComboDo
	Offset(offset int) IComboDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IComboDo
	Unscoped() IComboDo
	Create(values ...*models.Combo) error
	CreateInBatches(values []*models.Combo, batchSize int) error
	Save(values ...*models.Combo) error
	First() (*models.Combo, error)
	Take() (*models.Combo, error)
	Last() (*models.Combo, error)
	Find() ([]*models.Combo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Combo, err error)
	FindInBatches(result *[]*models.Combo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Combo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IComboDo
	Assign(attrs ...field.AssignExpr) IComboDo
	Joins(fields ...field.RelationField) IComboDo
	Preload(fields ...field.RelationField) IComboDo
	FirstOrInit() (*models.Combo, error)
	FirstOrCreate() (*models.Combo, error)
	FindByPage(offset int, limit int) (result []*models.Combo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IComboDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c comboDo) Debug() IComboDo {
	return c.withDO(c.DO.Debug())
}

func (c comboDo) WithContext(ctx context.Context) IComboDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c comboDo) ReadDB() IComboDo {
	return c.Clauses(dbresolver.Read)
}

func (c comboDo) WriteDB() IComboDo {
	return c.Clauses(dbresolver.Write)
}

func (c comboDo) Session(config *gorm.Session) IComboDo {
	return c.withDO(c.DO.Session(config))
}

func (c comboDo) Clauses(conds ...clause.Expression) IComboDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c comboDo) Returning(value interface{}, columns ...string) IComboDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c comboDo) Not(conds ...gen.Condition) IComboDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c comboDo) Or(conds ...gen.Condition) IComboDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c comboDo) Select(conds ...field.Expr) IComboDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c comboDo) Where(conds ...gen.Condition) IComboDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c comboDo) Order(conds ...field.Expr) IComboDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c comboDo) Distinct(cols ...field.Expr) IComboDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c comboDo) Omit(cols ...field.Expr) IComboDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c comboDo) Join(table schema.Tabler, on ...field.Expr) IComboDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c comboDo) LeftJoin(table schema.Tabler, on ...field.Expr) IComboDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c comboDo) RightJoin(table schema.Tabler, on ...field.Expr) IComboDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c comboDo) Group(cols ...field.Expr) IComboDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c comboDo) Having(conds ...gen.Condition) IComboDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c comboDo) Limit(limit int) IComboDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c comboDo) Offset(offset int) IComboDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c comboDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IComboDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c comboDo) Unscoped() IComboDo {
	return c.withDO(c.DO.Unscoped())
}

func (c comboDo) Create(values ...*models.Combo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c comboDo) CreateInBatches(values []*models.Combo, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c comboDo) Save(values ...*models.Combo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c comboDo) First() (*models.Combo, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Combo), nil
	}
}

func (c comboDo) Take() (*models.Combo, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Combo), nil
	}
}

func (c comboDo) Last() (*models.Combo, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Combo), nil
	}
}

func (c comboDo) Find() ([]*models.Combo, error) {
	result, err := c.DO.Find()
	return result.([]*models.Combo), err
}

func (c comboDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Combo, err error) {
	buf := make([]*models.Combo, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c comboDo) FindInBatches(result *[]*models.Combo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c comboDo) Attrs(attrs ...field.AssignExpr) IComboDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c comboDo) Assign(attrs ...field.AssignExpr) IComboDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c comboDo) Joins(fields ...field.RelationField) IComboDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c comboDo) Preload(fields ...field.RelationField) IComboDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c comboDo) FirstOrInit() (*models.Combo, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Combo), nil
	}
}

func (c comboDo) FirstOrCreate() (*models.Combo, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Combo), nil
	}
}

func (c comboDo) FindByPage(offset int, limit int) (result []*models.Combo, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c comboDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c comboDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c comboDo) Delete(models ...*models.Combo) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *comboDo) withDO(do gen.Dao) *comboDo {
	c.DO = *do.(*gen.DO)
	return c
}
