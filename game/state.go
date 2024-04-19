package game

import (
	"fmt"
	"math/rand"
)

// Start a new game with two player names and starting scores
func StartGame(pOneId, pTwoId string, startingScores int) *GameState {
	// Roll first hand
	var hand *Hand = &Hand{Values: []int32{1, 1, 1, 1, 1}}
	for i := range hand.Values {
		hand.Values[i] = diceRoll()
	}

	// Make players
	p1 := &Player{
		Id:    pOneId,
		Score: int32(startingScores),
		State: &Player_Bidder{},
	}
	p2 := &Player{
		Id:    pTwoId,
		Score: int32(startingScores),
		State: &Player_Challenger{},
	}

	// Initialize state
	return &GameState{
		PlayerOne:          p1,
		PlayerTwo:          p2,
		BidderId:           p1.Id,
		ChallengerId:       p2.Id,
		CurrentHand:        hand,
		CurrentBid:         nil,
		LastBid:            nil,
		CurrentActionState: ActionState_BIDDING,
		Winner:             nil,
	}
}

// Modify a game state by taking a player action.
func TakeAction(g *GameState, a *Action) error {
	player, err := g.getPlayerById(a.PlayerId)
	if err != nil {
		return err
	}

	valid, err := isValidAction(g, player, &a.Action)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("Action %v is not valid", a)
	}

	switch a.GetAction() {
	case PlayerAction_REROLL:
		doReroll(g, a.GetReroll())
		break
	case PlayerAction_BID:
		doBid(g, a.GetBid())
		break
	case PlayerAction_CALL:
		doChallenge(g, true)
		break
	case PlayerAction_STAY:
		doChallenge(g, false)
		break
	default:
		return fmt.Errorf("Action %d is not valid", a.GetAction())
	}

	return nil
}

//------------------------------------------------------------------------------
// Action Logic
//------------------------------------------------------------------------------

func getValidActions() map[ActionState]map[string][]PlayerAction {
	return map[ActionState]map[string][]PlayerAction{
		ActionState_BIDDING: {
			"bidder":     {PlayerAction_BID, PlayerAction_REROLL},
			"challenger": {},
		},
		ActionState_BIDDING_AFTER_REROLL: {
			"bidder":     {PlayerAction_BID},
			"challenger": {},
		},
		ActionState_CHALLENGING: {
			"bidder":     {},
			"challenger": {PlayerAction_CALL, PlayerAction_STAY},
		},
		ActionState_GAME_OVER: {
			"bidder":     {},
			"challenger": {},
		},
	}
}

func isValidAction(g *GameState, p *Player, a *PlayerAction) (bool, error) {
	role, err := getPlayerRoleAsString(p)
	if err != nil {
		return false, err
	}

	validActions := getValidActions()[g.CurrentActionState][role]
	for _, v := range validActions {
		if v == *a {
			return true, nil
		}
	}

	return false, nil
}

func doReroll(g *GameState, d *RerollData) {
	newHand(g, d)

	g.CurrentActionState = ActionState_BIDDING_AFTER_REROLL
}

func doBid(g *GameState, b *BidData) {
	g.LastBid = g.CurrentBid
	g.CurrentBid = b

	g.CurrentActionState = ActionState_CHALLENGING
}

func doChallenge(g *GameState, accuse bool) {
	// If there is a challenge, see who wins and distributes points
	if accuse {
		var actual ValidBid = g.CurrentHand.value()
		var bid = g.CurrentBid.GetBid()
		fmt.Println(fmt.Sprintf("actual %d, bid %d", actual, bid))
		if actual < bid { // Bidder Loses
			g.getBidder().Score -= g.CurrentBid.GetWager()
			g.getChallenger().Score += g.CurrentBid.GetWager()
		} else { // Challenger Loses
			g.getBidder().Score += g.CurrentBid.GetWager()
			g.getChallenger().Score -= g.CurrentBid.GetWager()
		}
		g.CurrentBid = nil
	}

	g.CurrentActionState = ActionState_BIDDING
    g.swapPlayerRoles()
	fullReroll := RerollData{Offsets: []int32{0, 1, 2, 3, 4}}
	newHand(g, &fullReroll)
	g.LastBid = g.CurrentBid
	g.CurrentBid = nil
}

//------------------------------------------------------------------------------
// Player Utilities
//------------------------------------------------------------------------------

func (g *GameState) swapPlayerRoles() {
    bidder, challenger := g.getBidder(), g.getChallenger()
	var temp string = g.BidderId
	g.BidderId = g.ChallengerId
	g.ChallengerId = temp
    bidder.State = &Player_Challenger{}
    challenger.State = &Player_Bidder{}
}

func getPlayerRoleAsString(p *Player) (string, error) {
	switch x := p.GetState().(type) {
	case *Player_Bidder:
		return "bidder", nil
	case *Player_Challenger:
		return "challenger", nil
	case nil:
		return "", fmt.Errorf("Player.State is not currently set for player id: %s", p.Id)
	default:
		return "", fmt.Errorf("Player.State has unexpected type %T", x)
	}
}

//------------------------------------------------------------------------------
// Hand Utilities
//------------------------------------------------------------------------------

func getHandValue(h *Hand) ValidBid {
	return h.value()
}

func newHand(g *GameState, d *RerollData) {
	var hand *Hand = g.CurrentHand
	for _, i := range d.Offsets {
		hand.Values[i] = diceRoll()
	}
}

func diceRoll() int32 {
	return int32(rand.Intn(6) + 1)
}

//------------------------------------------------------------------------------
// Game Utilities
//------------------------------------------------------------------------------

func (g *GameState) getPlayerById(id string) (*Player, error) {
	if id == g.PlayerOne.Id {
		return g.PlayerOne, nil
	} else if id == g.PlayerTwo.Id {
		return g.PlayerTwo, nil
	}

	return nil, fmt.Errorf("No player in game with ID %s", id)
}

// TODO: Make this handle the id not being found
func (g *GameState) getBidder() *Player {
	if g.PlayerOne.Id == g.BidderId {
		return g.PlayerOne
	}
	return g.PlayerTwo
}

func (g *GameState) getChallenger() *Player {
	if g.PlayerOne.Id == g.ChallengerId {
		return g.PlayerOne
	}
	return g.PlayerTwo
}
