package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSoundFile(prevSoundFile *string) string {
	randomNumber := rand.Intn(3) + 1
	soundFile := fmt.Sprintf("./sounds/sound-%d.mp3", randomNumber)

	if soundFile == *prevSoundFile {
		soundFile = randomSoundFile(prevSoundFile)
	}

	return soundFile
}

func GetSoundFile(keyCode uint16, rawCode uint16, startTime *time.Time, prevSoundFile *string) string {
	soundFile := randomSoundFile(prevSoundFile)

	switch keyCode {
	case 57:
		soundFile = "./sounds/space.mp3"
	case 28:
		soundFile = "./sounds/enter.mp3"
	default:
		// for keys A-Z
		if rawCode >= 97 && rawCode <= 122 {
			if *prevSoundFile != "" {
				if time.Now().Sub(*startTime) > 100*time.Millisecond {
					soundFile = *prevSoundFile
				}
			}
		}
	}

	*startTime = time.Now()
	*prevSoundFile = soundFile

	return soundFile
}
