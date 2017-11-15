package go_bgg_lib

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Boardgame struct {
	XMLName       xml.Name `xml:boardgame`
	Yearpublished string   `xml:"yearpublished"`
	Minplayers    int      `xml:"minplayers"`
	Maxplayers    int      `xml:"maxplayers"`
	Playingtime   int      `xml:"playingtime"`
	Minplaytime   int      `xml:"minplaytime"`
	Maxplaytime   int      `xml:"maxplaytime"`
	Age           int      `xml:"age"`
	Name          string   `xml:"name"`
	Description   string   `xml:"description"`
	Thumbnail     string   `xml:"thumbnail"`
	Image         string   `xml:"image"`
	Designer      string   `xml:"boardgamedesigner"`
}

type Boardgames struct {
	XMLName    xml.Name   `xml:"boardgames""`
	Boardgames *Boardgame `xml:"boardgame"`
}

func getGameXML(url string) Boardgames {
	response, err := http.Get(url)
	var v Boardgames
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

func GetBoardgame(id string) Boardgames {
	s := "https://boardgamegeek.com/xmlapi/boardgame/" + id
	game := getGameXML(s)
	return game
}
