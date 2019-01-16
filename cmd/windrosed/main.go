package main

import (
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("/Users/igreg/Downloads/GeoLite2-City_20190108/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("81.2.69.142")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(record)
}
