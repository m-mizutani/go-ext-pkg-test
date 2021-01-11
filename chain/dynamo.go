package chain

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

type DynamoDB struct {
	table dynamo.Table
}

type item struct {
	PKey  string `dynamo:"pk"`
	SKey  string `dynamo:"sk"`
	Value string `dynamo:"value"`
}

func NewDynamoDB(region, table string) *DynamoDB {
	ssn, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		panic(err)
	}

	return &DynamoDB{table: dynamo.New(ssn).Table(table)}
}

func (x *DynamoDB) AppendItem(key, value string) error {
	i := item{
		PKey:  key,
		SKey:  uuid.New().String(),
		Value: value,
	}
	if err := x.table.Put(i).Run(); err != nil {
		return err
	}
	return nil
}

func (x *DynamoDB) FetchItems(key string) ([]string, error) {
	var items []item
	if err := x.table.Get("pk", items).All(&items); err != nil {
		return nil, err
	}

	var res []string
	for _, i := range items {
		res = append(res, i.Value)
	}
	return res, nil
}
