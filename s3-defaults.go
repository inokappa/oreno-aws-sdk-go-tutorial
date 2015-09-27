package main

import (
    "os"
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/defaults"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {

  defaults.DefaultConfig.Region = aws.String(os.Getenv("AWS_REGION"))
  
  svc := s3.New(nil)
  var params *s3.ListBucketsInput
  result, err := svc.ListBuckets(params)
  if err != nil {
      log.Println(err.Error())
      return
  }
  log.Println(result)
}
