package main

import (
	"log"
	"mqt-tech-broker/broker"
	"mqt-tech-broker/db"
)

func main() {
	// init DB
	db.InitRedis()
	db.InitPostgres()

	// start broker
	log.Println("🚀 MQTT Broker starting...")
	broker.Start()
}
