package modules

import (
	"bytes"
	"github.com/sirupsen/logrus"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"image"
	"image/jpeg"
	_ "image/jpeg"
)

type COCOObjectDetectiveModel struct {
	ModelPath string `conf:"env"`
	model     *tf.SavedModel
	session   *tf.Session

	graph *tf.Graph

	inputImageOp    *tf.Operation
	outputBoxesOp   *tf.Operation
	outputScoresOp  *tf.Operation
	outputClassesOp *tf.Operation
	outputNumOp     *tf.Operation
}

func (mgr *COCOObjectDetectiveModel) Init() {
	if mgr.ModelPath == "" {
		logrus.Panic("[COCOObjectDetectiveModel] ModelPath should not be empty")
	}

	var err error
	mgr.model, err = mgr.loadModel([]string{"serve"})
	if err != nil {
		logrus.Panicf("[COCOObjectDetectiveModel] loadModel err: %v", err)
	}

	mgr.graph = mgr.model.Graph
	mgr.session = mgr.model.Session
	mgr.inputImageOp = mgr.graph.Operation("image_tensor")
	mgr.outputBoxesOp = mgr.graph.Operation("detection_boxes")
	mgr.outputScoresOp = mgr.graph.Operation("detection_scores")
	mgr.outputClassesOp = mgr.graph.Operation("detection_classes")
	mgr.outputNumOp = mgr.graph.Operation("num_detections")
}

func (mgr *COCOObjectDetectiveModel) Close() error {
	return mgr.session.Close()
}

func (mgr *COCOObjectDetectiveModel) Predict(sourceImage []byte) (num float32, boxes [][]float32, classes []float32, probabilities []float32, err error) {
	tensor, _, err := makeTensorFromImage(sourceImage)
	if err != nil {
		return
	}

	output, err := mgr.session.Run(
		map[tf.Output]*tf.Tensor{
			mgr.inputImageOp.Output(0): tensor,
		},
		[]tf.Output{
			mgr.outputBoxesOp.Output(0),
			mgr.outputScoresOp.Output(0),
			mgr.outputClassesOp.Output(0),
			mgr.outputNumOp.Output(0),
		}, nil)
	if err != nil {
		return
	}

	probabilities = output[1].Value().([][]float32)[0]
	classes = output[2].Value().([][]float32)[0]
	boxes = output[0].Value().([][][]float32)[0]
	num = output[3].Value().([]float32)[0]

	return
}

func (mgr *COCOObjectDetectiveModel) loadModel(modelNames []string) (*tf.SavedModel, error) {
	model, err := tf.LoadSavedModel(mgr.ModelPath, modelNames, nil) // 载入模型
	if err != nil {
		return nil, err
	}

	return model, nil
}

func makeTensorFromImage(content []byte) (*tf.Tensor, image.Image, error) {
	r := bytes.NewReader(content)
	img, _, err := image.Decode(r)

	if err != nil {
		return nil, nil, err
	}

	// DecodeJpeg uses a scalar String-valued tensor as input.
	tensor, err := tf.NewTensor(string(content))
	if err != nil {
		return nil, nil, err
	}
	// Creates a tensorflow graph to decode the jpeg image
	graph, input, output, err := decodeJpegGraph()
	if err != nil {
		return nil, nil, err
	}
	// Execute that graph to decode this one image
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, nil, err
	}
	defer session.Close()
	normalized, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		return nil, nil, err
	}
	return normalized[0], img, nil
}

func decodeJpegGraph() (graph *tf.Graph, input, output tf.Output, err error) {
	s := op.NewScope()
	input = op.Placeholder(s, tf.String)
	output = op.ExpandDims(s,
		op.DecodeJpeg(s, input, op.DecodeJpegChannels(3)),
		op.Const(s.SubScope("make_batch"), int32(0)))
	graph, err = s.Finalize()
	return graph, input, output, err
}

type DetectivedObject struct {
	Class       float32   `json:"class"`
	Box         []float32 `json:"box"`
	Probability float32   `json:"probability"`
}

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
