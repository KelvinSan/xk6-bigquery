package xk6_bigquery

import (
	"context"
	"encoding/json"
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

type Datasets struct {

Data []string `json:"data"`


}

type QueryResponse struct{

Data []bigquery.Value `json:"data"`


}



func (c *BQ) NewClient(serviceAccount string) *bigquery.Client {

	client,err := bigquery.NewClient(context.Background(),"moonlit-poetry-327116",option.WithCredentialsFile(serviceAccount))

	if err != nil{

		println("Connection Error",err.Error())

	}

	return client

}

func (r *BQ) GetDatasets(bqclient *bigquery.Client) string {

	it := bqclient.Datasets(context.Background())

	datasets := []string{}

	for {
		dataset, err := it.Next()
		if err == iterator.Done {
			break
		}
		datasets = append(datasets, dataset.DatasetID)
	}

	sets := Datasets{Data: datasets}

	bytes, err := json.Marshal(sets)

	if err != nil {
		panic("Error Occured receiving datasets "+err.Error())
	}

	return string(bytes)

}



func(r *BQ) Query(bqclient *bigquery.Client, query string) string{

q := bqclient.Query(query)

it, err := q.Read(context.Background())

if err != nil {
	panic("Error occured fetching query "+ err.Error())
}

data := []bigquery.Value{}

for {
    var values []bigquery.Value
    err := it.Next(&values)
    if err == iterator.Done {
        break
    }
    if err != nil {
		panic("Error occured fetching query "+ err.Error())
    }

	data = append(data, values)
    
}

fmt.Println(data)

// sets := QueryResponse{Data: data}

// bytes, err := json.Marshal(sets)

// 	if err != nil {
// 		panic("Error Occured receiving datasets "+err.Error())
// 	}

// 	return string(bytes)

return("Hello")	
}
