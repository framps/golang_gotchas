package broadcast

import "github.com/framps/golang_gotchas/broadcastChan/utils"

type broadcast struct {
	c chan broadcast
	v interface{}
}

type broadcastChannel chan (chan broadcast)

// Broadcaster -
type Broadcaster struct {
	// private fields:
	Listenc chan broadcastChannel // listenerchannel
	Sendc   chan<- interface{}    // send channel
}

// Receiver -
type Receiver struct {
	// private fields:
	C chan broadcast // a receiver has a broadcast channel
}

// NewBroadcaster - create a new broadcaster object.
func NewBroadcaster() Broadcaster {
	listenc := make(chan broadcastChannel) // listenerchannel which receives a broadcast channel
	sendc := make(chan interface{})        // send channel
	go func() {
		utils.Debugln("Starting broadcaster gofunc")
		currc := make(chan broadcast, 1)
		for {
			utils.Debugln("Waiting for send and listen channel")
			select {
			case v := <-sendc:
				utils.Debugln("Send received")
				if v == nil {
					utils.Debugln("Nil received")
					currc <- broadcast{}
					return
				}
				c := make(chan broadcast, 1) // create new broadcast channel
				b := broadcast{c: c, v: v}   // insert broadcast channel and value in new broadcast
				currc <- b
				currc = c
			case r := <-listenc:
				utils.Debugln("Listen received")
				r <- currc
			}
		}
	}()
	utils.Debugf("Returning Broadcaster with listener channel and send channel")
	return Broadcaster{
		Listenc: listenc,
		Sendc:   sendc,
	}
}

// Listen - start listening to the broadcasts.
func (b Broadcaster) Listen() Receiver {
	utils.Debugln("Listening")
	c := make(broadcastChannel, 0) // create a broadcastchannel
	b.Listenc <- c                 // write new broadcastchannel on listener channel
	utils.Debugln("Returning")
	return Receiver{<-c} // return receiver when broadcastchannel received
}

// Write - broadcast a value to all listeners.
func (b Broadcaster) Write(v interface{}) {
	utils.Debugf("Sending %v\n", v)
	b.Sendc <- v // write value on send channel
}

// read a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
	utils.Debugln("Reading")
	b := <-r.C // wait for a broadast channel
	v := b.v   // retrieve value from received broadcastchannel
	r.C <- b   // write same broadcastchannel to broadcastchannel
	r.C = b.c  // broadcastchannel now becomes bc from broadcast
	return v   // return received value
}
