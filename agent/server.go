package agent

type Agent struct {
	stop chan struct{}
}

func NewAgent() *Agent {
	return &Agent{
		make(chan struct{}),
	}
}

func (a *Agent) Start() {
	go func() {

	}()
}
