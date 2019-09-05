package detaction

import (
	"bytes"
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"image/jpeg"
)

func init() {
	Router.Register(courier.NewRouter(ObjectDetection{}))
}

// 物体检测
type ObjectDetection struct {
	httpx.MethodPost
	Body ObjectDetectionBody `name:"body" in:"body"`
}

type ObjectDetectionBody struct {
	Image []byte `json:"image"`
}

func (req ObjectDetection) Path() string {
	return "/object"
}

func (req ObjectDetection) Output(ctx context.Context) (result interface{}, err error) {
	data := make([]DetectivedObject, 0)
	if global.Config.RobotConfiguration.CameraMode == types.CAMERA_MODE__NORMAL {
		return data, nil
	}
	reader := bytes.NewReader(req.Body.Image)
	_, err = jpeg.Decode(reader)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}

	num, boxes, classes, probabilities, err := global.Config.COCOModel.Predict(req.Body.Image)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
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
