# TODO

# Build OpenAPI server
 ./oapi-codegen /home/ubuntu/git/AIctivity/services/logger/openapi.json > /home/ubuntu/git/AIctivity/services/logger/src/server.gen.go

# Testcall
 curl --header "Content-Type: application/json" --request POST --data '{"aX":1, "aY":2, "aZ":3, "gX":4, "gY":5, "gZ":6, "temp":1.2}' http://localhost:9000/data


 # Desploy container
 docker run -d -p 9000:9000 --name aictivity-logger -e PORT='9000' -e SQLDBPATH='/data' -v /var/run/docker.sock:/var/run/docker.sock --restart unless-stopped --mount type=bind,source=/home/pi/data,target=/data aictivitiy-logger:1.0
