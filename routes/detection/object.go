package detection

import (
	"bytes"
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"image/jpeg"
)

// 物体检测
type ObjectDetection struct {
	tp.CallCtx
}

type ObjectDetectionBody struct {
	Image []byte
}

func (r *Detection) Object(req *ObjectDetectionBody) ([]DetectivedObject, *tp.Status) {
	data := make([]DetectivedObject, 0)
	if global.Config.RobotConfiguration.CameraMode == types.CAMERA_MODE__NORMAL {
		return data, nil
	}
	reader := bytes.NewReader(req.Image)
	_, err := jpeg.Decode(reader)
	if err != nil {
		statusErr := errors.InternalError.StatusError().WithDesc(err.Error())
		return nil, tp.NewStatus(int32(statusErr.Code), statusErr.Desc, err)
	}

	num, boxes, classes, probabilities, err := global.Config.COCOModel.Predict(req.Image)
	if err != nil {
		statusErr := errors.InternalError.StatusError().WithDesc(err.Error())
		return nil, tp.NewStatus(int32(statusErr.Code), statusErr.Desc, err)
	}

	for i := 0; i < int(num); i++ {
		data = append(data, DetectivedObject{
			Class:       classes[i],
			Box:         boxes[i],
			Probability: probabilities[i],
		})
	}

	return data, nil
}

type DetectivedObject struct {
	Class       float32   `json:"class"`
	Box         []float32 `json:"box"`
	Probability float32   `json:"probability"`
}
