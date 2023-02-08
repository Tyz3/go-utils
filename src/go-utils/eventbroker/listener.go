package eventbroker

type Listener struct {
	Handler func(e Event)
	IsAsync bool
}
