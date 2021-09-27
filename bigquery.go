package xk6_bigquery

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.k6.io/k6/js/modules"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func init() {

	modules.Register("k6/x/bigquery", new(BQ))

}

type BQ struct {
}

func (c *BQ) NewClient(serviceAccount string) *bigquery.Client {

	client,err := bigquery.NewClient(context.Background(),"moonlit-poetry-327116",option.WithCredentialsFile(serviceAccount))

	if err != nil{

		println("Connection Error",err.Error())

	}

	it := client.Datasets(context.Background())

	for {
		dataset, err := it.Next()
		if err == iterator.Done {
			break
		}
		fmt.Println(dataset.DatasetID)
	}

	println("Client Logged In")

	return client

}

func (r *BQ) Query(bqclient *bigquery.Client, query string) string {

	return query

}
