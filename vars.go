package main

import "github.com/prometheus/client_golang/prometheus"

var (
	userPoints = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "foldingathome",
			Subsystem: "user",
			Name:      "points",
			Help:      "Total user points from FoldingAtHome api",
		},
		[]string{
			"user",
		},
	)
	userWus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "foldingathome",
			Subsystem: "user",
			Name:      "wus",
			Help:      "User completed working units",
		},
		[]string{
			"user",
		},
	)
	userRank = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "foldingathome",
			Subsystem: "user",
			Name:      "rank",
			Help:      "User rank",
		},
		[]string{
			"user",
		},
	)
	teamPoints = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "foldingathome",
			Subsystem: "team",
			Name:      "points",
			Help:      "Team user points from FoldingAtHome api",
		},
		[]string{
			"user",
			"team",
			"teamid",
		},
	)
)
