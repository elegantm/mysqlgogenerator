package app

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

func HandTemplate(list[] TableInfo) error {

	t ,_:= template.ParseFiles("query.tpl")
	
	var table TableList
	
	if len(list)<1 {
		log.Fatal("empty tabInfo")
		return nil
	}
	
	column := make([]Column,0)
	var temp Column
	table.TableName = list[0].TableName
	table.StrutName = ConvertToCamel(table.TableName)

	listLen := len(list)

	for k,v := range list{
		temp.ColumnName = v.ColumnName
		temp.DataType =  ConvertMysqlType(v.DataType)
		temp.Name =  ConvertToCamel(v.ColumnName)
		temp.ColumnComment = v.ColumnComment


		if k == (listLen-1) {
			temp.Separator = ""
		}else {
			temp.Separator = ","
		}

		column =append(column,temp)

	}
	
	table .Fields = column

	fmt.Printf("table dump: %+v \n",table)

	fileName := fmt.Sprintf("out/%s.go",table.TableName)
	 //f,_:= os.OpenFile(fileName,os.O_APPEND | os.O_WRONLY,0666)
	 //defer f.Close()

	 f,err1 := os.Create(fileName)
	if err1 != nil {
		log.Fatalf("file  Execute error :%s",err1)
		return  err1
	}

	 buf := new(bytes.Buffer)

	err := t.Execute(buf,table)

	if err != nil {
		log.Fatalf("template  Execute error :%s",err)
		return  err
	}

	log.Printf("buf: %v",buf)

  n,err :=  f.Write(buf.Bytes())
	log.Printf("buf n: %v",n)
	if err != nil {
		log.Fatalf("wrrite file error :%s",err)
		return  err
	}
	return nil
	
}