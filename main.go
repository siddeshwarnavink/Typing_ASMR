package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	hook "github.com/robotn/gohook"
)

func playSound() {
	fileBytes, err := os.ReadFile("./sounds/keypress-1.mp3")
	if err != nil {
		panic("Sound files are missing")
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}

	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}

func main() {
	fmt.Println("Coding with ASMR is running. Press 'SHIFT+Q' to stop.")

	hook.Register(hook.KeyDown, []string{"q", "shift"}, func(e hook.Event) {
		fmt.Println("Goodbye")
		hook.End()
	})

	hook.Register(hook.KeyUp, []string{"A"}, func(e hook.Event) {
		playSound()
	})

	s := hook.Start()
	<-hook.Process(s)
}
