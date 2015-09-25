package main

import (
    "os"
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {

  credentialsProvider := credentials.NewChainCredentials(
  []credentials.Provider{
    &credentials.EnvProvider{},
    &credentials.SharedCredentialsProvider{Filename: "", Profile: "default"},
  })

  svc := s3.New(&aws.Config{Credentials: credentialsProvider, Region: aws.String(os.Getenv("AWS_REGION"))})
  var params *s3.ListBucketsInput
  result, err := svc.ListBuckets(params)
  if err != nil {
      log.Println(err.Error())
      return
  }
  log.Println(result)
}
