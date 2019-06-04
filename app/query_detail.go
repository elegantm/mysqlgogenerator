package app

import (
	"log"
)

func QueryTableColumInfo(table string,schema string)(*[]TableInfo)  {
	 sql := "SELECT DISTINCT `table_name`,`column_name`,`data_type`,`column_comment`  FROM information_schema.`COLUMNS` WHERE table_name = ? AND table_schema =? "

	 var tableList []TableInfo

	 err := dbx.Select(&tableList,sql,table,schema)
	 if err != nil {
	 	log.Fatalf("Query Table Column Info error : %s",err)
	 	return nil
	 }

	 if len(tableList)==0 {
	 	log.Fatalf("empty table column %s",table)
		 return nil
	 }

	 for _,v := range tableList{
	 	log.Printf("table info: %+v",v)
	 }

	return &tableList
}