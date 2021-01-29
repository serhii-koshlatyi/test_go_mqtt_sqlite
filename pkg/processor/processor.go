package processor

import (
	"fmt"

	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/model"
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/store"
)

type Processor interface {
	Run(in <-chan model.MQTTMessage)
}

type StateProcessor struct {
	store store.Store
}

func NewStateProcessor(store store.Store) *StateProcessor {
	proc := &StateProcessor{
		store: store,
	}

	return proc
}

// Run runs message processor loop
func (p *StateProcessor) Run(messages <-chan model.MQTTMessage) {
	for msg := range messages {

		err := p.store.Create(msg)
		if err != nil {
			fmt.Printf("can`t save in db: %v", err)
			continue
		}

		fmt.Println("Data saved in db")
	}
}
