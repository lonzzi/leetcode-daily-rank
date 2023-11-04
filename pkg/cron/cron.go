package cron

import "github.com/robfig/cron/v3"

var c *cron.Cron

func Init() {
	c = cron.New()
}

func Start() {
	c.Start()
}

func AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return c.AddFunc(spec, cmd)
}

func Remove(id cron.EntryID) {
	c.Remove(id)
}

func Stop() {
	c.Stop()
}
