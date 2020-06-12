package config

import (
	"os"
	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/toml"
)

type Path struct{
	file string
	response int
}

type Server struct{
	Port string
	Paths []Path
}

func Read(conffile string) Server {
	var conf Server
	_, err := toml.DecodeFile(conffile, &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	return conf
}
