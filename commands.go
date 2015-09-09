package main

import (
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// TODO: More commands.

var Commands = map[string]func(string) string{
	"about": about,
	"pick":  pick,
}

func about(_ string) string {
	return "BlizzyBotGo: Written in Go " + runtime.Version()
}

func pick(choices string) string {
	options := strings.Split(choices, ",")
	rand.Seed(time.Now().Unix())
	choice := rand.Intn(len(choices))
	return "I randomly pick \"" + string(options[choice]) + "\""
}
