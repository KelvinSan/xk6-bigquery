import xk6_bigquery from 'k6/x/bigquery'

const client = xk6_bigquery.newClient("demo.json")

export default function(){

    console.log(xk6_bigquery.query(client,"Hello"))


}