package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	waitMilliseconds = 10000
	port             = ":8080"
)

var randomWait int32

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Wait till start: %d", randomWait)

	w.Write([]byte(message))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	randomWait = waitMilliseconds - rand.Int31n(waitMilliseconds/2)
	timeToSleep := time.Duration(randomWait) * time.Millisecond
	fmt.Printf("[%v] Time to sleep: %v\n", time.Now().Format("2006-01-02 15:04:05.000"), timeToSleep)
	time.Sleep(timeToSleep)
	fmt.Printf("[%v] Start listener at port %s\n", time.Now().Format("2006-02-01 15:04:05.000"), port)
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
