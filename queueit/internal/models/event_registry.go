package models

import (
	"encoding/json"
	"io/ioutil"
)

type EventRegistry struct {
	events map[string]Event
}

func NewEventRegistry() *EventRegistry {
	registry := &EventRegistry{events: make(map[string]Event)}

	file, _ := ioutil.ReadFile("./config/events.json")

	var data []Event

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data); i++ {
		registry.events[data[i].GetId()] = data[i]
	}

	return registry
}

func (er EventRegistry) GetEvent(id string) (event Event, success bool) {
	if event, ok := er.events[id]; ok {
		return event, true
	}
	return Event{}, false
}
