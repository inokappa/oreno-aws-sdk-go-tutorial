package main

import (
  "os"
  "log"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
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

  credentialsProvider := credentials.NewChainCredentials(
  []credentials.Provider{
    &credentials.EnvProvider{},
    &credentials.SharedCredentialsProvider{Filename: "", Profile: "default"},
  })

  containerStream := &logStream{
    logStreamName : "bar",
    logGroupName  : "docker-log",
    client        : cloudwatchlogs.New(&aws.Config{Credentials: credentialsProvider, Region: aws.String(os.Getenv("AWS_REGION"))}),
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
