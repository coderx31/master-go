package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
this implementation will help to override the ctx canceled
*/

/**
 - if we want to create a channel to send notifications without data, the appropriate way to do
	so in Go is a chan struct{}
*/

/**
- using unbuffered channel, sender will block until receiver receives the data from the channel
- with a buffered channel, sender can send messages until channel is full.
*/

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}

type Donation struct {
	cond    *sync.Cond
	balance int
}

func main() {

	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()

	//res := getName(ctx)
	//fmt.Println(res)

	//res := getName(detach{ctx: ctx})
	//fmt.Println(res)

	// nil channel

	//fmt.Println("starting")
	//var ch chan int
	//<-ch
	//ch <- 0
	//fmt.Println("ending")

	//wg := sync.WaitGroup{}
	//var v uint64
	//
	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go func() {
	//		//wg.Add(1)
	//		atomic.AddUint64(&v, 1)
	//		wg.Done()
	//	}()
	//}
	//
	//wg.Wait()
	//
	//fmt.Println(v)

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// listener goroutine
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go f(10)
	go f(15)

	// updater goroutine
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}

func getName(ctx context.Context) string {
	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("context canceled")
		return ""
	default:
		return "shenal"
	}
}
