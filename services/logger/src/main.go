/*
AIctivity Data Logger Service
This service is used to save data from the Arduino Nano 33 IoT to a SQLite
database. It exposes a single Websocket endpoint at /record.

###################################################################################

main.go
This is the main entrypoint of the service. It starts the service and handles
incoming data via a websocket.

###################################################################################

MIT License

Copyright (c) 2020 Bruno Hautzenberger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func stringToFloat(s string) float32 {
	if f, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(f)
	}
	log.Printf("Could not parse %v to float\n", s)
	return 0
}

func save(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		log.Printf("recv: %s", message)

		saver := func(message []byte) {
			var aX, aY, aZ, gX, gY, gZ float32
			values := strings.Split(string(message), ",")

			if len(values) != 6 {
				log.Println("Not enough data")
				return
			}

			aX = stringToFloat(values[0])
			aY = stringToFloat(values[1])
			aZ = stringToFloat(values[2])
			gX = stringToFloat(values[3])
			gY = stringToFloat(values[4])
			gZ = stringToFloat(values[5])

			var storage SQLiteStorage
			storage.Initialize(os.Getenv("SQLDBPATH"))

			err = storage.SaveData(aX, aY, aZ, gX, gY, gZ)
			if err != nil {
				log.Println("db write error:", err)
				return
			}
		}

		saver(message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/record", save)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil))
}
