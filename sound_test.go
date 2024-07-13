package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// When the user is typing fast, it should select different sound file
func TestReturnDifferentFile(t *testing.T) {
	myTime := time.Now()
	var prevSoundFile string

	file := GetSoundFile(69, 100, &myTime, &prevSoundFile)
	time.Sleep(25 * time.Millisecond)

	file1 := GetSoundFile(69, 100, &myTime, &prevSoundFile)
	file2 := GetSoundFile(69, 100, &myTime, &prevSoundFile)

	emptyMsg := "file should not be empty"

	assert.NotEqual(t, file, "", emptyMsg)
	assert.NotEqual(t, file1, "", emptyMsg)
	assert.NotEqual(t, file2, "", emptyMsg)

	message := "file name should be different when user is typing continuously"

	assert.NotEqual(t, file, file1, message)
	assert.NotEqual(t, file1, file2, message)
}

// When the user is typing slow, we should select the same sound file
func TestReturnSameFile(t *testing.T) {
	myTime := time.Now()
	var prevSoundFile string

	file := GetSoundFile(69, 100, &myTime, &prevSoundFile)
	time.Sleep(250 * time.Millisecond)

	file1 := GetSoundFile(69, 100, &myTime, &prevSoundFile)
	time.Sleep(200 * time.Millisecond)

	file2 := GetSoundFile(69, 100, &myTime, &prevSoundFile)

	emptyMsg := "file should not be empty"

	assert.NotEqual(t, file, "", emptyMsg)
	assert.NotEqual(t, file1, "", emptyMsg)
	assert.NotEqual(t, file2, "", emptyMsg)

	message := "file name should be the same when user is pausing while typing"

	assert.Equal(t, file, file1, message)
	assert.Equal(t, file1, file2, message)
}
