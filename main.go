package main

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"lnj.com/unix/sockets/message"
)

const (
	protocol = "unix"
	sockAddr = "/temp/stage-one-socket"
)

type InputMessage struct {
	Id          int64  `json:"id"`
	ShortCode   string `json:"shortcode"`
	Destination string `json:"destination"`
	Tag         string `json:"tag"`
}

func main() {
	var values []InputMessage

	inMsg := InputMessage{Id: 0, ShortCode: "KV231wx", Destination: "http://att.com", Tag: "CallAtt"}
	values = append(values, inMsg)

	inMsg = InputMessage{Id: 0, ShortCode: "AB43yyre", Destination: "http://somewhereoutthere.com", Tag: "Far out man"}
	values = append(values, inMsg)

	inMsg = InputMessage{Id: 0, ShortCode: "XYZZ321", Destination: "http://google.com", Tag: "Google Baby"}
	values = append(values, inMsg)

	inMsg = InputMessage{Id: 0, Destination: "http://huntington.com", Tag: "Huntington"}
	values = append(values, inMsg)

	inMsg = InputMessage{Id: 0, ShortCode: "abcd1234", Tag: "Dumb stuff"}
	values = append(values, inMsg)

	var idCount int64 = 1
	for idx := 0; idx < 500; idx++ {

		for _, d := range values {
			d.Id = idCount
			idCount++
			time.Sleep(time.Millisecond * 5)
			v, _ := json.Marshal(d)

			conn, err := net.Dial(protocol, sockAddr)
			if err != nil {
				log.Fatal(err)
			}

			func() {
				defer conn.Close()

				m := &message.Transport{
					Length: len(v),
					Data:   []byte(v),
				}

				err = m.Write(conn)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
	}
}
