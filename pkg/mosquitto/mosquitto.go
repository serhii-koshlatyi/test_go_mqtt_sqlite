package mosquitto

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/config"
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/model"
)

type MosquittoBroker interface {
	ResultChan() chan model.MQTTMessage
	Subscribe()
	Unsubscribe()
}

type MosquittoSubBroker struct {
	resultChan chan model.MQTTMessage
	client     mqtt.Client
	topic      string
}

func (b *MosquittoSubBroker) messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	var message model.MQTTMessage

	err := json.Unmarshal(msg.Payload(), &message)
	if err != nil {
		fmt.Printf("can`t unmarshal json: %s", err)
	}

	b.resultChan <- message
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func NewMosquittoBroker(cfg config.Mosquitto) (*MosquittoSubBroker, error) {
	broker := &MosquittoSubBroker{
		resultChan: make(chan model.MQTTMessage),
		topic:      cfg.Topic,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.Broker, cfg.Port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(cfg.User)
	opts.SetPassword(cfg.Password)
	opts.SetDefaultPublishHandler(broker.messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	broker.client = client

	return broker, nil
}

func (b *MosquittoSubBroker) ResultChan() chan model.MQTTMessage {
	return b.resultChan
}

func (b *MosquittoSubBroker) Subscribe() {
	token := b.client.Subscribe(b.topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", b.topic)
}

func (b *MosquittoSubBroker) Unsubscribe() {
	b.client.Unsubscribe(b.topic)
	close(b.resultChan)
	fmt.Printf("Unsubscribed from topic: %s \n", b.topic)
}
