package proxy

import "log"

type Subject interface {
	Send() string
}

type Proxy struct {
	realSubject Subject
}

func (p *Proxy) Send() string {
	if p.realSubject == nil {
		log.Println("creating real subject")

		p.realSubject = &RealSubject{}
	}

	return p.realSubject.Send()
}

type RealSubject struct{}

func (s *RealSubject) Send() string {
	return "I’ll be back!"
}
