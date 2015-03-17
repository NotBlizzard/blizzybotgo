package main

import (
	"runtime"
)

var Commands = map[string]func(string) string{
	"about": about,
	"lol":   lol,
}

func about(_ string) string {
	return "BlizzyBotGo: Written in Go " + runtime.Version()
}

func lol(_ string) string {
	return "LOL"
}
