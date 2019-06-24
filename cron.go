package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

var Cron = new(CronService)

type CronService struct {
	cron *cron.Cron
}

func (s *CronService) Boot() (err error) {
	s.cron = cron.New()
	err = s.cron.AddFunc("0 * * * * *", func() {
		fmt.Println("First scheduled task.")
	})
	if err != nil {
		return err
	}
	err = s.cron.AddFunc("0 0 0 * * *", func() {
		fmt.Println("Second scheduled task.")
	})
	if err != nil {
		return err
	}
	s.cron.Start()
	return nil
}

func main() {
	_ = Cron.Boot()
	for {
		fmt.Println("Waiting......")
		<-time.After(time.Second * 10)
	}
}
