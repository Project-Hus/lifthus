package helper

import "sync"

func WaitGroupWaiting(wait int) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(wait)
	return wg
}
