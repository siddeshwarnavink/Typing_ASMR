package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("Coding with ASMR is running. Press 'SHIFT+Q' to stop.")

	hook.Register(hook.KeyDown, []string{"q", "shift"}, func(e hook.Event) {
		fmt.Println("Goodbye")
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("pressed w!!")
	})

	s := hook.Start()
	<-hook.Process(s)
}
