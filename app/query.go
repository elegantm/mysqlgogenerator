package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type SchemaInfo struct {
	SchemaName string `db:"table_schema"`
	TableName string `db:"table_name"`
}

type Column struct {
	Name string
	DataType string
	ColumnName  string
	Separator string
	ColumnComment  string
}


type TableInfo struct {
	TableName string  `db:"table_name"`
	ColumnName  string `db:"column_name"`
	DataType string `db:"data_type"`
	ColumnComment  string `db:"column_comment"`
	Separator string
}

type TableList struct {
	TableName string
	StrutName string
	Fields    []Column
}

var dbx *sqlx.DB

func InitDatabasesConnect(Addr string) error  {
	var err error
	dbx, err = sqlx.Connect("mysql", Addr)
	if err != nil {
		log.Fatalf("connect data bases error %s",err)
		return  err
	}
	return nil
}



func QuerySchemaName()(string,error)  {
	sql := "SELECT SCHEMA()"

	var err error
	var schema string

	err = dbx.QueryRow(sql).Scan(&schema)

	if err != nil {
		log.Fatalf("query schema error %s",err)
		return  "",err
	}

	return  schema,nil
}

func Run()  {
	schema,err := QuerySchemaName()
	if err != nil {
		log.Fatal("init db error :%s",err)
		return
	}
	log.Printf("schema name :%s",schema)

	 table,err := QueryTablesList(schema)

	if err != nil {
		log.Fatal("QueryTablesList error :%s",err)
		return
	}

	 for _,v := range table{
	 	list :=QueryTableColumInfo(v.TableName,schema)
	 	fmt.Printf("column :%d \n",len(*list))

	 	err = HandTemplate(*list)

	 }



}

func QueryTablesList(schema string)([]SchemaInfo,error) {
	sql := "SELECT DISTINCT `table_name`,`table_schema` FROM information_schema.`COLUMNS` WHERE table_schema =?"
	var list []SchemaInfo

	err := dbx.Select(&list,sql,schema)
	if err != nil {
		log.Fatalf("query schema error %s",err)
		return  list,err
	}

	if list != nil {
		for _,v := range list{
			log.Printf("schema: %s table; %s",v.SchemaName,v.TableName)
		}

		return  list,err
	}

	log.Fatalf("empty schema %s",schema)

	return  list,fmt.Errorf("empty table list %s",schema)


}