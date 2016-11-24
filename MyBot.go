package main

import (
	"hlt"
    "math/rand"
)


func move (playerTag int, gameMap hlt.GameMap, location hlt.Location) hlt.Move {
    site := gameMap.GetSite(location, hlt.STILL)
    // Always STILL if moving won't achieve anything
    if site.Strength == 0 {
        return hlt.Move {location, hlt.STILL}
    }

    // Avoid moving into a site with higher strength
    var options []hlt.Direction
    for _, direction := range(hlt.CARDINALS) {
        option := gameMap.GetSite(location, direction)
        if option.Owner == playerTag || site.Strength > option.Strength {
            options = append(options, direction)
        }
    }
    return hlt.Move {location, options[rand.Intn(len(options))]}
}


func main () {
	conn, gameMap := hlt.NewConnection("MyBot")
	for {
		var moves hlt.MoveSet
		gameMap = conn.GetFrame()
		for y := 0; y < gameMap.Height; y++ {
			for x := 0; x < gameMap.Width; x++ {
				loc := hlt.NewLocation(x,y)
				if gameMap.GetSite(loc, hlt.STILL).Owner == conn.PlayerTag {
					moves = append(moves, move(conn.PlayerTag, gameMap, loc))
				}
			}
		}
		conn.SendFrame(moves)
	}
}
