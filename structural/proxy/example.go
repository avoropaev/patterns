package proxy

import "log"

func Example() {
	proxy := Proxy{}
	log.Println("proxy created")

	log.Println("call send")
	log.Println(proxy.Send())
}
