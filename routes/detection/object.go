package detection

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

// 物体检测
func (r *Detection) Object(req *models.CameraRequest) ([]models.DetectivedObject, *tp.Status) {
	data, err := models.DetectiveObject(req, global.Config.COCOModel)
	if err != nil {
		statusErr := errors.InternalError.StatusError().WithDesc(err.Error())
		return nil, tp.NewStatus(int32(statusErr.Code), statusErr.Desc, err)
	}

	return data, nil
}
