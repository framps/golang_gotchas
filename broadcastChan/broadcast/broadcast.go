package broadcast

import "github.com/framps/golang_gotchas/boadcastChan/utils"

type broadcast struct {
	c chan broadcast
	v interface{}
}

type broadcastChannel chan (chan broadcast)

type Broadcaster struct {
	// private fields:
	Listenc chan broadcastChannel
	Sendc   chan<- interface{}
}

type Receiver struct {
	// private fields:
	C chan broadcast
}

// create a new broadcaster object.
func NewBroadcaster() Broadcaster {
	listenc := make(chan broadcastChannel)
	sendc := make(chan interface{})
	go func() {
		utils.Debugln("Starting broadcaster gofunc")
		currc := make(chan broadcast, 1)
		for {
			utils.Debugln("Waiting for send and listen channel")
			select {
			case v := <-sendc:
				utils.Debugln("Send received")
				if v == nil {
					currc <- broadcast{}
					return
				}
				c := make(chan broadcast, 1)
				b := broadcast{c: c, v: v}
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

// start listening to the broadcasts.
func (b Broadcaster) Listen() Receiver {
	utils.Debugln("Listening")
	c := make(broadcastChannel, 0)
	b.Listenc <- c
	utils.Debugln("Returning")
	return Receiver{<-c}
}

// broadcast a value to all listeners.
func (b Broadcaster) Write(v interface{}) {
	utils.Debugf("Sending %v\n", v)
	b.Sendc <- v
}

// read a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
	utils.Debugln("Reading")
	b := <-r.C
	v := b.v
	r.C <- b
	r.C = b.c
	return v
}
