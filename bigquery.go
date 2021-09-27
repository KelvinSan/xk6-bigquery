package xk6_bigquery

import (

	"github.com/dailyburn/bigquery/client"
	"go.k6.io/k6/js/modules"
)

func init() {

	modules.Register("k6/x/bigquery", new(BQ))

}

type BQ struct {
}

func (c *BQ) NewClient(serviceAccount string) *client.Client{

 if serviceAccount == ""{

	panic("Please enter service account json path before continue")

 }

 return client.New(serviceAccount)

}

func (r *BQ) Query(bqclient *client.Client, dataset string, project string, query string) ([][]interface{}, []string, error) {

	rows,headers,err := bqclient.Query(dataset,project,query)

	if err != nil {
		panic("Something happened "+err.Error())
	}

	return rows,headers,err

}
