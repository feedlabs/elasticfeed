package event

import (
	"github.com/astaxie/beego/toolbox"

	"github.com/feedlabs/elasticfeed/elasticfeed/model"
)

const (
	EVENT_STORING = "storing"
	EVENT_PROCESSING = "processing"
	EVENT_DISTRIBUTING = "distributing"
	EVENT_LEARNING = "learning"

	EVENT_STORING_CREATE_ENTRY = "create-entry"
	EVENT_STORING_CREATE_VIEWER = "create-viewer"
	EVENT_PROCESSING_FEED_MAINTAINER = "feed-maintainer"
	EVENT_PROCESSING_SENSOR_UPDATE = "sensor-update"
	EVENT_DISTRIBUTING_PUSH_ENTRY = "push-entry"
	EVENT_LEARNING_CREATE_METRIC = "create-metric"
)

/**

	- COULD DEFINE EVENTS
	- COULD TRIGGER ON BIND-ED LISTENERS

	- COULD DEFINE ALARM CLOCK
	- COULD DEFINE INTERRUPTS

	EVENT
	- SHOULD HAVE DATA/CALLBACK
	- SHOULD HAVE TYPE
	- SHOULD HAVE PARENT
	- SHOULD HAVE DESTINATION

 */

type EventManager struct {
	engine model.Elasticfeed
	events map[string]interface{}
}

func (this *EventManager) Init() {
	toolbox.StartTask()
//	defer toolbox.StopTask()
}

func (this *EventManager) On(name string, callback func(event *Event)) {
	this.events[name] = callback
}

func (this *EventManager) Off(event string) {
	delete(this.events, event)
}

func (this *EventManager) Trigger(name string, data interface{}) {
	e := NewEvent(data)
	for _, i := range (this.events) {
		i.(func(*Event))(e)
	}
}

func (this *EventManager) GetEventsMap() map[string]interface{} {
	return map[string]interface{}{
		EVENT_STORING: []string{EVENT_STORING_CREATE_ENTRY, EVENT_STORING_CREATE_VIEWER},
		EVENT_PROCESSING: []string{EVENT_PROCESSING_FEED_MAINTAINER, EVENT_PROCESSING_SENSOR_UPDATE},
		EVENT_DISTRIBUTING: []string{EVENT_DISTRIBUTING_PUSH_ENTRY},
		EVENT_LEARNING: []string{EVENT_LEARNING_CREATE_METRIC},
	}
}

func (this *EventManager) InstallSchedule(name string, spec string, cb func() error) error {
	t := toolbox.NewTask(name, spec, cb)

	toolbox.AddTask(name, t)

	return nil
}

func NewEventManager(engine model.Elasticfeed) model.EventManager {
	e := make(map[string]interface{})
	return &EventManager{engine, e}
}
