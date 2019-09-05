package modules

import (
	"github.com/johnnyeven/eden-library/clients/client_id"
	"github.com/johnnyeven/eden-library/modules"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/service-vehicle-robot/database"
)

type CreateConfigurationBody struct {
	// StackID
	StackID uint64 `db:"F_stack_id" json:"stackID,string"`
	// Key
	Key string `db:"F_key" json:"key"`
	// Value
	Value string `db:"F_value" json:"value"`
}

func CreateConfiguration(req CreateConfigurationBody, db *sqlx.DB, clientID *client_id.ClientID) error {
	id, err := modules.NewUniqueID(clientID)
	if err != nil {
		return err
	}
	model := &database.Configuration{
		ConfigurationID: id,
		StackID:         req.StackID,
		Key:             req.Key,
		Value:           req.Value,
	}

	return model.Create(db)
}
