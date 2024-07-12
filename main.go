package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	hook "github.com/robotn/gohook"
)

var otoCtx *oto.Context

func initializeOtoContext() {
	op := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	var err error
	var readyChan <-chan struct{}
	otoCtx, readyChan, err = oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}

	<-readyChan
}

func playSound(e hook.Event) {
	var soundFile string

	switch e.Keycode {
	case 57:
		soundFile = "./sounds/space.mp3"
	case 28:
		soundFile = "./sounds/enter.mp3"
	default:
		randomNumber := rand.Intn(3) + 1
		soundFile = fmt.Sprintf("./sounds/sound-%d.mp3", randomNumber)
	}

	fileBytes, err := os.ReadFile(soundFile)
	if err != nil {
		panic("Sound files are missing")
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}

func main() {
	initializeOtoContext()

	fmt.Println("Coding with ASMR is running. Press 'SHIFT+Q' to stop.")

	hook.Register(hook.KeyDown, []string{"q", "shift"}, func(e hook.Event) {
		fmt.Println("Goodbye")
		hook.End()
	})

	hook.Register(hook.KeyUp, []string{"A"}, func(e hook.Event) {
		go playSound(e)
	})

	s := hook.Start()
	<-hook.Process(s)
}
