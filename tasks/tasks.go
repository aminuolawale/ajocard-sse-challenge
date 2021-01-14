package main

import (
	"aminuolawale/ajocard/helpers"
	"time"

	"gopkg.in/robfig/cron.v2"
)

func main() {

		c := cron.New()
		c.AddFunc("@every 0h0m50s", helpers.SaveStatus)
        c.Start()
        time.Sleep(100 * time.Second)
        c.Stop() 
}