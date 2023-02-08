package eventbroker

type Priority int

const (
	LOWEST Priority = iota
	LOW
	MEDIUM
	HIGH
	HIGHEST
	MONITOR
)

var Values = []Priority{LOWEST, LOW, MEDIUM, HIGH, HIGHEST, MONITOR}

type event struct {
	name      string
	listeners map[Priority][]*Listener
}

type Event interface {
	Call()
	GetEventName() string
	registerListener(listener *Listener, priority Priority)
}

func NewEvent(name string) Event {
	return &event{
		name:      name,
		listeners: make(map[Priority][]*Listener, len(Values)),
	}
}

func (e *event) GetEventName() string {
	return e.name
}

func (e *event) registerListener(listener *Listener, priority Priority) {
	list, exists := e.listeners[priority]
	if !exists {
		list = make([]*Listener, 0, 1)
	}

	list = append(list, listener)
	e.listeners[priority] = list
}

func (e *event) Call() {
	for _, p := range Values {
		if listeners, ok := e.listeners[p]; ok {
			for _, l := range listeners {
				if l.IsAsync {
					go l.Handler(e)
				} else {
					l.Handler(e)
				}
			}
		}
	}
}
