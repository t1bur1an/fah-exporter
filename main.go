package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
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
	userWus.WithLabelValues(respJson.Name).Set(float64(respJson.Wus))
	userRank.WithLabelValues(respJson.Name).Set(float64(respJson.Rank))
	active7.WithLabelValues(respJson.Name).Set(float64(respJson.Active7))
	userPoints.WithLabelValues(respJson.Name).Set(float64(respJson.Score))
	active50.WithLabelValues(respJson.Name).Set(float64(respJson.Active50))
}

func main() {
	var promVars = []prometheus.Collector{
		userPoints,
		teamPoints,
		userWus,
		userRank,
		active50,
		active7,
	}
	for _, promVar := range promVars {
		prometheus.MustRegister(promVar)
	}

	config := readUsers("config.yaml")
	if config.ListenPort > 20000 {
		log.Fatal("Listen port is bigger than 20000")
		os.Exit(2)
	}

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
