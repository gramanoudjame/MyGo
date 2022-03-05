package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const PHILOS = 5
const EATERS = 2
const TIMES = 3

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

func (p Philo) pickUpChopsticks() {
	rand.Seed(time.Now().UnixNano())
	// random pick up order for the chopsticks
	if rand.Intn(2) != 0 {
		p.leftCS.Lock()
		p.rightCS.Lock()
	} else {
		p.rightCS.Lock()
		p.leftCS.Lock()
	}
}

func (p Philo) releaseChopsticks() {
	rand.Seed(time.Now().UnixNano())
	// random release order for the chopsticks
	if rand.Intn(2) != 0 {
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	} else {
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func (p Philo) eat(times int, host Host) {
	for i := 0; i < times; i++ {
		<-host.permissionToEat
		p.pickUpChopsticks()
		fmt.Printf("starting to eat %v\n", p.number)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("finishing eating %v\n", p.number)
		p.releaseChopsticks()
		host.finishedEating <- p.number
	}
}

type Host struct {
	permissionToEat chan int
	finishedEating  chan int
}

func (h Host) feed(concurrentEaters int) {
	for i := 0; i < concurrentEaters; i++ {
		h.permissionToEat <- 1
	}
	for range h.finishedEating {
		h.permissionToEat <- 1
	}
}

func main() {
	// Initialisation
	CSticks := make([]*ChopS, PHILOS)
	for i := 0; i < PHILOS; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, PHILOS)
	for i := 0; i < PHILOS; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%PHILOS]}
	}
	// Start host
	host := Host{make(chan int, EATERS), make(chan int, PHILOS)}
	go host.feed(EATERS)
	// Start eating
	var wg sync.WaitGroup
	for i := 0; i < PHILOS; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			philos[i].eat(TIMES, host)
		}(&wg, i)
	}
	wg.Wait()
}
