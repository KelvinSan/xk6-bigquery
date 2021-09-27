import xk6_bigquery from 'k6/x/bigquery'
import { sleep } from 'k6';

const client = xk6_bigquery.newClient("demo.json")

export default function(){


     
     xk6_bigquery.dataLoader(client,"users","gs://demo-bucket-ks/myFile0.csv","user-data")
     sleep(2);

}