package table

import (
	"fmt"
	"log"

	"github.com/prfc0/aksha/internal/player"
)

type Table struct {
	Players        []*player.Player // List of players at the table
	DealerPosition int              // Position of the dealer button
	MaxPlayers     int              // Maximum number of players allowed at the table
}

func NewTable(maxPlayers int) *Table {
	return &Table{
		Players:        make([]*player.Player, 0),
		DealerPosition: 0,
		MaxPlayers:     maxPlayers,
	}
}

func (t *Table) AddPlayer(player *player.Player) error {
	if len(t.Players) >= t.MaxPlayers {
		return fmt.Errorf("table is full: max %d players allowed", t.MaxPlayers)
	}
	t.Players = append(t.Players, player)
	log.Printf("Player %s joined the table.\n", player.Name)
	return nil
}

func (t *Table) RemovePlayer(playerID string) {
	for i, player := range t.Players {
		if player.ID == playerID {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			log.Printf("Player %s left the table.\n", player.Name)
			return
		}
	}
}

func (t *Table) RotateDealer() {
	t.DealerPosition = (t.DealerPosition + 1) % len(t.Players)
	log.Printf("Dealer button moved to player %s.\n", t.Players[t.DealerPosition].Name)
}

// ActivePlayers returns a list of active players at the table.
func (t *Table) ActivePlayers() []*player.Player {
	activePlayers := make([]*player.Player, 0)
	for _, player := range t.Players {
		if player.Active {
			activePlayers = append(activePlayers, player)
		}
	}
	return activePlayers
}
