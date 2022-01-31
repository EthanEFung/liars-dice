package game

import (
	"math/rand"
	"time"
)

type Die struct {
	Face int
	Seed int64
}
func (d Die) Roll() {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + d.Seed))
	d.Face = r.Intn(6) + 1
}

type Wager struct {
	Amount int
	Face int
}

type Player struct {
	Name string
	Dice []Die
}
func (p Player) Bid(r *Round, wager Wager) bool {
	last := r.Last.Wager
	if wager.Amount < last.Amount {
		return false
	}
	if wager.Amount == last.Amount && wager.Face <= last.Face {
		return false
	}

	r.Last = Turn{
		Player: p,
		Wager: wager,
	}
	var next Player
	for i, player := range r.Players {
		if player.Name == p.Name {
			next = *r.Players[i+1]
		} else if i == len(r.Players) - 1 {
			next = *r.Players[0]
		}
	}
	r.Current = Turn{
		Player: next,
	}
	return true
}
func (p Player) Call(r *Round) {
	var amt int // count the number of dice with the face
	wager := r.Last.Wager
	for _, player := range r.Players {
		for _, die := range player.Dice {
			if die.Face == wager.Face {
				amt++
			}
		}
	}
	if wager.Amount < amt {
		// bidder wins
		r.Loser = p
	} else {
		// caller wins
		r.Loser = r.Last.Player
	}
	// remove a die from the losing player
	r.Loser.Dice = r.Loser.Dice[:1]
}

type Turn struct {
	Player Player
	Wager Wager
}

type Round struct {
	Last Turn
	Current Turn
	Players []*Player
	Loser Player
}
func (r Round) Roll() {
	for _, player := range r.Players {
		for _, die := range player.Dice {
			die.Roll()
		}
	}
}

type Game struct {
	Host *Host
	Rounds []*Round
	Winner Player
	Players []*Player
}
func(g Game) NewPlayer(username string) *Player {
	dice := []Die{
		{1, time.Now().UnixNano()},
		{1, time.Now().UnixNano()},
		{1, time.Now().UnixNano()},
		{1, time.Now().UnixNano()},
		{1, time.Now().UnixNano()},
	};
	player := &Player{
		Name: username,
		Dice: dice,
	}
	g.Players = append(g.Players, player)
	return player
}

type Host = Player
func (h Host) StartGame(game *Game) {
  firstTurn := Turn{
		Player: h,
	}
	firstRound := &Round{ Current: firstTurn, Players: game.Players }
	firstRound.Roll()
	game.Rounds = append(game.Rounds, firstRound)
}

func NewGame(username string) *Game {
	host := &Host{
		Name: username,
		Dice: []Die{
			{1, time.Now().UnixNano()},
			{1, time.Now().UnixNano()},
			{1, time.Now().UnixNano()},
			{1, time.Now().UnixNano()},
			{1, time.Now().UnixNano()},
		},
	}

	game := &Game{ Host: host }
	game.Players = append(game.Players, host)
	return game
}