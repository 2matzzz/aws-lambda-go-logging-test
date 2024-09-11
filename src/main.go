package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

func handler(ctx context.Context) {
	input, _ := json.Marshal(&ctx)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Trace("trace")
	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Debug("debug")
	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Info("info")

	input = append(input, '}')

	var m map[string]interface{}
	err := json.Unmarshal(input, &m)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"input": string(input),
			"error": err,
		}).Error("input json unmarshal error")
		return
	}

	// fmt.Fprintf(os.Stderr, "context: %s", string(json))
}

func main() {
	lambda.Start(handler)
}
