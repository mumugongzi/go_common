package task

import (
	"context"
	"sync"
)

/*
并行运行框架
*/

type ConcurrentTask struct {
	Ctx context.Context

	// 批量查询的股票代码列表
	ArgsList []interface{}

	RunFunc func(context.Context, interface{}) (interface{}, error)

	// 设置并发度
	Concurrent int

	ResultList []*Result

	mutex sync.Mutex
}

type Result struct {
	Args   interface{}
	Result interface{}
	Error  error
}

// concurrent: 设置任务的并发度
func NewConcurrentTask(ctx context.Context, runFunc func(context.Context, interface{}) (interface{}, error), concurrent int, args ...interface{}) *ConcurrentTask {
	return &ConcurrentTask{
		Ctx:        ctx,
		ArgsList:   args,
		RunFunc:    runFunc,
		Concurrent: concurrent,
		ResultList: []*Result{},
		mutex:      sync.Mutex{},
	}
}

func concurrentRun(t *ConcurrentTask) []*Result {
	start, delta := 0, t.Concurrent
	argsList, mutex := t.ArgsList, t.mutex
	for start < len(argsList) {
		if start+delta > len(argsList) {
			delta = len(argsList) - start
		}
		group := sync.WaitGroup{}
		group.Add(delta)
		for i := start; i < start+delta; i++ {
			args := argsList[i]
			go func() {
				result, err := t.RunFunc(t.Ctx, args)
				mutex.Lock()
				t.ResultList = append(t.ResultList, &Result{
					Args:   args,
					Result: result,
					Error:  err,
				})
				mutex.Unlock()
				group.Done()
			}()

		}
		group.Wait()
		start = start + delta
	}
	return t.ResultList
}

// 返回结果具有无序性, 如果要返回有序结果, 将concurrent设置为1
func ConcurrentRun(ctx context.Context, runFunc func(context.Context, interface{}) (interface{}, error), concurrent int, args ...interface{}) []*Result {
	t := NewConcurrentTask(ctx, runFunc, concurrent, args...)
	return concurrentRun(t)
}