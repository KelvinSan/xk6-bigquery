import xk6_bigquery from 'k6/x/bigquery'

const client = xk6_bigquery.newClient("hello")

export default function(){

    console.log(xk6_bigquery.query(client,"hello-dev","inside","Select * from table"))


}