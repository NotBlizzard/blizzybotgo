package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	bot := &PSBot{}
	f, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal([]byte(f), &bot)
	bot.Connect()
}
