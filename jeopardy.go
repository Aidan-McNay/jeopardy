package main

import (
	"fmt"
	"logic"
)

func main() {
	player := logic.Player{}
	player.SetName("Joey")
	player.ResetScore()
	player.IncrScore(500)

	fmt.Println(player.AsString())
}
