package main

import (
  "fmt"
	"os"
	"github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"log"
  "net/url"
)

// Environment variables
const (
	AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
	StripeApiKey             = "STRIPE_API_KEY"
)

type Request struct {
    Token     string `schema:"stripeToken"`
    TokenType string `schema:"stripeTokenType"`
    Email     string `schema:"stripeEmail"`
}

func die(err error) (events.APIGatewayProxyResponse, error) {
   return events.APIGatewayProxyResponse{
    Body:       "Something went wrong",
    StatusCode: 500,
  }, err
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)  {
	log.Printf("Handling request")
	stripe.Key, _ = os.LookupEnv(StripeApiKey)

  q, err := url.ParseQuery(request.Body)
  if err != nil {
    return die(err)
  }
  email := q["stripeEmail"][0]

  return events.APIGatewayProxyResponse{
    Body:       fmt.Sprintf("Thanks, %s!", email),
    StatusCode: 200,
  }, nil


	chargeParams := &stripe.ChargeParams{
		Amount:   2000,
		Currency: "usd",
		Desc:     "Charge for isabella.brown@example.com",
	}
	chargeParams.SetSource("tok_visa")
	_, err = charge.New(chargeParams)
  return events.APIGatewayProxyResponse{}, err
}

func main() {
	_, ok := os.LookupEnv(AWSLambdaFunctionVersion)
	if ok {
		log.Printf("Running in AWS lambda environment, starting lambda handler.")
		lambda.Start(HandleRequest)
		os.Exit(0)
	}

	log.Printf("Not running in AWS lambda environment, exiting.")
	os.Exit(1)
}
