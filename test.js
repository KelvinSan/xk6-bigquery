import xk6_bigquery from 'k6/x/bigquery'

const client = xk6_bigquery.newClient("demo.json")

export default function(){

    console.log(xk6_bigquery.getDatasets(client,"Hello"))

    xk6_bigquery.query(client,"SELECT * FROM `moonlit-poetry-327116.users.user-data` LIMIT 1000")
    xk6_bigquery.dataLoader(client,"users","gs://demo-bucket-ks/MOCK_DATA.csv","users-data")

}