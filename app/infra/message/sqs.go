package message

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

type SqsIncomingMessage struct {
	QueueUrl string
	Message  *sqs.Message
}

func NewSqs(logger *zap.Logger) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		logger.Error("Failed to load AWS Config", zap.Error(err))
		return nil, err
	}
	// todo need check
	// logger.Info("Success Initializing AWS SDK", zap.String("URL", *sess.Config.Endpoint))
	return sess, err
}
