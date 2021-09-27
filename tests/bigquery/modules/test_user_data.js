import xk6_bigquery from 'k6/x/bigquery'
import { sleep } from 'k6';

export function getAllRecords(){

    xk6_bigquery.query(client,"SELECT * FROM `moonlit-poetry-327116.users.user-data` LIMIT 1000")
    sleep(1)

}