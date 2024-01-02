package main

type FAHUserInfo struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	Score    int    `json:"score"`
	Wus      int    `json:"wus"`
	Rank     int64  `json:"rank"`
	Active50 int    `json:"active_50"`
	Active7  int    `json:"active_7"`
	Last     string `json:"last"`
	Users    int    `json:"users"`
	Teams    []struct {
		Team     int    `json:"team"`
		Name     string `json:"name"`
		Trank    int    `json:"trank"`
		Tscore   int64  `json:"tscore"`
		Twus     int    `json:"twus"`
		Founder  string `json:"founder"`
		URL      string `json:"url"`
		Logo     string `json:"logo"`
		Score    int    `json:"score"`
		Wus      int    `json:"wus"`
		Last     int    `json:"last"`
		Active50 int    `json:"active_50"`
		Active7  int    `json:"active_7"`
	} `json:"teams"`
}

type YamlConfig struct {
	Users      []string `yaml:"users"`
	Sleep      int64    `yaml:"sleep"`
	ListenPort int64    `yaml:"port"`
}
