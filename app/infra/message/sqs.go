package message

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
)

// todo create single struct contains url and message
// *sqs.Message

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
