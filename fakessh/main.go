package main

import (
	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		for { s.Write([]byte{0x20}) }
	})

	ssh.ListenAndServe(":2222", nil)
 }
