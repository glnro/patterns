package main

import (
	"fmt"
	"sort"
)

type AuctionType string

// Strategy Interface
type AuctionStrategy interface {
	Execute(a *Auction)
}

// Concrete Type
type ReserveAuction struct {
}

// Concrete Type
type SecondBidAuction struct {
}

func (auction *ReserveAuction) Execute(a *Auction) {
	winner := a.GetWinner(len(a.Bids) - 1)
	fmt.Printf("%v Auction :: Selecting the Highest Bid: %d\n", a.AType, winner)
}

func (auction *SecondBidAuction) Execute(a *Auction) {
	winner := a.GetWinner(len(a.Bids) - 2)
	fmt.Printf("%v Auction :: Selecting Second Highest Bid: %d\n", a.AType, winner)
}

// Context
type Auction struct {
	Strategy AuctionStrategy
	AType    AuctionType
	Bids     []int
}

func initAuction(aStrat AuctionStrategy, aType AuctionType, bids []int) *Auction {
	return &Auction{
		Strategy: aStrat,
		AType:    aType,
		Bids:     bids,
	}
}

func (a *Auction) Execute() {
	a.Strategy.Execute(a)
}

func (a *Auction) GetWinner(wIdx int) int {
	sortedBids := a.Bids
	sort.Ints(sortedBids)
	return sortedBids[wIdx]
}

func main() {
	reserveStrat := &ReserveAuction{}
	secondStrat := &SecondBidAuction{}

	bds := []int{1, 4, 7, 3, 5, 8, 2}

	reserveAuction := initAuction(reserveStrat, "RESERVE", bds)
	secondAuction := initAuction(secondStrat, "SECOND", bds)

	auctions := []Auction{*reserveAuction, *secondAuction}

	for _, a := range auctions {
		a.Execute()
	}
}
