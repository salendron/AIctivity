# AIctivity Data Logger Service
This service is used to save data from the Arduino Nano 33 IoT to a SQLite database. It exposes a single Websocket endpoint at /record.

## Docker deployment
### Build Container
```
docker build --tag aictivity-logger:1.0 .
```

### Run Container
This command starts the service on port 9000 and mount local directory /home/pi/data to /data in the container. It will save the SQLite Database to /date/aictivity.db.
```
docker run -d -p 9000:9000 --name aictivity-logger -e PORT='9000' -e SQLDBPATH='/data/aictivity.db' -v /var/run/docker.sock:/var/run/docker.sock --restart unless-stopped --mount type=bind,source=/home/pi/data,target=/data aictivity-logger:1.0
```
