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

type SqsWrapper struct {
	logger *zap.Logger
	sqs    *sqs.SQS
}

func NewSqsWrapper(logger *zap.Logger, sess *session.Session) SqsWrapper {
	sqsSvc := sqs.New(sess)
	return SqsWrapper{
		logger: logger,
		sqs:    sqsSvc,
	}
}

func (s SqsWrapper) PollMessages(queueUrl string, maxNumberMessage int64, waitTimeout int64, chn chan<- *sqs.Message) {
	for {
		output, err := s.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueUrl),
			MaxNumberOfMessages: aws.Int64(maxNumberMessage),
			WaitTimeSeconds:     aws.Int64(waitTimeout),
		})

		if err != nil {
			s.logger.Error("Failed to fetch sqs message",
				zap.String("QueueUrl", queueUrl),
				zap.Error(err),
			)
		}

		for _, message := range output.Messages {
			chn <- message
		}
	}

}

func (s SqsWrapper) DeleteMessage(queueUrl string, msg *sqs.Message) {
	_, err := s.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		s.logger.Error("Failed to delete sqs message",
			zap.String("QueueUrl", queueUrl),
			zap.String("MessageId", *msg.MessageId),
			zap.String("ReceiptHandle", *msg.ReceiptHandle),
			zap.String("MessageBody", *msg.Body),
			zap.Error(err),
		)
	}
	s.logger.Info("Success to delete sqs message", zap.String("QueueUrl", queueUrl),
		zap.String("MessageId", *msg.MessageId),
		zap.String("ReceiptHandle", *msg.ReceiptHandle),
		zap.String("MessageBody", *msg.Body),
	)
}
