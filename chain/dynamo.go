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

func NewDynamoDB(region, table string) *DynamoDB {
	ssn, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		panic(err)
	}

	return &DynamoDB{table: dynamo.New(ssn).Table(table)}
}

func (x *DynamoDB) SaveItem(key, value string) error {
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

func (x *DynamoDB) GetItems(key string) ([]string, error) {
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
