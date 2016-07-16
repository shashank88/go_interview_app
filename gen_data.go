package main

import (
	"math/rand"
	"time"
)

const (
	MaxBet        = 10
	MinOdds       = 10
	NumUsers      = 100
	NumAffiliates = 10
	NumBets       = 500
)

// Poor Man's Database.
var users Users
var bets Bets
var affiliates Affiliates
var dummyWebsites []string
var dummyNames []string

func genAffiliates(num int) Affiliates {
	dummy_affiliates := make(Affiliates, num)

	for i := 0; i < num; i++ {
		dummy_affiliates[i] = Affiliate{i, selectOne(dummyNames), selectOne(dummyWebsites)}
	}

	return dummy_affiliates
}

func genUsers(num int, aff_num int) Users {
	dummy_users := make(Users, num)

	for i := 0; i < num; i++ {
		dummy_users[i] = User{i, selectOne(dummyNames), rand.Intn(aff_num), time.Now()}
	}

	return dummy_users
}

func genBets(num int, user_num int) Bets {
	dummy_bets := make(Bets, num)
	var amount_bet int
	var stake float32
	var percentage_odds int

	for i := 0; i < num; i++ {
		amount_bet = rand.Intn(MaxBet)
		percentage_odds = (MinOdds + rand.Intn(100-MinOdds)) % 100 // precautionary %
		stake = (float32(100) / float32(percentage_odds)) * float32(amount_bet)
		dummy_bets[i] = Bet{i, rand.Intn(user_num), amount_bet, stake, time.Now(), winLoss(percentage_odds)}
	}

	return dummy_bets
}

func genDummyData(affiliate_count int, user_count int, bet_count int) {
	rand.Seed(time.Now().Unix())
	dummyNames, dummyWebsites = loadFileData()
	affiliates = genAffiliates(affiliate_count)
	users = genUsers(user_count, affiliate_count)
	bets = genBets(bet_count, user_count)
}

func initApp() {
	genDummyData(NumAffiliates, NumUsers, NumBets)
}
