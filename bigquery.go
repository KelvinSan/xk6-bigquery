package xk6_bigquery

import (
	"context"
	"github.com/dailyburn/bigquery/client"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {

	modules.Register("k6/x/bigquery", new(BQ))

}

type BQ struct {
}

type Client struct {
	bqclient *client.Client
}

func (r *BQ) XClient(ctxPtr *context.Context, serviceAccount string) interface{} {
	rt := common.GetRuntime(*ctxPtr)
	return common.Bind(rt, &Client{bqclient: client.New(serviceAccount)}, ctxPtr)
}

func (r *BQ) Query(name string) string {

	return name + "hello"

}