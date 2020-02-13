package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var sqsConfig = map[string]string{
	SQS_QUEUE_URL: "https://sqs.ap-southeast-1.amazonaws.com/",
}

var sqsInterfaceConfig = map[string]interface{}{
	I_SQS_MAX_NUMBER_OF_MESSAGES: 10,
	I_SQS_WAIT_TIME_SECONDS:      20,
}

// GetSqsAwsSession construct aws Session for sqs configuration
func GetSqsAwsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(Get(AWS_REGION)),
		Credentials: credentials.NewStaticCredentials(
			Get(AWS_ACCESS_KEY_ID),
			Get(AWS_SECRET_ACCESS_KEY),
			Get(AWS_SESSION_TOKEN),
		),
	})
	if nil != err {
		panic(err)
	}

	return sess
}

// GetReceiveMessageInput construct ReceiveMessageInput for sqs configuration
func GetReceiveMessageInput() *sqs.ReceiveMessageInput {
	return &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(GetIfInt64(I_SQS_MAX_NUMBER_OF_MESSAGES)),
		QueueUrl:            aws.String(Get(SQS_QUEUE_URL)),
		WaitTimeSeconds:     aws.Int64(GetIfInt64(I_SQS_WAIT_TIME_SECONDS)),
	}
}

const (
	SQS_QUEUE_URL = "SQS_QUEUE_URL"
	SQS_FLAG      = "SQS_FLAG"
)

const (
	I_SQS_MAX_NUMBER_OF_MESSAGES = "I_SQS_MAX_NUMBER_OF_MESSAGES"
	I_SQS_WAIT_TIME_SECONDS      = "I_SQS_WAIT_TIME_SECONDS"
)
