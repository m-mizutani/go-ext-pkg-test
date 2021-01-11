package simple_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/m-mizutani/go-ext-pkg-test/simple"
)

type dummyS3Client struct {
	input []*s3.GetObjectInput
}

func (x *dummyS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	x.input = append(x.input, input)
	return &s3.GetObjectOutput{
		Body: ioutil.NopCloser(strings.NewReader("blue")),
	}, nil
}

func TestS3Access(t *testing.T) {
	var client dummyS3Client
	data, err := simple.GetData(&client, "mybucket", "path/to/object")
	require.NoError(t, err)
	require.Equal(t, 1, len(client.input))
	assert.Equal(t, "mybucket", *client.input[0].Bucket)
	assert.Equal(t, "path/to/object", *client.input[0].Key)
	assert.Equal(t, "blue", string(data))
}
