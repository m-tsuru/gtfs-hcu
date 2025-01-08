package main

import (
	"log"
	"os"

	"github.com/m-tsuru/gtfs-hcu/lib"
)

var static = "https://ajt-mobusta-gtfs.mcapps.jp/static/8/current_data.zip"
var realtime_tripUpdate = "https://ajt-mobusta-gtfs.mcapps.jp/realtime/8/trip_updates.bin"
var realtime_vehiclePosition = "https://ajt-mobusta-gtfs.mcapps.jp/realtime/8/vehicle_position.bin"
var realtime_alert = "https://ajt-mobusta-gtfs.mcapps.jp/realtime/8/alerts.bin"

var out_dir = "./data"
var fn = "static.zip"

func getStaticData() error {
	err := lib.Download(static, out_dir, fn, true)
	if err != nil {
		log.Fatalln("Error:", err)
		return err
	}
	return nil
}

func main() {
	err := getStaticData()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	os.Exit(0)
}
