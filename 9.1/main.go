package main

import (
	"container/ring"
	"fmt"
)

func main() {
	players := 465
	lastMarble := 71498

	scores := make([]int, players, players)
	marbleValue := 0

	game := ring.New(1)
	game.Value = marbleValue

	for {
		marbleValue++

		if marbleValue == lastMarble {
			break
		}

		if marbleValue%23 == 0 {
			player := marbleValue % players
			scores[player] += marbleValue
			game = game.Move(-7)
			scores[player] += game.Value.(int)

			game = game.Prev()
			game.Unlink(1)
			game = game.Next()

			continue
		}

		game = game.Next()

		r := ring.New(1)
		r.Value = marbleValue

		game = game.Link(r)
		game = game.Prev()
	}

	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}

	fmt.Printf("%d\n", max)
}
