package detection

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules"
)

// 物体检测
type ObjectDetection struct {
	tp.CallCtx
}

type ObjectDetectionBody struct {
	Image []byte
}

func (r *Detection) Object(req *ObjectDetectionBody) ([]modules.DetectivedObject, *tp.Status) {
	data := make([]modules.DetectivedObject, 0)
	var err error
	if global.Config.RobotConfiguration.CameraMode == types.CAMERA_MODE__NORMAL {
		return data, nil
	}

	data, err = modules.DetectiveObject(req.Image, global.Config.COCOModel)
	if err != nil {
		statusErr := errors.InternalError.StatusError().WithDesc(err.Error())
		return nil, tp.NewStatus(int32(statusErr.Code), statusErr.Desc, err)
	}

	return data, nil
}
