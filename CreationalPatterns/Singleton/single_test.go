package Singleton

import (
	"sync"
	"testing"
)

func TestSingle(t *testing.T) {
	cnt := 30
	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := 0; i < 30; i++ {
		go func() {
			getInstanceByOnce()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSingleWithComment(t *testing.T) {
	cnt := 30
	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := 0; i < 30; i++ {
		go func() {
			getInstanceWithComment()
			wg.Done()
		}()
	}
	wg.Wait()
}
