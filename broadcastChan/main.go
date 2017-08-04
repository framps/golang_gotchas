package main

// Broadcast channel
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Original code modified and enhanced
//
// See https://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/

import (
	"strconv"
	"time"

	"github.com/framps/golang_gotchas/broadcastChan/broadcast"
	"github.com/framps/golang_gotchas/broadcastChan/utils"
)

// Demo broadcaster
var b = broadcast.NewBroadcaster()

// Demo listener
func listen(id int, r broadcast.Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
		// go listen(id, r)
		utils.Debugf("Listener %d: Received: '%v'\n", id, v)
	}
}

func main() {

	// start broadcaster to listen for any broadcasts
	r := b.Listen()
	// start two demo listener
	go listen(1, r)
	go listen(2, r)
	// now broadcast
	for i := 0; i < 3; i++ {
		utils.Debugf("Broadcasting %d\n", i)
		b.Write("Broadcast " + strconv.Itoa(i))
	}

	utils.Debugln("Broadcasting nil")
	b.Write(nil)

	utils.Debugln("Waiting for broadcasting to finish")
	time.Sleep(time.Second * 3)
}
