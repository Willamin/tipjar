package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"log"
	"net/url"
	"os"
)

const (
	AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
	StripeApiKey             = "STRIPE_API_KEY"
)

type Product struct {
	Name string
	Cost uint64
}

func allProducts() []Product {
	return []Product{
		Product{
			Name: "Sticker",
			Cost: 200,
		},
		Product{
			Name: "Hat",
			Cost: 1000,
		},
	}
}

func errorResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Something went wrong",
		StatusCode: 500,
	}, err
}

func getTokenFromBody(body string) (string, error) {
	q, err := url.ParseQuery(body)

	token := q.Get("stripeToken")
	return token, err
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error = nil

	log.Printf("Handling request:")
	log.Printf("%v", request)
	log.Printf("Params:")
	log.Printf("%v", request.body)

	key, found := os.LookupEnv(StripeApiKey)
	stripe.Key = key
	if found == false {
		err = errors.New(fmt.Sprintf("%s not found", StripeApiKey))
		return errorResponse(err)
	}

	token, err := getTokenFromBody(request.Body)
	if err != nil {
		return errorResponse(err)
	}

	product := allProducts()[0]

	chargeParams := &stripe.ChargeParams{
		Amount:   product.Cost,
		Currency: "usd",
		Desc:     fmt.Sprintf("Charge for %s", product.Name),
	}
	chargeParams.SetSource(token)

	_, err = charge.New(chargeParams)
	if err != nil {
		return errorResponse(err)
	}

	headers := map[string]string{
		"Location": "https://tip.wflewis.com/thanks",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers:    headers,
	}, nil
}

func main() {
	_, ok := os.LookupEnv(AWSLambdaFunctionVersion)
	if ok {
		log.Printf("Running in AWS lambda environment, starting lambda handler.")
		lambda.Start(HandleRequest)
		os.Exit(0)
	}

	log.Printf("Not running in AWS lambda environment.")
	os.Exit(0)
}
