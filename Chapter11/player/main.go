package main

import "player/gadget"

func main() {
	player := gadget.TapePlayer{}
	songList := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playlist(player, songList)
}

func playlist(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}
