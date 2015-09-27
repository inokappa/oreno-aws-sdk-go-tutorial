package main

import (
  "os"
  "log"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/defaults"
  "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type logStream struct {
  logStreamName string
  logGroupName  string
  client        api
}

type api interface {
  CreateLogStream(*cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error)
  PutLogEvents(*cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error)
}

func main() {

  defaults.DefaultConfig.Region = aws.String(os.Getenv("AWS_REGION"))

  containerStream := &logStream{
    logStreamName : "bar",
    logGroupName  : "docker-log",
    client        : cloudwatchlogs.New(nil),
  }
  err := containerStream.create()
  if err != nil {
    log.Println(err)
  }
}

func (l *logStream) create() error {
  params := &cloudwatchlogs.CreateLogStreamInput{
    LogGroupName: aws.String(l.logGroupName),
    LogStreamName: aws.String(l.logStreamName),
  }

  _, err := l.client.CreateLogStream(params)
  if err != nil {
    log.Println(err)
  }
  return err
}
