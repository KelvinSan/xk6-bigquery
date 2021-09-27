import xk6_bigquery from 'k6/x/bigquery'

export default function(){

    console.log(xk6_bigquery.query("cloud-dev"))


}