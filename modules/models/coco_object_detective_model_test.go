package models

import (
	"io/ioutil"
	"testing"
)

func TestLoadModel(t *testing.T) {
	mgr := COCOObjectDetectiveModel{
		ModelPath: "./config/mobilenet",
	}
	mgr.Init()

	// load image
	f, err := ioutil.ReadFile("./test.jpg")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	num, boxes, classes, probabilities, err := mgr.Predict(f)

	t.Log(num)
	t.Log(classes)
	t.Log(probabilities)
	t.Log(boxes)
}
