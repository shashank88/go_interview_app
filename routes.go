package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserIndex",
		"GET",
		"/users",
		UserIndex,
	},
	Route{
		"AffiliateIndex",
		"GET",
		"/affiliates",
		AffiliateIndex,
	},
	Route{
		"BetIndex",
		"GET",
		"/bets",
		BetIndex,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userId}",
		UserShow,
	},
	Route{
		"BetShow",
		"GET",
		"/bets/{betId}",
		BetShow,
	},
	Route{
		"AffiliateShow",
		"GET",
		"/affiliates/{affiliateId}",
		AffiliateShow,
	},
}
