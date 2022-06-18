package main

import (
	bot "github.com/datagrove/ironbot_go"
)

func main() {
	id := map[string]bool{
		"jimh@datagrove.com": true,
		"4843664923":         true,
	}

	bot.CertifyBot(func(
		v *bot.CertificationRequest,
		out *bot.Certification) error {

		if id[v.Email] || id[v.Phone] {
			out.CertifiedSeconds = 24 * 3600
		}
		return nil
	})
}
