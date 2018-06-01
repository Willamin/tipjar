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

type Product struct {
  Cost uint64
  Name string
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
  token := q.Get("stripeToken")

  product := Product{
    Cost: 100,
    Name: "Tip",
  }

	chargeParams := &stripe.ChargeParams{
		Amount:   product.Cost,
		Currency: "usd",
		Desc:     fmt.Sprintf("Charge for %s", product.Name),
	}
  chargeParams.SetSource(token)

	chargeResponse, err := charge.New(chargeParams)

  return events.APIGatewayProxyResponse{
    Body:     fmt.Sprintf("%#v", chargeResponse),
    StatusCode: 200,
    }, err
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
