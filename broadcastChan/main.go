package main

// Broadcast channel
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Original code modified and enhanced
//
// See https://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/

import (
	"time"

	"github.com/framps/golang_gotchas/boadcastChan/broadcast"
	"github.com/framps/golang_gotchas/boadcastChan/utils"
)

var b = broadcast.NewBroadcaster()

func listen(id int, r broadcast.Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
		// go listen(id, r)
		utils.Debugf("Rcv - %d: %d\n", id, v)
	}
}

func main() {
	r := b.Listen()
	go listen(1, r)
	go listen(2, r)
	for i := 0; i < 3; i++ {
		utils.Debugf("Broadcasting %d\n", i)
		b.Write(i)
	}
	utils.Debugln("Broadcasting nil")
	b.Write(nil)

	utils.Debugln("Waiting")
	time.Sleep(time.Second * 3)
}
