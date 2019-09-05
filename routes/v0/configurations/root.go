package configurations

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(ConfigurationGroup{})

type ConfigurationGroup struct {
	courier.EmptyOperator
}

func (ConfigurationGroup) Path() string {
	return "/configurations"
}
