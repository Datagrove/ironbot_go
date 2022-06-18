package main

import (
	bot "github.com/datagrove/ironbot_go"
)

// sample program that simulates an adjudicator
// this chooses json as the file format

type PricedClaim struct {
	MemberId string
	Line     []PricedLine
	Issuer   string
	Plan     string
}
type PricedLine struct {
	NPI       string
	CodeQual  string
	Code      string
	Price     string
	InNetwork bool
}

type Aeob struct {
	Aeob []AeobLine
}
type AeobLine struct {
	Category string
	Code     string
	Amount   float64
}

func processFile(in, out string) {

}

// for this sample we just pay the network
func main() {
	bot.AnswerBot[PricedClaim, Aeob](
		func(in *PricedClaim, out *Aeob) {

		})
}
