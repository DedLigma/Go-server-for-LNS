package main

var data struct {
	Solv struct {
		Start_timer int    `json:"start_timer"`
		Time_data   string `json:"time_data"`
		Timer       string `json:"timer"`
		Coordinates struct {
			Latitude int `json:"latitude"`
			Longtude int `json:"longtude"`
		} `json:"coordinates"`
		Height        int    `json:"height"`
		Svs           string `json:"svs"`
		Res           int    `json:"res"`
		Modi_ns       string `json:"modi_ns"`
		Modi_ms       string `json:"modi_ms"`
		Pdop          int    `json:"pdop"`
		E_distance    int    `json:"e_distance"`
		Appar_counter int    `json:"appar_counter"`
	}
	Lct struct {
		Number      []int `json:"number"`
		Time_data   []int `json:"time_data"`
		Timer       []int `json:"timer"`
		Coordinates struct {
			Latitude []int `json:"latitude"`
			Longtude []int `json:"longtude"`
		} `json:"coordinates"`
		Height []int   `json:"heights"`
		Resids [][]int `json:"resids"`
	}
}
