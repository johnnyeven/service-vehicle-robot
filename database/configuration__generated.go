package database

import (
	"fmt"
	"time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var ConfigurationTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	ConfigurationTable = DBRobot.Register(&Configuration{})
}

func (configuration *Configuration) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBRobot
}

func (configuration *Configuration) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return ConfigurationTable
}

func (configuration *Configuration) TableName() string {
	return "t_configuration"
}

type ConfigurationFields struct {
	ID              *github_com_johnnyeven_libtools_sqlx_builder.Column
	ConfigurationID *github_com_johnnyeven_libtools_sqlx_builder.Column
	StackID         *github_com_johnnyeven_libtools_sqlx_builder.Column
	Key             *github_com_johnnyeven_libtools_sqlx_builder.Column
	Value           *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime      *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime      *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled         *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var ConfigurationField = struct {
	ID              string
	ConfigurationID string
	StackID         string
	Key             string
	Value           string
	CreateTime      string
	UpdateTime      string
	Enabled         string
}{
	ID:              "ID",
	ConfigurationID: "ConfigurationID",
	StackID:         "StackID",
	Key:             "Key",
	Value:           "Value",
	CreateTime:      "CreateTime",
	UpdateTime:      "UpdateTime",
	Enabled:         "Enabled",
}

func (configuration *Configuration) Fields() *ConfigurationFields {
	table := configuration.T()

	return &ConfigurationFields{
		ID:              table.F(ConfigurationField.ID),
		ConfigurationID: table.F(ConfigurationField.ConfigurationID),
		StackID:         table.F(ConfigurationField.StackID),
		Key:             table.F(ConfigurationField.Key),
		Value:           table.F(ConfigurationField.Value),
		CreateTime:      table.F(ConfigurationField.CreateTime),
		UpdateTime:      table.F(ConfigurationField.UpdateTime),
		Enabled:         table.F(ConfigurationField.Enabled),
	}
}

func (configuration *Configuration) IndexFieldNames() []string {
	return []string{"ConfigurationID", "ID", "Key", "StackID"}
}

func (configuration *Configuration) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := configuration.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(configuration)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range configuration.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	return condition
}

func (configuration *Configuration) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (configuration *Configuration) Indexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{"I_stack": github_com_johnnyeven_libtools_sqlx.FieldNames{"StackID"}}
}
func (configuration *Configuration) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{
		"U_configuration_id": github_com_johnnyeven_libtools_sqlx.FieldNames{"ConfigurationID", "Enabled"},
		"U_stack_key":        github_com_johnnyeven_libtools_sqlx.FieldNames{"StackID", "Key", "Enabled"},
	}
}
func (configuration *Configuration) Comments() map[string]string {
	return map[string]string{
		"ConfigurationID": "业务ID",
		"CreateTime":      "",
		"Enabled":         "",
		"ID":              "",
		"Key":             "Key",
		"StackID":         "StackID",
		"UpdateTime":      "",
		"Value":           "Value",
	}
}

func (configuration *Configuration) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if configuration.CreateTime.IsZero() {
		configuration.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	configuration.UpdateTime = configuration.CreateTime

	stmt := configuration.D().
		Insert(configuration).
		Comment("Configuration.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		configuration.ID = uint64(lastInsertID)
	}

	return err
}

func (configuration *Configuration) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := configuration.T()

	stmt := table.Delete().
		Comment("Configuration.DeleteByStruct").
		Where(configuration.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (configuration *Configuration) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if configuration.CreateTime.IsZero() {
		configuration.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	configuration.UpdateTime = configuration.CreateTime

	table := configuration.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(configuration, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range configuration.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(m, field)
		}
	}

	if len(m) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !m[field] {
			delete(fieldValues, field)
		}
	}

	stmt := table.
		Insert().Columns(cols).Values(vals...).
		OnDuplicateKeyUpdate(table.AssignsByFieldValues(fieldValues)...).
		Comment("Configuration.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (configuration *Configuration) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Select().
		Comment("Configuration.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(configuration.ID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Select().
		Comment("Configuration.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(configuration.ID),
			table.F("Enabled").Eq(configuration.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Delete().
		Comment("Configuration.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(configuration.ID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Configuration.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(configuration.ID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	dbRet := db.Do(stmt).Scan(configuration)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return configuration.FetchByID(db)
	}
	return nil
}

func (configuration *Configuration) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(configuration, zeroFields...)
	return configuration.UpdateByIDWithMap(db, fieldValues)
}

func (configuration *Configuration) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Configuration.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(configuration.ID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	dbRet := db.Do(stmt).Scan(configuration)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return configuration.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (configuration *Configuration) FetchByConfigurationID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Select().
		Comment("Configuration.FetchByConfigurationID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ConfigurationID").Eq(configuration.ConfigurationID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) FetchByConfigurationIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Select().
		Comment("Configuration.FetchByConfigurationIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ConfigurationID").Eq(configuration.ConfigurationID),
			table.F("Enabled").Eq(configuration.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) DeleteByConfigurationID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Delete().
		Comment("Configuration.DeleteByConfigurationID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ConfigurationID").Eq(configuration.ConfigurationID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) UpdateByConfigurationIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Configuration.UpdateByConfigurationIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ConfigurationID").Eq(configuration.ConfigurationID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	dbRet := db.Do(stmt).Scan(configuration)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return configuration.FetchByConfigurationID(db)
	}
	return nil
}

func (configuration *Configuration) UpdateByConfigurationIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(configuration, zeroFields...)
	return configuration.UpdateByConfigurationIDWithMap(db, fieldValues)
}

func (configuration *Configuration) SoftDeleteByConfigurationID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Configuration.SoftDeleteByConfigurationID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ConfigurationID").Eq(configuration.ConfigurationID),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	dbRet := db.Do(stmt).Scan(configuration)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return configuration.DeleteByConfigurationID(db)
		}
		return err
	}
	return nil
}

func (configuration *Configuration) FetchByStackIDAndKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Select().
		Comment("Configuration.FetchByStackIDAndKey").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("StackID").Eq(configuration.StackID),
			table.F("Key").Eq(configuration.Key),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) FetchByStackIDAndKeyForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Select().
		Comment("Configuration.FetchByStackIDAndKeyForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("StackID").Eq(configuration.StackID),
			table.F("Key").Eq(configuration.Key),
			table.F("Enabled").Eq(configuration.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) DeleteByStackIDAndKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()
	stmt := table.Delete().
		Comment("Configuration.DeleteByStackIDAndKey").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("StackID").Eq(configuration.StackID),
			table.F("Key").Eq(configuration.Key),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	return db.Do(stmt).Scan(configuration).Err()
}

func (configuration *Configuration) UpdateByStackIDAndKeyWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Configuration.UpdateByStackIDAndKeyWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("StackID").Eq(configuration.StackID),
			table.F("Key").Eq(configuration.Key),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	dbRet := db.Do(stmt).Scan(configuration)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return configuration.FetchByStackIDAndKey(db)
	}
	return nil
}

func (configuration *Configuration) UpdateByStackIDAndKeyWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(configuration, zeroFields...)
	return configuration.UpdateByStackIDAndKeyWithMap(db, fieldValues)
}

func (configuration *Configuration) SoftDeleteByStackIDAndKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	configuration.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := configuration.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Configuration.SoftDeleteByStackIDAndKey").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("StackID").Eq(configuration.StackID),
			table.F("Key").Eq(configuration.Key),
			table.F("Enabled").Eq(configuration.Enabled),
		))

	dbRet := db.Do(stmt).Scan(configuration)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return configuration.DeleteByStackIDAndKey(db)
		}
		return err
	}
	return nil
}

type ConfigurationList []Configuration

// deprecated
func (configurationList *ConfigurationList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*configurationList, count, err = (&Configuration{}).FetchList(db, size, offset, conditions...)
	return
}

func (configuration *Configuration) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (configurationList ConfigurationList, count int32, err error) {
	configurationList = ConfigurationList{}

	table := configuration.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}

func (configuration *Configuration) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (configurationList ConfigurationList, err error) {
	configurationList = ConfigurationList{}

	table := configuration.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.List").
		Where(condition)

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}

func (configuration *Configuration) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (configurationList ConfigurationList, err error) {
	configurationList = ConfigurationList{}

	table := configuration.T()

	condition := configuration.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}

// deprecated
func (configurationList *ConfigurationList) BatchFetchByConfigurationIDList(db *github_com_johnnyeven_libtools_sqlx.DB, configurationIDList []uint64) (err error) {
	*configurationList, err = (&Configuration{}).BatchFetchByConfigurationIDList(db, configurationIDList)
	return
}

func (configuration *Configuration) BatchFetchByConfigurationIDList(db *github_com_johnnyeven_libtools_sqlx.DB, configurationIDList []uint64) (configurationList ConfigurationList, err error) {
	if len(configurationIDList) == 0 {
		return ConfigurationList{}, nil
	}

	table := configuration.T()

	condition := table.F("ConfigurationID").In(configurationIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.BatchFetchByConfigurationIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}

// deprecated
func (configurationList *ConfigurationList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*configurationList, err = (&Configuration{}).BatchFetchByIDList(db, idList)
	return
}

func (configuration *Configuration) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (configurationList ConfigurationList, err error) {
	if len(idList) == 0 {
		return ConfigurationList{}, nil
	}

	table := configuration.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}

// deprecated
func (configurationList *ConfigurationList) BatchFetchByKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, keyList []string) (err error) {
	*configurationList, err = (&Configuration{}).BatchFetchByKeyList(db, keyList)
	return
}

func (configuration *Configuration) BatchFetchByKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, keyList []string) (configurationList ConfigurationList, err error) {
	if len(keyList) == 0 {
		return ConfigurationList{}, nil
	}

	table := configuration.T()

	condition := table.F("Key").In(keyList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.BatchFetchByKeyList").
		Where(condition)

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}

// deprecated
func (configurationList *ConfigurationList) BatchFetchByStackIDList(db *github_com_johnnyeven_libtools_sqlx.DB, stackIDList []uint64) (err error) {
	*configurationList, err = (&Configuration{}).BatchFetchByStackIDList(db, stackIDList)
	return
}

func (configuration *Configuration) BatchFetchByStackIDList(db *github_com_johnnyeven_libtools_sqlx.DB, stackIDList []uint64) (configurationList ConfigurationList, err error) {
	if len(stackIDList) == 0 {
		return ConfigurationList{}, nil
	}

	table := configuration.T()

	condition := table.F("StackID").In(stackIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Configuration.BatchFetchByStackIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&configurationList).Err()

	return
}
