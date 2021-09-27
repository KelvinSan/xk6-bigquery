package xk6_bigquery

import (
	
	"go.k6.io/k6/js/modules"
)

func init() {

	modules.Register("k6/x/bigquery", new(BQ))

}

type BQ struct {
}

func (r *BQ) Query(dataset string, project string, query string) (string,string,string) {

	return dataset, project, query

}
