package simple

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Client interface {
	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
}

func GetData(client s3Client, bucket, key string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := client.GetObject(input)
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	return ioutil.ReadAll(output.Body)
}
