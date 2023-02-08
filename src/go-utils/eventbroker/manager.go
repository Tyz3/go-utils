package eventbroker

var manager EventManager

type EventManager interface {
	RegisterListener(e Event, handler func(e Event), priority Priority, isAsync bool)
}

type eventManager struct {
}

func GetEventManager() EventManager {
	if manager == nil {
		manager = &eventManager{}
	}
	return manager
}

func (m *eventManager) RegisterListener(e Event, handler func(e Event), priority Priority, isAsync bool) {
	e.registerListener(&Listener{
		Handler: handler,
		IsAsync: isAsync,
	}, priority)
}
