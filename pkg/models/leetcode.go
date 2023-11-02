package models

type UserProfileCalendarResp struct {
	Data struct {
		UserCalendar struct {
			Streak             int    `json:"streak"`
			TotalActiveDays    int    `json:"totalActiveDays"`
			SubmissionCalendar string `json:"submissionCalendar"`
			ActiveYears        []int  `json:"activeYears"`
			MonthlyMedals      []struct {
				Name       string `json:"name"`
				ObtainDate string `json:"obtainDate"`
				Category   string `json:"category"`
				Config     struct {
					Icon              string `json:"icon"`
					IconGif           string `json:"iconGif"`
					IconGifBackground string `json:"iconGifBackground"`
				} `json:"config"`
				Progress int    `json:"progress"`
				ID       string `json:"id"`
				Year     int    `json:"year"`
				Month    int    `json:"month"`
			} `json:"monthlyMedals"`
			RecentStreak int `json:"recentStreak"`
		} `json:"userCalendar"`
	} `json:"data"`
}

type RecentACSubmissionsResp struct {
	Data struct {
		RecentACSubmission struct {
			SubmissionID int   `json:"submissionId"`
			SubmitTime   int64 `json:"submitTime"`
			Question     struct {
				Title              string `json:"title"`
				TranslatedTitle    string `json:"translatedTitle"`
				TitleSlug          string `json:"titleSlug"`
				QuestionFrontendID string `json:"questionFrontendId"`
			} `json:"question"`
		} `json:"recentACSubmissions"`
	} `json:"data"`
}

type UserProfileResp struct {
	Data struct {
		UserProfile struct {
			HaveFollowed interface{} `json:"haveFollowed"`
			SiteRanking  int         `json:"siteRanking"`
			Profile      struct {
				UserSlug         string      `json:"userSlug"`
				RealName         string      `json:"realName"`
				AboutMe          string      `json:"aboutMe"`
				ASCIICode        string      `json:"asciiCode"`
				UserAvatar       string      `json:"userAvatar"`
				Gender           string      `json:"gender"`
				Websites         []string    `json:"websites"`
				SkillTags        []string    `json:"skillTags"`
				IPRegion         string      `json:"ipRegion"`
				Birthday         interface{} `json:"birthday"`
				Location         string      `json:"location"`
				UseDefaultAvatar bool        `json:"useDefaultAvatar"`
				Github           interface{} `json:"github"`
				School           struct {
					SchoolID string `json:"schoolId"`
					Logo     string `json:"logo"`
					Name     string `json:"name"`
				} `json:"school"`
				Company        string      `json:"company"`
				Job            interface{} `json:"job"`
				GlobalLocation struct {
					Country      string `json:"country"`
					Province     string `json:"province"`
					City         string `json:"city"`
					OverseasCity bool   `json:"overseasCity"`
				} `json:"globalLocation"`
				SocialAccounts []struct {
					Provider   string `json:"provider"`
					ProfileURL string `json:"profileUrl"`
				} `json:"socialAccounts"`
				SkillSet struct {
					LangLevels []struct {
						LangName        string `json:"langName"`
						LangVerboseName string `json:"langVerboseName"`
						Level           int    `json:"level"`
					} `json:"langLevels"`
					Topics []struct {
						Slug           string `json:"slug"`
						Name           string `json:"name"`
						TranslatedName string `json:"translatedName"`
					} `json:"topics"`
					TopicAreaScores []struct {
						Score     int `json:"score"`
						TopicArea struct {
							Name string `json:"name"`
							Slug string `json:"slug"`
						} `json:"topicArea"`
					} `json:"topicAreaScores"`
				} `json:"skillSet"`
			} `json:"profile"`
			EducationRecordList []struct {
				UnverifiedOrganizationName string `json:"unverifiedOrganizationName"`
			} `json:"educationRecordList"`
			OccupationRecordList []interface{} `json:"occupationRecordList"`
		} `json:"userProfilePublicProfile"`
	} `json:"data"`
}
