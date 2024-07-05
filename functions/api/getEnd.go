package api

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// NOTES:
// why are you json like that it's because Go lang understands string and int it doesn't understand Json ,
// and we will be sending json to the lambda function so that why we have defined how the json will look like


type MyEvent struct {
	Name string `json:"What is your name?"`
	Age int `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HanderLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}


func main() {
	lambda.Start(HanderLambdaEvent)
}