package go_bgg_lib

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Stats struct {
	MinPlayers  int   `xml:"minplayers,attr"`
	MaxPlayers  int   `xml:"maxplayers,attr"`
	MinPlaytime int   `xml:"minplaytime"`
	MaxPlaytime int   `xml:"maxplaytime"`
	Playingtime int   `xml:"playingtime"`
	TotalOwners int64 `xml:"numowned"`
	Rating      struct {
		Average struct {
			Average float32 `xml:"value,attr"`
		} `xml:"average"`
		BayesAverage struct {
			BayesAverage float32 `xml:"value,attr"`
		} `xml:"bayesaverage"`
	} `xml:"rating"`
}

type Items struct {
	XMLName   xml.Name `xml:items`
	Name      string   `xml:"name"`
	Published string   `xml:"yearpublished"`
	Image     string   `xml:"image"`
	Thumbnail string   `xml:"thumbnail"`
	Comment   string   `xml:"comment"`

	Stats struct {
		MinPlayers  int   `xml:"minplayers,attr"`
		MaxPlayers  int   `xml:"maxplayers,attr"`
		MinPlaytime int   `xml:"minplaytime"`
		MaxPlaytime int   `xml:"maxplaytime"`
		Playingtime int   `xml:"playingtime"`
		TotalOwners int64 `xml:"numowned"`
		Rating      struct {
			Average struct {
				Average float32 `xml:"value,attr"`
			} `xml:"average"`
			BayesAverage struct {
				BayesAverage float32 `xml:"value,attr"`
			} `xml:"bayesaverage"`
		} `xml:"rating"`
	} `xml:"stats"`

	Status struct {
		Owned           int `xml:"own,attr"`
		PreviouslyOwned int `xml:"prevowned,attr"`
		ForTrade        int `xml:"fortrade,attr"`
		Wanted          int `xml:"want,attr"`
		ToPlay          int `xml:"wanttoplay,attr"`
		ToBuy           int `xml:"wanttobuy,attr"`
		Wishlist        int `xml:"wishlist,attr"`
		Preordered      int `xml:"preordered,attr"`
	} `xml:"status"`
	Played int `xml:"numplays"`
}

type Collection struct {
	XMLName xml.Name `xml:"items""`
	Items   []Items  `xml:"item"`
}

func getCollectionXML(url string) Collection {
	response, err := http.Get(url)
	if response.StatusCode == 202 {
		log.Fatal("The collection you requested was not available but is now queued. Try again later.")
	}
	var v Collection
	if err != nil {
		defer response.Body.Close()
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

func GetCollection(user string) Collection {
	s := "https://boardgamegeek.com/xmlapi/collection/" + user
	collection := getCollectionXML(s)
	return collection
}
