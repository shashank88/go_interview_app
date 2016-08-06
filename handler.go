package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("interview_question.html")
	t.Execute(w, nil)
}

func Bonus(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("bonus_question.html")
	t.Execute(w, nil)
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func AffiliateIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(affiliates); err != nil {
		panic(err)
	}
}

func BetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bets); err != nil {
		panic(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var userId int
	var err error
	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}
	if validUserId(userId) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(users[userId]); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "User not found."}); err != nil {
		panic(err)
	}
}

func UserBetShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var userId int
	var err error
	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}
	if validUserId(userId) {
		userBets := betsByUser(userId)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(userBets); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "User not found."}); err != nil {
		panic(err)
	}
}

func BetShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var betId int
	var err error
	if betId, err = strconv.Atoi(vars["betId"]); err != nil {
		panic(err)
	}
	if validBetId(betId) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(bets[betId]); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Bet not found."}); err != nil {
		panic(err)
	}
}

func AffiliateShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var affiliateId int
	var err error
	if affiliateId, err = strconv.Atoi(vars["affiliateId"]); err != nil {
		panic(err)
	}
	if validAffiliateId(affiliateId) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(affiliates[affiliateId]); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Affiliate not found."}); err != nil {
		panic(err)
	}
}
