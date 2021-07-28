package safe

type WorkerGroup struct {
	work func()
	num  int
}

func NewWorkerGroup(work func(), num int) WorkerGroup {
	return WorkerGroup{
		work: work,
		num:  num,
	}
}

func (w WorkerGroup) Start() {
	group := NewRoutineGroup()
	for i := 0; i < w.num; i++ {
		group.Run(w.work)
	}
	group.Wait()
}
