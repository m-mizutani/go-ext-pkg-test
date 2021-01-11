package chain_test

import (
	"os"
	"testing"

	"github.com/m-mizutani/go-ext-pkg-test/chain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsecaseWithDynamo(t *testing.T) {
	tableName := os.Getenv("TEST_TABLE_NAME")
	if tableName == "" {
		t.Skip("TEST_TABLE_NAME is not set")
	}
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		t.Skip("AWS_REGION is not set")
	}

	testUsecase(t, chain.NewDynamoDB(awsRegion, tableName))
}

func TestUsecaseWithInMemory(t *testing.T) {
	testUsecase(t, chain.NewInMemoryDB())
}

func testUsecase(t *testing.T, repo chain.Repository) {
	usecase := chain.NewUsecase(repo)

	user0, err := usecase.GetUsers()
	require.NoError(t, err)
	assert.Equal(t, 0, len(user0))

	require.NoError(t, usecase.SaveUser("blue"))
	users1, err := usecase.GetUsers()
	require.NoError(t, err)
	assert.Equal(t, 1, len(users1))
	assert.Contains(t, users1, "blue")

	require.NoError(t, usecase.SaveUser("orange"))
	users2, err := usecase.GetUsers()
	require.NoError(t, err)
	assert.Equal(t, 2, len(users2))
	assert.Contains(t, users2, "blue")
	assert.Contains(t, users2, "orange")
}
