package leetcode

import (
	"github.com/lonzzi/leetcode-daily-rank/data"
	"github.com/lonzzi/leetcode-daily-rank/models"
	"github.com/lonzzi/leetcode-daily-rank/models/api"
	lt "github.com/lonzzi/leetcode-daily-rank/pkg/leetcode"
)

func SaveUserProfile(userSlug string) error {
	userProfileResp, err := lt.GetUserProfilePublicProfile(userSlug)
	if err != nil {
		return err
	}

	userCalendarResp, err := lt.GetUserProfileCalendar(userSlug)
	if err != nil {
		return err
	}

	recentACSubmissionsResp, err := lt.GetRecentAcSubmissions(userSlug)
	if err != nil {
		return err
	}

	user := models.User{
		UserSlug:         userProfileResp.Data.UserProfile.Profile.UserSlug,
		RealName:         userProfileResp.Data.UserProfile.Profile.RealName,
		AboutMe:          userProfileResp.Data.UserProfile.Profile.AboutMe,
		UserAvatar:       userProfileResp.Data.UserProfile.Profile.UserAvatar,
		Gender:           userProfileResp.Data.UserProfile.Profile.Gender,
		Websites:         userProfileResp.Data.UserProfile.Profile.Websites,
		SkillTags:        userProfileResp.Data.UserProfile.Profile.SkillTags,
		IPRegion:         userProfileResp.Data.UserProfile.Profile.IPRegion,
		Location:         userProfileResp.Data.UserProfile.Profile.Location,
		UseDefaultAvatar: userProfileResp.Data.UserProfile.Profile.UseDefaultAvatar,
	}

	userCalendar := models.UserCalendar{
		UserSlug:           user.UserSlug,
		Streak:             userCalendarResp.Data.UserCalendar.Streak,
		TotalActiveDays:    userCalendarResp.Data.UserCalendar.TotalActiveDays,
		SubmissionCalendar: userCalendarResp.Data.UserCalendar.SubmissionCalendar,
		ActiveYears:        userCalendarResp.Data.UserCalendar.ActiveYears,
		RecentStreak:       userCalendarResp.Data.UserCalendar.RecentStreak,
	}

	recentACSubmissions := models.RecentACSubmission{
		UserSlug:           user.UserSlug,
		SubmissionID:       recentACSubmissionsResp.Data.RecentACSubmission[0].SubmissionID,
		SubmitTime:         recentACSubmissionsResp.Data.RecentACSubmission[0].SubmitTime,
		QuestionFrontendID: recentACSubmissionsResp.Data.RecentACSubmission[0].Question.QuestionFrontendID,
	}

	question := models.Question{
		Title:              recentACSubmissionsResp.Data.RecentACSubmission[0].Question.Title,
		TranslatedTitle:    recentACSubmissionsResp.Data.RecentACSubmission[0].Question.TranslatedTitle,
		TitleSlug:          recentACSubmissionsResp.Data.RecentACSubmission[0].Question.TitleSlug,
		QuestionFrontendID: recentACSubmissionsResp.Data.RecentACSubmission[0].Question.QuestionFrontendID,
	}

	db := data.Get()
	db.Where(models.User{UserSlug: userSlug}).FirstOrCreate(&user)
	db.Where(models.User{UserSlug: userSlug}).FirstOrCreate(&userCalendar)
	db.Where(models.User{UserSlug: userSlug}).FirstOrCreate(&recentACSubmissions)
	db.FirstOrCreate(&question)

	return nil
}

func GetUserProfile(userSlug string) (api.User, error) {
	dbUser := models.User{}
	dbUserCalendar := models.UserCalendar{}
	dbRecentACSubmission := models.RecentACSubmission{}
	dbQuestion := models.Question{}

	db := data.Get()
	db.Where("user_slug = ?", userSlug).First(&dbUser)
	db.Where("user_slug = ?", userSlug).First(&dbUserCalendar)
	db.Where("user_slug = ?", userSlug).First(&dbRecentACSubmission)
	db.Where("question_frontend_id = ?", dbRecentACSubmission.QuestionFrontendID).First(&dbQuestion)

	return api.User{
		UserSlug:         dbUser.UserSlug,
		RealName:         dbUser.RealName,
		AboutMe:          dbUser.AboutMe,
		UserAvatar:       dbUser.UserAvatar,
		Gender:           dbUser.Gender,
		Websites:         dbUser.Websites,
		SkillTags:        dbUser.SkillTags,
		IPRegion:         dbUser.IPRegion,
		Location:         dbUser.Location,
		UseDefaultAvatar: dbUser.UseDefaultAvatar,
		UserCalendar:     dbUserCalendar,
		RecentACSubmission: struct {
			RecentACSubmission models.RecentACSubmission
			Question           models.Question
		}{
			RecentACSubmission: dbRecentACSubmission,
			Question:           dbQuestion,
		},
	}, nil
}
