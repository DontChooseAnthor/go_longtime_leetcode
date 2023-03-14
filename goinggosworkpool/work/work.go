package work

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	addRoutine = 1
	rmvRoutine = 2
)

var ErrorInvalidMinRoutines = errors.New("Invalid minimum numbrer of routines")

var ErrorInvalidStatTime = errors.New("Invalid duration for stat time")

type Worker interface {
	Work(id int)
}

type Pool struct {
	minRoutine int
	statTime   time.Duration
	counter    int
	tasks      chan Worker
	control    chan int // 记录worker的数量
	kill       chan struct{}
	shutdown   chan struct{}
	wg         sync.WaitGroup
	routines   int64 
	active     int64  // atomic记录活跃worker数量
	pending    int64
	logFunc    func(message string)
}

func (p *Pool) Add(routines int) {
	if routines == 0 {
		return
	}

	cmd := addRoutine

	if routines < 0 {
		routines = routines * -1
		cmd = routines
	}

	for i := 0; i < routines; i++ {
		p.control <- cmd
	}
}

func (p *Pool) work(id int) {
done:
	for {
		select {
		case t := <-p.tasks:
			atomic.AddInt64(&p.active, 1)
			{
				t.Work(id)
			}
		case <-p.kill:
			break done
		}
	}

	atomic.AddInt64(&p.routines, -1)
	p.wg.Done()

	p.log("worker: shut down")
}

func (p *Pool) Run(work Worker) {
	atomic.AddInt64(&p.pending, 1)
	{
		p.tasks <- work
	}
	atomic.AddInt64(&p.pending, -1)
}

func (p *Pool) Shutdown() {
	close(p.shutdown)
	p.wg.Wait()
}

func (p *Pool) log(message string) {
	if p.logFunc != nil {
		p.logFunc(message)
	}
}

func (p *Pool) manager() {
	p.wg.Add(1)

	go func() {
		p.log("work manager: started")
		timer := time.NewTimer(p.statTime)

		for {
			select {
			case <-p.shutdown:
				routines := int(atomic.LoadInt64(&p.routines))
				for i := 0; i < routines; i++ {
					p.kill <- struct{}{}
				}

				p.wg.Done()
				return
			case c := <-p.control:
				switch c {
				case addRoutine:
					p.log("work manager: add routine")
					p.counter++
					p.wg.Add(1)
					atomic.AddInt64(&p.routines, 1)
					go p.work(p.counter)
				case rmvRoutine:
					p.log("work manager: remove routine")
					routines := int(atomic.LoadInt64(&p.routines))
					if routines <= p.minRoutine {
						p.log("work mana")
					}
				}
			case <-timer.C:
				routines := atomic.LoadInt64(&p.routines)
				pending := atomic.LoadInt64(&p.pending)
				active := atomic.LoadInt64(&p.active)
				p.log(fmt.Sprintf("work manager: stats : G[%d] P[%d] A[%d]", routines, pending, active))
				timer.Reset(p.statTime)
			}
		}
	}()
}

func New(minRoutine int, statTime time.Duration, logFunc func(message string)) (*Pool, error) {
	if minRoutine <= 0 {
		return nil, ErrorInvalidMinRoutines
	}
	if statTime < time.Millisecond {
		return nil, ErrorInvalidStatTime
	}
	p := Pool{
		minRoutine: minRoutine,
		statTime:   statTime,
		tasks:      make(chan Worker),
		control:    make(chan int),
		kill:       make(chan struct{}),
		shutdown:   make(chan struct{}),
		logFunc:    logFunc,
	}
	p.manager()
	p.Add(minRoutine)
	return &p, nil
}
