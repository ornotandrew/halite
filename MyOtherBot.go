package main

import (
	"hlt"
    "math/rand"
)


func move (gameMap hlt.GameMap, location hlt.Location) hlt.Move {
    site := gameMap.GetSite(location, hlt.STILL)
    if site.Strength == 0 {
        return hlt.Move {
            Location: location,
            Direction: hlt.Direction(hlt.STILL),
        }
    }
    return hlt.Move {
        Location: location,
        Direction: hlt.Direction(rand.Int() % 5),
    }
}


func main () {
	conn, gameMap := hlt.NewConnection("MyOtherBot")
	for {
		var moves hlt.MoveSet
		gameMap = conn.GetFrame()
		for y := 0; y < gameMap.Height; y++ {
			for x := 0; x < gameMap.Width; x++ {
				loc := hlt.NewLocation(x,y)
				if gameMap.GetSite(loc, hlt.STILL).Owner == conn.PlayerTag {
					moves = append(moves, move(gameMap, loc))
				}
			}
		}
		conn.SendFrame(moves)
	}
}
