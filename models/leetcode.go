package models

type User struct {
	UserSlug         string `gorm:"primaryKey"`
	RealName         string
	AboutMe          string
	UserAvatar       string
	Gender           string
	Websites         Slices[string]
	SkillTags        Slices[string]
	IPRegion         string
	Location         string
	UseDefaultAvatar bool
}

type UserCalendar struct {
	UserSlug           string `gorm:"primaryKey"`
	Streak             int    // 连续提交
	TotalActiveDays    int    // 总活跃天数
	SubmissionCalendar string // 每日提交
	ActiveYears        Slices[int]
	RecentStreak       int
}

type RecentACSubmission struct {
	UserSlug           string `gorm:"primaryKey"`
	SubmissionID       int
	SubmitTime         int64
	QuestionFrontendID string
}

type Question struct {
	Title              string
	TranslatedTitle    string
	TitleSlug          string
	QuestionFrontendID string `gorm:"primaryKey"`
}
