package go_bgg_lib

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type SearchBoardgames struct {
	XMLName   xml.Name `xml:"boardgames""`
	Boardgame struct {
		Id            string `xml:"objectid,attr"`
		Name          string `xml:"name"`
		Yearpublished string `xml:"yearpublished"`
	} `xml:"boardgame"`
}

func getSearchXML(url string) SearchBoardgames {
	response, err := http.Get(url)
	var v SearchBoardgames
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		err = xml.Unmarshal([]byte(data), &v)
		if err != nil {
			log.Fatal(err)
		}
	}
	return v
}

// Used to search for games. For strict searches exact should be true.
// E.g if you want only Go SearchBoardgame("Go", true) instead of every
// game that contains the word Go. This might give false positives.
func SearchBoardgame(game string, exact bool) SearchBoardgames {
	s := "https://boardgamegeek.com/xmlapi/search/?search=" + game
	if exact == true {
		s = s + "&exact=1"
	}
	games := getSearchXML(s)
	return games
}
