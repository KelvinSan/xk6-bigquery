package xk6_bigquery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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



func(r *BQ) Query(bqclient *bigquery.Client, query string){

q := bqclient.Query(query)

it, err := q.Read(context.Background())

if err != nil {
	panic("Error occured fetching query "+ err.Error())
}



for {
    var values []bigquery.Value
    err := it.Next(&values)
    if err == iterator.Done {
        break
    }
    if err != nil {
        panic("Error occured fetching query "+ err.Error())
    }
	time.Sleep(100 * time.Millisecond)
    fmt.Println("Query executed here are rows ",values)
}

}

func(r *BQ) DataLoader(bqclient *bigquery.Client, dataset string, gsbucket string,table string){

	myDataset := bqclient.Dataset(dataset)
	gcsRef := bigquery.NewGCSReference(gsbucket)
	gcsRef.AllowJaggedRows = true
	loader := myDataset.Table(table).LoaderFrom(gcsRef)
	loader.CreateDisposition = bigquery.CreateNever
	job, err := loader.Run(context.Background())
	
	if err != nil {
		panic("Job Loader Configuration Failed "+ err.Error())
	}

	status, err2 := job.Status(context.Background())

	if err2 != nil {
        panic("Job Loader Configuration Failed"+ err.Error())
    }
	
    if status.Done() {
        if status.Err() != nil {
            log.Fatalf("Job failed with error %v", status.Err())
        }
		fmt.Println("Job was success data has been loaded")
    }
	

}
