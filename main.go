package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

func gatherInfo(userName string) {
	log.Println("Start goroutine for " + userName)
	var respJson FAHUserInfo
	resp, err := http.Get("https://api2.foldingathome.org/user/" + userName)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respJson)
	if err != nil {
		log.Fatal(err)
	}
	for _, team := range respJson.Teams {
		teamPoints.WithLabelValues(respJson.Name, team.Name, strconv.Itoa(team.Team)).Set(float64(team.Score))
	}
	userPoints.WithLabelValues(respJson.Name).Set(float64(respJson.Score))
	userWus.WithLabelValues(respJson.Name).Set(float64(respJson.Wus))
	userRank.WithLabelValues(respJson.Name).Set(float64(respJson.Rank))
}

func main() {
	prometheus.MustRegister(userPoints)
	prometheus.MustRegister(teamPoints)
	prometheus.MustRegister(userWus)
	prometheus.MustRegister(userRank)

	config := readUsers("config.yaml")

	go func(users []string) {
		for true {
			for _, userName := range config.Users {
				go gatherInfo(userName)
			}
			log.Println("Sleep for " + strconv.FormatInt(config.Sleep, 10) + " minutes before next gather")
			time.Sleep(time.Duration(config.Sleep) * time.Minute)
		}
	}(config.Users)

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":"+strconv.FormatInt(config.ListenPort, 10), nil)
	if err != nil {
		panic(err)
	}
}
