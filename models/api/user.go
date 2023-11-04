package api

import "github.com/lonzzi/leetcode-daily-rank/models"

type User struct {
	UserSlug           string
	RealName           string
	AboutMe            string
	UserAvatar         string
	Gender             string
	Websites           models.Slices[string]
	SkillTags          models.Slices[string]
	IPRegion           string
	Location           string
	UseDefaultAvatar   bool
	UserCalendar       models.UserCalendar
	RecentACSubmission struct {
		RecentACSubmission models.RecentACSubmission
		Question           models.Question
	}
}
