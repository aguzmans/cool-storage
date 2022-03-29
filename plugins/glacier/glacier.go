package glacier

import (
	"context"
	"log"

	"github.com/aguzmans/cool-storage/configread"
	"github.com/aws/aws-sdk-go-v2/config"
)

func Glacier(configFile configread.Profile) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(configFile.Name))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(cfg)
	// sess, err := session.NewSessionWithOptions(session.Options{
	// 	Config: aws.Config{Region: aws.String(region),
	// 			  CredentialsChainVerboseErrors: aws.Bool(true)},
	// 	Profile: profile,
	// })

	// sess, errSess := session.NewSession(&aws.Config{
	// 	Region: aws.String(configFile.Aws.Region)},
	// )
	// if errSess != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Panicln(sess)
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
