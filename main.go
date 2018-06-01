package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"log"
)

// Environment variables
const (
	AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
	StripeApiKey             = "STRIPE_API_KEY"
)

func check(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func HandleRequest(ctx context.Context) error {
	log.Printf("Handling request")
  stripe.Key = os.LookupEnv(StripeApiKey)
  token := r.FormValue("stripeToken")

  params := &stripe.ChargeParams{
    Amount:      100,
    Currency:    "usd",
    Description: "Tip",
  }
  params.SetSource(token)
  ch, err := charge.New(params)

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

	log.Printf("Not running in AWS lambda environment, exiting.")
	check(err)
}
