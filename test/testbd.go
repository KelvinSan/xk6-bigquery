package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {

		client,err := bigquery.NewClient(context.Background(),"moonlit-poetry-327116",option.WithCredentialsFile("../demo.json"))

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


}
