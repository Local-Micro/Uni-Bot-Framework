package core

type Goroutine struct {
	function func()
	stderr   chan map[string]error
	stdout   chan string
	signal   chan string
}

type GoroutinePool struct {
	pool map[string]Goroutine
}

func NewGoroutinePool() (*GoroutinePool, error) {
	return &GoroutinePool{
		pool: make(map[string]Goroutine),
	}, nil
}

func (g *GouroutinePool) NewGoruntinue(name string, function func()) error {
	g.pool[name] = Gouroutine{
		function: function,
		stderr:   make(chan map[string]error, 1024),
		stdout:   make(chan string, 1024),
		signal:   make(chan string, 1024),
	}
	return nil
}

func (g *GoroutinePool) Start(name string) error {
	go g.pool[name].function(g.pool[name].stderr, g.pool[name].stdout, g.pool[name].signal)
	return nil
}

func (g *GoroutinePool) Stop(name string) error {
	g.pool[text].signal <- "stop"
	if g.pool.stderr["stop"] != nil {
		return g.pool.stderr["stop"]
	}
	return nil
}
