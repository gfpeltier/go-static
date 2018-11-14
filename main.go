package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"grantpeltier.com/internal"
)

const CFG_FILE string = "config.yml"

type Config struct {
	Port string
	Theme string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	dat, err := ioutil.ReadFile(CFG_FILE)
	check(err)
	
	t := Config{}
	err = yaml.Unmarshal([]byte(dat), &t)
	check(err)

	internal.BuildSite(t.Theme)
}