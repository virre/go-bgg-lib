package go_bgg_lib

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Geeklist struct {
	XMLName     xml.Name `xml:geeklist`
	Postded     string   `xml:"postdate"`
	Edited      string   `xml:"editdate"`
	Thumbs      int64    `xml:"thumbs"`
	NumItems    int64    `xml:"numitems"`
	Poster      string   `xml:"username"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	ListItem    []Item   `xml:"item"`
}

type Item struct {
	Id      string `xml:"objectid,attr"`
	BggType string `xml:"subtype,attr"`
	Title   string `xml:"objectname,attr"`
	Thumbs  int64  `xml:"thumbs, attr"`
	ImageId string `xml:"imageid, attr"`
}

func getGeeklistXML(url string) Geeklist {
	response, err := http.Get(url)
	var v Geeklist
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

func GetGeeklist(id string) Geeklist {
	s := "https://boardgamegeek.com/xmlapi/geeklist/" + id
	geeklist := getGeeklistXML(s)
	return geeklist
}
