# test_go_mqtt_sqlite

## Development

### Run app

```bash
make

./bin/dbcreate

docker-compose up-d

./bin/broker

mosquitto_pub -t test/topic -m "{\"message\": \"test\", \"name\": \"mike\"}" -u user -P password

```