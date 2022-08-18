package notifiers

import "DummyAlerts/messages"

type Notifiers interface {
	Notify(*messages.Message) error
}
