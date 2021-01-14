package main

import (
	"aminuolawale/ajocard/helpers"
	"fmt"
	"time"

	"gopkg.in/robfig/cron.v2"
)

func main() {

		c := cron.New()
		c.AddFunc("@every 0h0m50s", helpers.SaveStatus)
		fmt.Println(c)
        c.Start()

        // Funcs are invoked in their own goroutine, asynchronously.

        // Funcs may also be added to a running Cron
        // c.AddFunc("@daily", func() { fmt.Println("Every day") })

        // Added time to see output
        time.Sleep(100 * time.Second)

        c.Stop() // Stop the scheduler (does not stop any jobs already running).

}