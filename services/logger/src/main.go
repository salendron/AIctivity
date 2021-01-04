/*
LOGGER SERVICE

This service is used to simply record the data gathered by the arduino hardware on
the user's body to a SQLite Database.
This Data will later on be used to train our model.

###################################################################################

main.go
This is the main entrypoint of the service. It starts the service and
starts the reverse proxy.

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

 ./oapi-codegen /home/ubuntu/git/AIctivity/services/logger/openapi.json > /home/ubuntu/git/AIctivity/services/logger/src/server.gen.go


 curl --header "Content-Type: application/json" --request POST --data '{"aX":1, "aY":2, "aZ":3, "gX":4, "gY":5, "gZ":6, "temp":1.2}' http://localhost:9000/data
*/

package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	var storage SQLiteStorage
	storage.Initialize(os.Getenv("SQLDBPATH"))

	var api API
	api.SetStorage(&storage)

	e := echo.New()
	RegisterHandlers(e, &api)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
