package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsConfig = map[string]string{
	AWS_REGION:            "ap-southeast-1",
	AWS_ACCESS_KEY_ID:     "access-key",
	AWS_SECRET_ACCESS_KEY: "secret",
	AWS_SESSION_TOKEN:     "",
}

// GetAwsSession construct aws Session for sqs configuration
func GetAwsSession() *session.Session {
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

const (
	AWS_REGION            = "AWS_REGION"
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN     = "AWS_SESSION_TOKEN"
)
