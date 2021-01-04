# TODO

# Build OpenAPI server
 ./oapi-codegen /home/ubuntu/git/AIctivity/services/logger/openapi.json > /home/ubuntu/git/AIctivity/services/logger/src/server.gen.go

# Testcall
 curl --header "Content-Type: application/json" --request POST --data '{"aX":1, "aY":2, "aZ":3, "gX":4, "gY":5, "gZ":6, "temp":1.2}' http://localhost:9000/data
