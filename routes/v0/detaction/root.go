package detaction

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(Detection{})

type Detection struct {
	courier.EmptyOperator
}

func (Detection) Path() string {
	return "/detections"
}
