package models

import (
	"bytes"
	"image/jpeg"
)

func DetectiveObject(req *CameraRequest, model *COCOObjectDetectiveModel) ([]DetectivedObject, error) {
	data := make([]DetectivedObject, 0)
	reader := bytes.NewReader(req.Frame)
	_, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}

	num, boxes, classes, probabilities, err := model.Predict(req.Frame)
	if err != nil {
		return nil, err
	}

	for i := 0; i < int(num); i++ {
		data = append(data, DetectivedObject{
			Class:       classes[i],
			Label:       model.GetLabel(i, probabilities, classes),
			Box:         boxes[i],
			Probability: probabilities[i],
		})
	}

	return data, nil
}
