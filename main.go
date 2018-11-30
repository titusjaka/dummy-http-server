package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const waitMilliseconds = 10000

var randomWait int32

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = fmt.Sprintf("Wait till start: %d", randomWait)

	w.Write([]byte(message))
}

func main() {
	randomWait = waitMilliseconds - rand.Int31n(waitMilliseconds/2)
	time.Sleep(time.Duration(randomWait) * time.Microsecond)
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
