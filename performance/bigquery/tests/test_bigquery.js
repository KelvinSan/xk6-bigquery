import xk6_bigquery from 'k6/x/bigquery'
import { sleep } from 'k6';
import { group } from 'k6';
import {getAllRecords} from '../modules/user_data.js'
import {testDataload} from '../modules/dataload_user_data.js'


const client = xk6_bigquery.newClient("../../../demo.json")

export default function(){


     group("gathering records",() => {

        getAllRecords(client)

     })

     sleep(2)

   //   group("dataload test",() => {

   //      testDataload(client)


   //   })

     sleep(2)
     

}