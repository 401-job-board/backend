package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	msgPtr := s3Event.Records[0].S3.Object.Key
	topicPtr := os.Getenv("TOPIC_ARN")
	flag.Parse()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, resp := svc.PublishRequest(&sns.PublishInput{
		Message:  &msgPtr,
		TopicArn: &topicPtr,
	})
	err := result.Send()

	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println(resp)

	}
}
