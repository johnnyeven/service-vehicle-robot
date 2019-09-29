package models

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func LoadLabels(labelsFile string) []string {
	labels := make([]string, 0)
	file, err := os.Open(labelsFile)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logrus.Infof("ERROR: failed to read %s: %v", labelsFile, err)
	}
	return labels
}

func GetLabel(labels []string, idx int, probabilities []float32, classes []float32) string {
	index := int(classes[idx])
	label := fmt.Sprintf("%s (%2.0f%%)", labels[index], probabilities[idx]*100.0)

	return label
}
