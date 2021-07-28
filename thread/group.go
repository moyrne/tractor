package thread

import "sync"

type RoutineGroup struct {
	group sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup {
	return &RoutineGroup{group: sync.WaitGroup{}}
}

func (g *RoutineGroup) Run(fn func()) {
	g.group.Add(1)
	Go(func() {
		defer g.group.Done()
		fn()
	})
}

func (g *RoutineGroup) Wait() {
	g.group.Wait()
}
