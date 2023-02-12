package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/gocarina/gocsv"
)

type rate struct {
	Account string  `csv:"account"`
	User    string  `csv:"user"`
	Rate    float32 `csv:"rate"`
}

// match scores how well an account and user matches
// the rate. The user is heavily weighted:
//   - user match: +100
//   - rate user not set: 0
//   - user's don't match: return immediately with score of 0
//
// the account score is based on the number of characters
// in the rate account that match the query account param.
// For instance, if the rate account is cust:a, and the query
// account is cust:a:proj1, then we match 6 charcters and add
// 6 to the score. If we find a rate later of cust:a:proj1, then
// that would add 12 to the score.
func (r *rate) match(account string, user string) int {
	var score int

	if r.User != "" {
		if r.User == user {
			score += 100
		} else {
			return 0
		}
	}

	if strings.HasPrefix(account, r.Account) {
		score += len(r.Account)
	}

	return score
}

type rates []rate

func (r *rates) find(account string, user string) (float32, bool) {
	var ret rate
	var score int

	for _, rt := range *r {
		s := rt.match(account, user)
		if s > score {
			ret = rt
			score = s
		}
	}

	if score <= 0 {
		return 0, false
	}

	return ret.Rate, true
}

func parseRates(r io.Reader) (rates, error) {
	var rates rates

	if err := gocsv.Unmarshal(r, &rates); err != nil {
		return rates, fmt.Errorf("csv error: %v", err)
	}

	return rates, nil
}
