package task

import (
	"context"
	"fmt"
	"testing"
)

func TestConcurrentRun(t *testing.T) {
	runFunc := func(ctx context.Context, a interface{}) (interface{}, error) {
		return fmt.Sprintf("Patch Test: %v", a), nil
	}

	ctx := context.Background()

	var list []interface{}
	for i := 0; i < 200; i++ {
		list = append(list, i+1)
	}

	resultList := ConcurrentRun(ctx, runFunc, 3, list...)
	for _, r := range resultList {
		fmt.Printf("args: %v, result: %v, err: %v\n", r.Args, r.Result, r.Error)
	}
}

