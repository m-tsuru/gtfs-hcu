package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/m-tsuru/gtfs-hcu/lib"
)

var static = "https://ajt-mobusta-gtfs.mcapps.jp/static/8/current_data.zip"
var realtime_tripUpdate = "https://ajt-mobusta-gtfs.mcapps.jp/realtime/8/trip_updates.bin"
var realtime_vehiclePosition = "https://ajt-mobusta-gtfs.mcapps.jp/realtime/8/vehicle_position.bin"
var realtime_alert = "https://ajt-mobusta-gtfs.mcapps.jp/realtime/8/alerts.bin"

var data_dir = "./data"

var static_db = "static.sqlite3"
var static_fn = "static.zip"

func getStaticData() error {
	err := lib.Download(static, data_dir, static_fn, true)
	if err != nil {
		log.Fatalln("Error:", err)
		return err
	}
	return nil
}

func main() {
	db := filepath.Join(data_dir, static_db)

	err := lib.InitDatabase(db)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = getStaticData()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	os.Exit(0)
}
