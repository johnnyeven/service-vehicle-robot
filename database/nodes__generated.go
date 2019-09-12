package database

import (
	fmt "fmt"
	time "time"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
	github_com_johnnyeven_libtools_sqlx "github.com/johnnyeven/libtools/sqlx"
	github_com_johnnyeven_libtools_sqlx_builder "github.com/johnnyeven/libtools/sqlx/builder"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

var NodesTable *github_com_johnnyeven_libtools_sqlx_builder.Table

func init() {
	NodesTable = DBRobot.Register(&Nodes{})
}

func (nodes *Nodes) D() *github_com_johnnyeven_libtools_sqlx.Database {
	return DBRobot
}

func (nodes *Nodes) T() *github_com_johnnyeven_libtools_sqlx_builder.Table {
	return NodesTable
}

func (nodes *Nodes) TableName() string {
	return "t_nodes"
}

type NodesFields struct {
	ID         *github_com_johnnyeven_libtools_sqlx_builder.Column
	NodesID    *github_com_johnnyeven_libtools_sqlx_builder.Column
	Key        *github_com_johnnyeven_libtools_sqlx_builder.Column
	Secret     *github_com_johnnyeven_libtools_sqlx_builder.Column
	Comment    *github_com_johnnyeven_libtools_sqlx_builder.Column
	CreateTime *github_com_johnnyeven_libtools_sqlx_builder.Column
	UpdateTime *github_com_johnnyeven_libtools_sqlx_builder.Column
	Enabled    *github_com_johnnyeven_libtools_sqlx_builder.Column
}

var NodesField = struct {
	ID         string
	NodesID    string
	Key        string
	Secret     string
	Comment    string
	CreateTime string
	UpdateTime string
	Enabled    string
}{
	ID:         "ID",
	NodesID:    "NodesID",
	Key:        "Key",
	Secret:     "Secret",
	Comment:    "Comment",
	CreateTime: "CreateTime",
	UpdateTime: "UpdateTime",
	Enabled:    "Enabled",
}

func (nodes *Nodes) Fields() *NodesFields {
	table := nodes.T()

	return &NodesFields{
		ID:         table.F(NodesField.ID),
		NodesID:    table.F(NodesField.NodesID),
		Key:        table.F(NodesField.Key),
		Secret:     table.F(NodesField.Secret),
		Comment:    table.F(NodesField.Comment),
		CreateTime: table.F(NodesField.CreateTime),
		UpdateTime: table.F(NodesField.UpdateTime),
		Enabled:    table.F(NodesField.Enabled),
	}
}

func (nodes *Nodes) IndexFieldNames() []string {
	return []string{"ID", "Key", "NodesID"}
}

func (nodes *Nodes) ConditionByStruct() *github_com_johnnyeven_libtools_sqlx_builder.Condition {
	table := nodes.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(nodes)

	conditions := []*github_com_johnnyeven_libtools_sqlx_builder.Condition{}

	for _, fieldName := range nodes.IndexFieldNames() {
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

func (nodes *Nodes) PrimaryKey() github_com_johnnyeven_libtools_sqlx.FieldNames {
	return github_com_johnnyeven_libtools_sqlx.FieldNames{"ID"}
}
func (nodes *Nodes) UniqueIndexes() github_com_johnnyeven_libtools_sqlx.Indexes {
	return github_com_johnnyeven_libtools_sqlx.Indexes{
		"U_key":      github_com_johnnyeven_libtools_sqlx.FieldNames{"Key", "Enabled"},
		"U_nodes_id": github_com_johnnyeven_libtools_sqlx.FieldNames{"NodesID", "Enabled"},
	}
}
func (nodes *Nodes) Comments() map[string]string {
	return map[string]string{
		"Comment":    "描述",
		"CreateTime": "",
		"Enabled":    "",
		"ID":         "",
		"Key":        "key",
		"NodesID":    "业务ID",
		"Secret":     "secret",
		"UpdateTime": "",
	}
}

func (nodes *Nodes) Create(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if nodes.CreateTime.IsZero() {
		nodes.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	nodes.UpdateTime = nodes.CreateTime

	stmt := nodes.D().
		Insert(nodes).
		Comment("Nodes.Create")

	dbRet := db.Do(stmt)
	err := dbRet.Err()

	if err == nil {
		lastInsertID, _ := dbRet.LastInsertId()
		nodes.ID = uint64(lastInsertID)
	}

	return err
}

func (nodes *Nodes) DeleteByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (err error) {
	table := nodes.T()

	stmt := table.Delete().
		Comment("Nodes.DeleteByStruct").
		Where(nodes.ConditionByStruct())

	err = db.Do(stmt).Err()
	return
}

func (nodes *Nodes) CreateOnDuplicateWithUpdateFields(db *github_com_johnnyeven_libtools_sqlx.DB, updateFields []string) error {
	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	if nodes.CreateTime.IsZero() {
		nodes.CreateTime = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}
	nodes.UpdateTime = nodes.CreateTime

	table := nodes.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(nodes, updateFields...)

	delete(fieldValues, "ID")

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	m := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		m[field] = true
	}

	// fields of unique index can not update
	delete(m, "CreateTime")

	for _, fieldNames := range nodes.UniqueIndexes() {
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
		Comment("Nodes.CreateOnDuplicateWithUpdateFields")

	return db.Do(stmt).Err()
}

func (nodes *Nodes) FetchByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Select().
		Comment("Nodes.FetchByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(nodes.ID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) FetchByIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Select().
		Comment("Nodes.FetchByIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(nodes.ID),
			table.F("Enabled").Eq(nodes.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) DeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Delete().
		Comment("Nodes.DeleteByID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(nodes.ID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) UpdateByIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Nodes.UpdateByIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(nodes.ID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	dbRet := db.Do(stmt).Scan(nodes)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return nodes.FetchByID(db)
	}
	return nil
}

func (nodes *Nodes) UpdateByIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(nodes, zeroFields...)
	return nodes.UpdateByIDWithMap(db, fieldValues)
}

func (nodes *Nodes) SoftDeleteByID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Nodes.SoftDeleteByID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("ID").Eq(nodes.ID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	dbRet := db.Do(stmt).Scan(nodes)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return nodes.DeleteByID(db)
		}
		return err
	}
	return nil
}

func (nodes *Nodes) FetchByNodesID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Select().
		Comment("Nodes.FetchByNodesID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("NodesID").Eq(nodes.NodesID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) FetchByNodesIDForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Select().
		Comment("Nodes.FetchByNodesIDForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("NodesID").Eq(nodes.NodesID),
			table.F("Enabled").Eq(nodes.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) DeleteByNodesID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Delete().
		Comment("Nodes.DeleteByNodesID").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("NodesID").Eq(nodes.NodesID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) UpdateByNodesIDWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Nodes.UpdateByNodesIDWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("NodesID").Eq(nodes.NodesID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	dbRet := db.Do(stmt).Scan(nodes)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return nodes.FetchByNodesID(db)
	}
	return nil
}

func (nodes *Nodes) UpdateByNodesIDWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(nodes, zeroFields...)
	return nodes.UpdateByNodesIDWithMap(db, fieldValues)
}

func (nodes *Nodes) SoftDeleteByNodesID(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Nodes.SoftDeleteByNodesID").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("NodesID").Eq(nodes.NodesID),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	dbRet := db.Do(stmt).Scan(nodes)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return nodes.DeleteByNodesID(db)
		}
		return err
	}
	return nil
}

func (nodes *Nodes) FetchByKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Select().
		Comment("Nodes.FetchByKey").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("Key").Eq(nodes.Key),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) FetchByKeyForUpdate(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Select().
		Comment("Nodes.FetchByKeyForUpdate").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("Key").Eq(nodes.Key),
			table.F("Enabled").Eq(nodes.Enabled),
		)).
		ForUpdate()

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) DeleteByKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()
	stmt := table.Delete().
		Comment("Nodes.DeleteByKey").
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("Key").Eq(nodes.Key),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	return db.Do(stmt).Scan(nodes).Err()
}

func (nodes *Nodes) UpdateByKeyWithMap(db *github_com_johnnyeven_libtools_sqlx.DB, fieldValues github_com_johnnyeven_libtools_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()

	delete(fieldValues, "ID")

	stmt := table.Update().
		Comment("Nodes.UpdateByKeyWithMap").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("Key").Eq(nodes.Key),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	dbRet := db.Do(stmt).Scan(nodes)
	err := dbRet.Err()
	if err != nil {
		return err
	}

	rowsAffected, _ := dbRet.RowsAffected()
	if rowsAffected == 0 {
		return nodes.FetchByKey(db)
	}
	return nil
}

func (nodes *Nodes) UpdateByKeyWithStruct(db *github_com_johnnyeven_libtools_sqlx.DB, zeroFields ...string) error {
	fieldValues := github_com_johnnyeven_libtools_sqlx.FieldValuesFromStructByNonZero(nodes, zeroFields...)
	return nodes.UpdateByKeyWithMap(db, fieldValues)
}

func (nodes *Nodes) SoftDeleteByKey(db *github_com_johnnyeven_libtools_sqlx.DB) error {
	nodes.Enabled = github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE

	table := nodes.T()

	fieldValues := github_com_johnnyeven_libtools_sqlx_builder.FieldValues{}
	fieldValues["Enabled"] = github_com_johnnyeven_libtools_courier_enumeration.BOOL__FALSE

	if _, ok := fieldValues["UpdateTime"]; !ok {
		fieldValues["UpdateTime"] = github_com_johnnyeven_libtools_timelib.MySQLTimestamp(time.Now())
	}

	stmt := table.Update().
		Comment("Nodes.SoftDeleteByKey").
		Set(table.AssignsByFieldValues(fieldValues)...).
		Where(github_com_johnnyeven_libtools_sqlx_builder.And(
			table.F("Key").Eq(nodes.Key),
			table.F("Enabled").Eq(nodes.Enabled),
		))

	dbRet := db.Do(stmt).Scan(nodes)
	err := dbRet.Err()
	if err != nil {
		dbErr := github_com_johnnyeven_libtools_sqlx.DBErr(err)
		if dbErr.IsConflict() {
			return nodes.DeleteByKey(db)
		}
		return err
	}
	return nil
}

type NodesList []Nodes

// deprecated
func (nodesList *NodesList) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (count int32, err error) {
	*nodesList, count, err = (&Nodes{}).FetchList(db, size, offset, conditions...)
	return
}

func (nodes *Nodes) FetchList(db *github_com_johnnyeven_libtools_sqlx.DB, size int32, offset int32, conditions ...*github_com_johnnyeven_libtools_sqlx_builder.Condition) (nodesList NodesList, count int32, err error) {
	nodesList = NodesList{}

	table := nodes.T()

	condition := github_com_johnnyeven_libtools_sqlx_builder.And(conditions...)

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Nodes.FetchList").
		Where(condition)

	errForCount := db.Do(stmt.For(github_com_johnnyeven_libtools_sqlx_builder.Count(github_com_johnnyeven_libtools_sqlx_builder.Star()))).Scan(&count).Err()
	if errForCount != nil {
		err = errForCount
		return
	}

	stmt = stmt.Limit(size).Offset(offset)

	stmt = stmt.OrderDescBy(table.F("CreateTime"))

	err = db.Do(stmt).Scan(&nodesList).Err()

	return
}

func (nodes *Nodes) List(db *github_com_johnnyeven_libtools_sqlx.DB, condition *github_com_johnnyeven_libtools_sqlx_builder.Condition) (nodesList NodesList, err error) {
	nodesList = NodesList{}

	table := nodes.T()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Nodes.List").
		Where(condition)

	err = db.Do(stmt).Scan(&nodesList).Err()

	return
}

func (nodes *Nodes) ListByStruct(db *github_com_johnnyeven_libtools_sqlx.DB) (nodesList NodesList, err error) {
	nodesList = NodesList{}

	table := nodes.T()

	condition := nodes.ConditionByStruct()

	condition = github_com_johnnyeven_libtools_sqlx_builder.And(condition, table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Nodes.ListByStruct").
		Where(condition)

	err = db.Do(stmt).Scan(&nodesList).Err()

	return
}

// deprecated
func (nodesList *NodesList) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (err error) {
	*nodesList, err = (&Nodes{}).BatchFetchByIDList(db, idList)
	return
}

func (nodes *Nodes) BatchFetchByIDList(db *github_com_johnnyeven_libtools_sqlx.DB, idList []uint64) (nodesList NodesList, err error) {
	if len(idList) == 0 {
		return NodesList{}, nil
	}

	table := nodes.T()

	condition := table.F("ID").In(idList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Nodes.BatchFetchByIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&nodesList).Err()

	return
}

// deprecated
func (nodesList *NodesList) BatchFetchByKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, keyList []string) (err error) {
	*nodesList, err = (&Nodes{}).BatchFetchByKeyList(db, keyList)
	return
}

func (nodes *Nodes) BatchFetchByKeyList(db *github_com_johnnyeven_libtools_sqlx.DB, keyList []string) (nodesList NodesList, err error) {
	if len(keyList) == 0 {
		return NodesList{}, nil
	}

	table := nodes.T()

	condition := table.F("Key").In(keyList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Nodes.BatchFetchByKeyList").
		Where(condition)

	err = db.Do(stmt).Scan(&nodesList).Err()

	return
}

// deprecated
func (nodesList *NodesList) BatchFetchByNodesIDList(db *github_com_johnnyeven_libtools_sqlx.DB, nodesIDList []uint64) (err error) {
	*nodesList, err = (&Nodes{}).BatchFetchByNodesIDList(db, nodesIDList)
	return
}

func (nodes *Nodes) BatchFetchByNodesIDList(db *github_com_johnnyeven_libtools_sqlx.DB, nodesIDList []uint64) (nodesList NodesList, err error) {
	if len(nodesIDList) == 0 {
		return NodesList{}, nil
	}

	table := nodes.T()

	condition := table.F("NodesID").In(nodesIDList)

	condition = condition.And(table.F("Enabled").Eq(github_com_johnnyeven_libtools_courier_enumeration.BOOL__TRUE))

	stmt := table.Select().
		Comment("Nodes.BatchFetchByNodesIDList").
		Where(condition)

	err = db.Do(stmt).Scan(&nodesList).Err()

	return
}
