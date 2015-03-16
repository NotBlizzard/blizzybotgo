package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"strings"
)

type ShowdownBot struct {
	Name   string
	Server string
	Rooms  []string
	ws     *websocket.Conn
}

func (s ShowdownBot) Send(msg string) {
	if _, err := s.ws.Write([]byte(msg)); err != nil {
		panic(err)
	}
}

func (s ShowdownBot) Connect() {
	base := "http://play.pokemonshowdown.com/action.php"
	ws_url := "ws://" + s.Server + "/showdown/websocket"
	ws_origin := "http://play.pokemonshowdown.com"
	ws, err := websocket.Dial(ws_url, "", ws_origin)
	if err != nil {
		panic(err)
	}
	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			panic(err)
		}
		s.ws = ws
		message := strings.Split(string(msg[:n]), "|")
		fmt.Println(string(msg[:n]))
		if !strings.Contains(string(msg[:n]), "|") {
			//crash
		} else {
			switch message[1] {
			case "challstr":
				key := message[2]
				challenge := message[3]
				url_get := base + "?act=getassertion&userid=blizzybotgo&challengekeyid=" + key + "&challenge=" + challenge
				response, err := http.Get(url_get)
				if err != nil {
				}
				defer response.Body.Close()
				data, err := ioutil.ReadAll(response.Body)
				if err != nil {
					//panic(err)
				}
				trn := "|/trn blizzybotgo,0," + string(data)
				s.Send(trn)
				for _, e := range s.Rooms {
					s.Send("|/join " + e)
				}
			case "c:":
				fmt.Println("msg is " + message[4])
				//
			}
		}

	}
}

func main() {
	b := &ShowdownBot{"Kek", "198.27.117.206:8000", []string{"lobby"}, nil}
	b.Connect()
}
