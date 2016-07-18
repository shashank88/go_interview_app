package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func selectOne(name_list []string) string {
	return name_list[rand.Intn(len(name_list))]
}

func winLoss(percentage int) bool {
	if (rand.Int() % 100) <= percentage {
		return true
	}
	return false
}

func loadFileData() ([]string, []string) {
	fnames, err := os.Open("dummy_names.txt")
	if err != nil {
		panic(err)
	}
	defer fnames.Close()
	scanner := bufio.NewScanner(fnames)
	for scanner.Scan() {
		dummyNames = append(dummyNames, strings.TrimSpace(scanner.Text()))
	}

	fsites, err := os.Open("dummy_websites.txt")
	if err != nil {
		panic(err)
	}
	defer fsites.Close()
	scanner = bufio.NewScanner(fsites)
	for scanner.Scan() {
		dummyWebsites = append(dummyWebsites, scanner.Text()+".com")
	}

	return dummyNames, dummyWebsites
}

func validUserId(userId int) bool {
	return userId >= 0 && userId < NumUsers
}

func validBetId(betId int) bool {
	return betId >= 0 && betId < NumBets
}

func validAffiliateId(affiliateId int) bool {
	return affiliateId >= 0 && affiliateId < NumAffiliates
}

func betsByUser(userId int) Bets {
	var userBets Bets
	for i := 0; i < NumBets; i++ {
		if bets[i].UserId == userId {
			userBets = append(userBets, bets[i])
		}
	}
	return userBets
}
