package grtrack

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetApi() {
	const (
		artist    string = "https://groupietrackers.herokuapp.com/api/artists"
		locations string = "https://groupietrackers.herokuapp.com/api/relation"
	)
	err := Decode(artist, &Info.Art)
	if err != nil {
		log.Fatal(err)
	}
	err = Decode(locations, &Info.Rel)
	if err != nil {
		log.Fatal(err)
	}
}

func Decode(url string, Info interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, Info)
}
