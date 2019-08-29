package detaction

import (
	"bytes"
	"context"
	"fmt"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"image/jpeg"
	"os"
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
	reader := bytes.NewReader(req.Body.Image)
	jpgHandler, err := jpeg.Decode(reader)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}

	imageFile, err := os.OpenFile("./test.jpg", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		err = errors.InternalError.StatusError().WithDesc(err.Error())
		return
	}

	err = jpeg.Encode(imageFile, jpgHandler, &jpeg.Options{Quality: 75})
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
