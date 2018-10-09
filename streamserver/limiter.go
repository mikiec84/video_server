package main

import (
	"log"
	"fmt"
)

type ConnLimiter struct {
	concurrentConn int 
	bucket chan int 
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter {
		concurrentConn: cc,
		bucket: make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	fmt.Println("len of bucket ", len(cl.bucket))
	if len(cl.bucket) >= cl.concurrentConn {
		log.Print("Reached the rate limitation")
		return false 
	}

	cl.bucket <- 1
	return true 
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <- cl.bucket
	log.Printf("New connection coming: %d", c)
}



