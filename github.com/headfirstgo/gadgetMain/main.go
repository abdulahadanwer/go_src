package main

import (
	"fmt"

	"github.com/headfirstgo/gadget"
)

// Player interface
type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func tryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a tape recorder!")
	}

}
func main() {
	/* 	mixTape := []string{"Jessie's girl", "whip it", "9 to 5"}
	   	var player Player = gadget.TapePlayer{}
	   	playList(player, mixTape)
	   	player = gadget.TapeRecorder{}
		   playList(player, mixTape) */
	tryOut(gadget.TapeRecorder{})
	tryOut(gadget.TapePlayer{})
}
