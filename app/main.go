package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/common"
	. "github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/iam"
	"github.com/felixa1996/go_next_be/app/infra/message"
)

var (
	sqsSvc *sqs.SQS
)

type SqsMessage struct {
	uri     string
	message *sqs.Message
}

func pollMessages(chn chan<- *sqs.Message) {

	for {
		output, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String("https://sqs.us-east-1.amazonaws.com/065561208089/test_sqs_data.fifo"),
			MaxNumberOfMessages: aws.Int64(2),
			WaitTimeSeconds:     aws.Int64(15),
		})

		if err != nil {
			fmt.Println("failed to fetch sqs message %v", err)
		}

		for _, message := range output.Messages {
			chn <- message
		}

	}

}

func handleMessage(msg *sqs.Message) {
	fmt.Println("RECEIVING MESSAGE >>> ")
	fmt.Println(*msg.Body)
}

func deleteMessage(msg *sqs.Message) {
	sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String("https://sqs.us-east-1.amazonaws.com/065561208089/test_sqs_data.fifo"),
		ReceiptHandle: msg.ReceiptHandle,
	})
}

func main() {
	// init config
	ReadConfigFromEnv()

	// init logging
	logger := InitLogging()
	// defer logger.Sync()

	// init database
	dbManager := database.NewDatabaseManager(logger, viper.GetString("MONGODB_URI"), viper.GetString("MONGODB_DB"))

	keycloakIam := iam.NewKeycloakIAM()

	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-1")},
	// )
	// if err != nil {
	// 	log.Fatalf("unable to load SDK config, %v", err)
	// }

	// sqsSvc = sqs.New(sess)

	// chnMessages := make(chan *sqs.Message, 2)
	// go pollMessages(chnMessages)

	// for message := range chnMessages {
	// 	handleMessage(message)
	// 	deleteMessage(message)
	// }

	// init sqs
	sqsSession, _ := message.NewSqs(logger)
	sqsSvc = sqs.New(sqsSession)

	// init app
	InitApp(dbManager, sqsSvc, logger, keycloakIam)

	err := common.Application.Echo.Start(":" + viper.GetString("PORT"))
	if err != nil {
		logger.Fatal("Failed to start app", zap.Error(err))
		panic(err)
	}
}
