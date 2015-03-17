package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type PSBot struct {
	User   string
	Pass   string
	Server string
	Symbol string
	Rooms  []string
	WS     *websocket.Conn
}

type Data struct {
}

func (s PSBot) Send(msg string) {
	if _, err := s.WS.Write([]byte(msg)); err != nil {
		panic(err)
	}
}

func (s PSBot) Command(msg string) string {
	i := strings.Index(msg, " ")
	msg = strings.Replace(msg, " ", "|", -1)
	msg = strings.TrimSpace(msg)
	var cmd string
	if i != -1 {
		cmd = msg[1:i]
	} else {
		cmd = msg[1:]
	}
	if _, ok := Commands[cmd]; ok {
		args := "" //strings.Split(msg, cmd+" ")[1]
		return Commands[cmd](args)
	}
	return msg
}

func (s PSBot) Connect() {
	base := "http://play.pokemonshowdown.com/action.php"
	ws_url := "ws://" + s.Server + "/showdown/websocket"
	ws_origin := "http://play.pokemonshowdown.com"
	ws, err := websocket.Dial(ws_url, "", ws_origin)
	if err != nil {
		panic(err)
	}
	for {
		var message = make([]byte, 512)
		var n int
		n, _ = ws.Read(message)
		s.WS = ws

		str_msg := strings.Replace(string(message), "\n", "", -1)
		str_msg = strings.TrimPrefix(string(str_msg), ">")
		message = []byte(str_msg)
		msg := strings.Split(string(message[:n]), "|")
		fmt.Println(string(message[:n]))
		if strings.Contains(string(message[:n]), "|") {
			switch msg[1] {
			case "challstr":
				var response *http.Response
				var err error
				var data []byte
				var data_str string

				key := msg[2]
				challenge := msg[3]
				if s.Pass == "" {
					response, err = http.Get(base + "?act=getassertion&userid=" + s.User + "&challengekeyid=" + key + "&challenge=" + challenge)
					if err != nil {
						panic(err)
					}
					data, _ = ioutil.ReadAll(response.Body)
				} else {
					response, err = http.PostForm(base, url.Values{
						"act":            {"login"},
						"name":           {s.User},
						"pass":           {s.Pass},
						"challengekeyid": {key},
						"challenge":      {challenge},
					})
					if err != nil {
						panic(err)
					}
					data, _ = ioutil.ReadAll(response.Body)
					data_str = strings.TrimPrefix(string(data), "]")
					data_str = strings.Split(data_str, "\"assertion\":\"")[1]
					data_str = strings.Split(data_str, "\"")[0]
				}

				defer response.Body.Close()
				s.Send("|/trn " + s.User + ",0," + string(data_str))

				for _, e := range s.Rooms {
					s.Send("|/join " + e)
				}
			case "c:":
				room := msg[0]
				if strings.HasPrefix(msg[4], s.Symbol) {
					s.Send(room + "|" + s.Command(msg[4]))
					fmt.Println(room + "|" + s.Command(strings.TrimSpace(msg[4])))
				}
			}
		}
	}
}
