package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

func handler(ctx context.Context) {
	input, _ := json.Marshal(&ctx)

	if os.Getenv("AWS_LAMBDA_LOG_FORMAT") == "JSON" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logLevel := os.Getenv("AWS_LAMBDA_LOG_LEVEL")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("log level configuration error")
	}
	logrus.SetLevel(level)

	logrus.WithField("AWS_LAMBDA_LOG_FORMAT", os.Getenv("AWS_LAMBDA_LOG_FORMAT")).Info()
	logrus.WithField("AWS_LAMBDA_LOG_LEVEL", os.Getenv("AWS_LAMBDA_LOG_LEVEL")).Info()

	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Trace("trace")
	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Debug("debug")
	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Info("info")
	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Warn("warn")
	logrus.WithFields(logrus.Fields{
		"input": string(input),
	}).Error("error")

	input = append(input, '}')

	var m map[string]interface{}
	err = json.Unmarshal(input, &m)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"input": string(input),
			"error": err,
		}).Error("input json unmarshal error")
	}

	// The following 5 entries are valid, and the output to CloudWatch logs depends on the application's log level in Lambda config.
	fmt.Fprintf(os.Stderr, `{"msg":"mylog","out":"stderr","level":"%s"}`+"\n", logrus.TraceLevel.String())
	fmt.Fprintf(os.Stderr, `{"msg":"mylog","out":"stderr","level":"%s"}`+"\n", logrus.DebugLevel.String())
	fmt.Fprintf(os.Stderr, `{"msg":"mylog","out":"stderr","level":"%s"}`+"\n", logrus.InfoLevel.String())
	fmt.Fprintf(os.Stderr, `{"msg":"mylog","out":"stderr","level":"%s"}`+"\n", logrus.WarnLevel.String())
	fmt.Fprintf(os.Stderr, `{"msg":"mylog","out":"stderr","level":"%s"}`+"\n", logrus.ErrorLevel.String())

	// The following 5 entries are valid, and the output to CloudWatch logs depends on the application's log level in Lambda config.
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"%s"}`+"\n", logrus.TraceLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"%s"}`+"\n", logrus.DebugLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"%s"}`+"\n", logrus.InfoLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"%s"}`+"\n", logrus.WarnLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"%s"}`+"\n", logrus.ErrorLevel.String())

	// The following 5 entries are invalid
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","Level":"%s"}`+"\n", logrus.ErrorLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","LEVEL":"%s"}`+"\n", logrus.ErrorLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","logLevel":"%s"}`+"\n", logrus.ErrorLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","log-level":"%s"}`+"\n", logrus.ErrorLevel.String())
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","log_level":"%s"}`+"\n", logrus.ErrorLevel.String())

	// The following first 2 entries are valid
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"ERROR"}`+"\n")
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"Error"}`+"\n")
	fmt.Fprintf(os.Stdout, `{"msg":"mylog","out":"stdout","level":"eRROr"}`+"\n")
}

func main() {
	lambda.Start(handler)
}
