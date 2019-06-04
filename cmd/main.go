package main

import (
	"git.ucloudadmin.com/securityhouse/mysqlgogenerator/app"
	"git.ucloudadmin.com/securityhouse/mysqlgogenerator/config"
	"log"
)

import "fmt"

func main() {
	fmt.Print("mysql go generator")
	fmt.Printf("mysql addr : %s \n",config.Conf.MysqlAddr)

	err := app.InitDatabasesConnect(config.Conf.MysqlAddr)
	if err != nil {
		log.Fatal("init db error :%s",err)
		return
	}

	app.Run()



}
