package rules

import (
	"fmt"
	"sort"

	"github.com/prfc0/aksha/internal/action"
	"github.com/prfc0/aksha/internal/player"
)

// ValidateAction validates a player's action based on the game rules.
func ValidateAction(playerAction *action.Action, player *player.Player, currentBet int) error {
	switch playerAction.Type {
	case action.Bet, action.Raise:
		if playerAction.Amount > player.Stack {
			return fmt.Errorf("player does not have enough chips to %s %d", playerAction.Type, playerAction.Amount)
		}
		if playerAction.Amount < currentBet {
			return fmt.Errorf("minimum %s amount is %d", playerAction.Type, currentBet)
		}
	case action.Call:
		if playerAction.Amount > player.Stack {
			return fmt.Errorf("player does not have enough chips to call %d", playerAction.Amount)
		}
	case action.Fold:
		// No validation needed for Fold
	default:
		return fmt.Errorf("invalid action type: %s", playerAction.Type)
	}
	return nil
}

func DetermineNextPlayer(players []*player.Player, currentPosition int) int {
	for i := 1; i <= len(players); i++ {
		nextPosition := (currentPosition + i) % len(players)
		if players[nextPosition].Active {
			return nextPosition
		}
	}
	return -1
}

func CalculateSidePots(players []*player.Player) map[*player.Player]int {
	sidePots := make(map[*player.Player]int)

	// Sort players by their contribution to the pot
	sortedPlayers := make([]*player.Player, len(players))
	copy(sortedPlayers, players)
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return sortedPlayers[i].Stack < sortedPlayers[j].Stack
	})

	for i, player := range sortedPlayers {
		if player.Stack > 0 {
			pot := player.Stack * (len(sortedPlayers) - i)
			sidePots[player] = pot
		}
	}

	return sidePots
}
