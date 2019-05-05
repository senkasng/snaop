package util

import (
	//"github.com/robfig/cron"
	//"../handler"
)

type Crontab struct {
	Rule string
	Function func()
}

var Crontabs []Crontab

//func init() {
//	cron1 := Crontab {
//		"*/10 * * * * ?",
//		handler.UpdateTopic,
//	}

//	Crontabs = append(Crontabs,cron1)
//}