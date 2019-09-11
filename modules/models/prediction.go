package models

import (
	"bytes"
	"image/jpeg"
)

func DetectiveObject(image []byte, model *COCOObjectDetectiveModel) ([]DetectivedObject, error) {
	data := make([]DetectivedObject, 0)
	reader := bytes.NewReader(image)
	_, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}

	num, boxes, classes, probabilities, err := model.Predict(image)
	if err != nil {
		return nil, err
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
