package main

import (
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	AffiliateId int       `json:"affiliate_id"`
	Created     time.Time `json:"created"`
}

type Bet struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	Amount         float32   `json:"amount"`
	PercentageOdds int       `json:"percentage_odds"`
	Timestamp      time.Time `json:"timestamp"`
	Result         bool      `'json:result'`
}

type Affiliate struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Website string `json:"website"`
}

type Users []User
type Bets []Bet
type Affiliates []Affiliate
