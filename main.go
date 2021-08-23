package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type TestHandler struct {
	str string
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("HandleFunc")
	w.Write([]byte(string("HandleFunc")))
	go func() {
		fmt.Println("start")
		time.Sleep(5 * time.Second)
		fmt.Println("end")
	}()
}

//ServeHTTP方法，绑定TestHandler
func (th *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle")
	w.Write([]byte(string("Handle")))
}

func main() {
	http.Handle("/", &TestHandler{"Hi"}) //根路由
	http.HandleFunc("/t", SayHello)      //test路由
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
