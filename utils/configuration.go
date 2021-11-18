package utils

import (
	db "SDCC/database"
	"encoding/json"
	"log"
	"os"
	"time"
)

const (
	AwsRegion = "us-east-1"
	Replicas  = 4
	Timeout   = 5 * time.Second
	CostRead  = 1
	CostWrite = 2
)

var N = Replicas + 1
var Threshold uint64
var DynamoTable string
var MigrationWindowMinutes int
var TestingServer string
var MigrationThreshold int
var TestingMode bool
var MigrationPeriodSeconds int

type Configuration struct {
	Database  db.Database
	awsRegion string
}

func GetConfiguration() Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	decoder := json.NewDecoder(file)
	parser := struct {
		Database               string
		ReplicationFactor      int
		OffloadingThreshold    uint64
		DynamoTable            string
		MigrationWindowMinutes int
		TestingServer          string
		MigrationThreshold     int
		TestingMode            bool
		MigrationPeriodSeconds int
	}{}
	err = decoder.Decode(&parser)
	if err != nil {
		log.Fatal(err)
	}

	var database db.Database
	if parser.Database == "badger" {
		database = db.BadgerDB{Db: db.GetBadgerDb()}
	} else if parser.Database == "redis" {
		database = db.RedisDB{Db: db.ConnectToRedis()}
	} else {
		database = nil // TODO handle default
	}

	N = parser.ReplicationFactor
	Threshold = parser.OffloadingThreshold
	DynamoTable = parser.DynamoTable
	MigrationWindowMinutes = parser.MigrationWindowMinutes
	MigrationThreshold = parser.MigrationThreshold
	TestingServer = parser.TestingServer
	TestingMode = parser.TestingMode
	MigrationPeriodSeconds = parser.MigrationPeriodSeconds

	return Configuration{Database: database}
}
