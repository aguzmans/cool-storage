package glacier

import (
	// "fmt"

	// "context"

	"github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/aws/session"
	"context"
	"log"

	"github.com/aguzmans/cool-storage/configread"
)

const (
	S3_REGION = "ap-southeast-1"
)

func Glacier(configFile configread.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(configFile.Aws.Profile))
	if err != nil {
		log.Println(err)
		return
	}
	log.Print(cfg)
	// create an AWS session which can be
	// reused if we're uploading many files
	// s, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(S3_REGION),
	// 	Credentials: credentials.NewStaticCredentials(
	// 		"XXX",
	// 		"YYY",
	// 		""),
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = uploadFileToS3(s, "discharge-letter-787653.pdf")

	// if err != nil {
	// 	log.Fatal(err)
	// }
}
