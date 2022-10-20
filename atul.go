package main

import (
	"fmt"
	"time"
)

type Player struct {
	Name   string
	Health uint
	Power  int
}

// func NewPlayer(name string, hp uint, power int )Player{
// 	return Player{
// 		Name : name,
// 		Health : hp,
// 		Power : power,
// 	}
// }

func main() {

	// playerA := NewPlayer("green",150,22)
	// playerB := NewPlayer("reyna",111,90)

	playerA := &Player{
		Name:   "green",
		Health: 989,
		Power:  888,
	}

	playerB := Player{
		Name:   "reyna",
		Health: 999,
		Power:  1100,
	}
	fmt.Printf("the name of first player is %s \n", playerA.Name)
	fmt.Println("the name of second player is ", playerB.Name)

	fmt.Println("the first player damage second player by ", playerA.Power)
	fmt.Println("the second player survives with hp of ", playerB.Health)

	fmt.Println("the second player damage first player by ", playerB.Power)
	fmt.Println("it is a critical damage ")
	fmt.Println("ONE HIT KO")

	die(playerA)

	fmt.Println("the health of player one is ", playerA.Health)
	fmt.Println("the winner of this game is ", playerB.Name)

}

func die(player *Player) {
	time.Sleep(time.Second * 2)

	player.Health = 0
	fmt.Println("the game is over  ")
}
