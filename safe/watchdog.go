package safe

import (
	"time"
)

// WatchDog TODO 看门狗
type WatchDog struct {
	done      chan struct{}
	doneErr   chan struct{}
	heartbeat chan struct{}
	timer     time.Duration
	start     Start
	errDo     ErrDo
	runner    Runner
}

type (
	Start  func() (Runner, ErrDo, error)
	ErrDo  func()
	Runner func() error
)

// NewWatchDog create a new object
func NewWatchDog(timer time.Duration, start Start) *WatchDog {
	return &WatchDog{
		done:      make(chan struct{}),
		doneErr:   make(chan struct{}),
		heartbeat: make(chan struct{}),
		timer:     timer,
		start:     start,
	}
}

func (w *WatchDog) HeartBeat() {
	w.heartbeat <- struct{}{}
}

func (w *WatchDog) DoneErr() {
	w.doneErr <- struct{}{}
}

func (w *WatchDog) Start() (err error) {
	w.runner, w.errDo, err = w.start()
	if err != nil {
		return err
	}
	Go(w.Watch)
	defer func() {
		close(w.done)
	}()
	return w.runner()
}

func (w *WatchDog) Watch() {
	timer := time.NewTimer(w.timer)
	window := [3]bool{true, true, true}
	for {
		timer.Reset(w.timer)
		select {
		case <-w.done:
			// 主动退出
			return
		case <-w.doneErr:
			// 错误结束
			w.errDo()
			return
		case <-w.heartbeat:
			window[0], window[1], window[2] = true, true, true
		case <-timer.C:
			// 心跳出错
			window[0], window[1], window[2] = window[1], window[2], false
			if !window[0] && !window[1] && !window[2] {
				w.errDo()
				return
			}
		}
	}
}
