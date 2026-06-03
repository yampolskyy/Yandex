package main

import "sort"

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

var players []Player

func NewPlayer(name string, goals, misses, assists int) Player {
	player := Player{name, goals, misses, assists, calculateRating(goals, misses, assists)}
	players = append(players, player)
	return player
}

func calculateRating(goals, misses, assists int) float64 {
	if misses != 0 {
		return (float64(goals) + float64(assists)/2) / float64(misses)
	}
	return float64(goals) + float64(assists)/2
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			return players[i].Name < players[j].Name
		}
		return players[i].Goals > players[j].Goals
	})
	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Misses == 0 && players[j].Misses == 0 {
			if players[i].Goals == players[j].Goals {
				return players[i].Name < players[j].Name
			}
			return players[i].Goals > players[j].Goals
		}

		if players[i].Misses == 0 {
			return true
		}

		if players[j].Misses == 0 {
			return false
		}

		if players[i].Goals/players[i].Misses == players[j].Goals/players[j].Misses {
			return players[i].Name < players[j].Name
		}
		return players[i].Goals/players[i].Misses > players[j].Goals/players[j].Misses
	})
	return players
}
