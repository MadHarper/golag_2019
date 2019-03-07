package daemon

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	Connect string `yaml:"connect"`
}

func Initialize() *Conf {
	conf := Conf{}
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &conf
}
