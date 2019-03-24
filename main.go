package main

import (
	"encoding/json"
	"flag"
	"github.com/agirot/crawler-maxtrain/data"
	"github.com/agirot/crawler-maxtrain/manager"
	"log"
	"os"
	"time"
)

func init() {
	data.ConfigPath = *flag.String("config", "config.json", "config file path (default in bin location)")
	data.SmsKeyArg = flag.String("sms_key", "", "List days to send alert")
	data.SmsUserArg = flag.String("sms_user", "", "List days to send alert")
	flag.Parse()
	err := hydrateConfig()

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	for {
		for _, watchDay := range data.Config {
			err := manager.Process(watchDay)
			if err != nil {
				log.Println(err.Error())
			}
		}
		time.Sleep(30 * time.Minute)
	}
}

//checkConf valid configuration file
func hydrateConfig() error {
	file, err := os.Open(data.ConfigPath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data.Config)
	if err != nil {
		panic(err)
	}

	for _, input := range data.Config {
		err := data.CheckDayExist(input.Day)
		if err != nil {
			return err
		}
	}
	return nil
}
