package main

import (
	"fmt"

	"llewellyn.dev/liars_dice/game"
)

func main() {
	var instance *game.GameState = game.StartGame("kevin", "anna", 100)
	fmt.Println(fmt.Sprintf("%v", instance))

	fmt.Println("---Rerolling dice 1 and 3---")
	rerollAction := &game.Action_Reroll{
		Reroll: &game.RerollData{Offsets: []int32{0, 2}},
	}
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "kevin",
		Action:   game.PlayerAction_REROLL,
		Data:     rerollAction,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%v", instance))

	fmt.Println("---Can't reroll once already done once---")
	rerollActionAgain := &game.Action_Reroll{
		Reroll: &game.RerollData{Offsets: []int32{0, 2}},
	}
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "kevin",
		Action:   game.PlayerAction_REROLL,
		Data:     rerollActionAgain,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%v", instance))

	fmt.Println("---Kevin can make a bid---")
	BidAction := &game.Action_Bid{
		Bid: &game.BidData{
			Bid:   game.ValidBid_FULL_HOUSE,
			Wager: 20,
		},
	}
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "kevin",
		Action:   game.PlayerAction_BID,
		Data:     BidAction,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%v", instance))

	fmt.Println("---Kevin cannot challenge---")
	challenge := &game.Action_Stay{}
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "kevin",
		Action:   game.PlayerAction_STAY,
		Data:     challenge,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%v", instance))

	fmt.Println("---Anna can challenge---")
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "anna",
		Action:   game.PlayerAction_CALL,
		Data:     challenge,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%v", instance))

	fmt.Println("---Anna reroll and bid now---")
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "anna",
		Action:   game.PlayerAction_REROLL,
		Data:     rerollAction,
	}); err != nil {
		fmt.Println(err)
	}
	annaBid := &game.Action_Bid{
		Bid: &game.BidData{
			Bid:   game.ValidBid_FIVE_OF_A_KIND,
			Wager: 40,
		},
	}
	if err := game.TakeAction(instance, &game.Action{
		PlayerId: "anna",
		Action:   game.PlayerAction_BID,
		Data:     annaBid,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%v", instance))
}
