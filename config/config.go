package config

import "github.com/koding/multiconfig"

type ConfigStruct struct {
	MysqlAddr string `required:"true"`
	DBTag string
}

var Conf ConfigStruct

func init()  {
	m := multiconfig.NewWithPath("config.toml")
	m.MustLoad(&Conf)
}
