package main

import (
  "os"
    "context"

  "log"
  "github.com/aws/aws-lambda-go/lambda"
)

// Environment variables
const (
  AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
)

func action() error {
  return nil
}

func check(err error) {
  if err != nil {
    log.Fatal("Error: ", err)
  }
}

func HandleRequest(ctx context.Context) error {
  log.Printf("Handling request")
  err := action()
  log.Printf("Done handling request")
  return err
}

func main() {
  _, ok := os.LookupEnv(AWSLambdaFunctionVersion)
  if ok {
    log.Printf("Running in AWS lambda environment, starting lambda handler.")
    lambda.Start(HandleRequest)
    os.Exit(0)
  }

  log.Printf("Not running in AWS lambda environment, running test request.")
  err := action()
  check(err)
}
