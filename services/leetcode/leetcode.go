package leetcode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/lonzzi/leetcode-daily-rank/data"
	"github.com/lonzzi/leetcode-daily-rank/models"
	"github.com/lonzzi/leetcode-daily-rank/models/api"
	lt "github.com/lonzzi/leetcode-daily-rank/pkg/leetcode"
)

func structToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Struct {
		typ := reflect.TypeOf(data)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name
			result[fieldName] = field.Interface()
		}
	}
	return result
}

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

	var submissionCalendar map[string]int

	err = json.Unmarshal([]byte(userCalendarResp.Data.UserCalendar.SubmissionCalendar), &submissionCalendar)
	if err != nil {
		return err
	}

	maxTime := int64(0)
	todaySubmissions := 0
	for k := range submissionCalendar {
		t, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return err
		}
		if t > maxTime {
			maxTime = t
		}
	}

	now := time.Now().Unix()
	currentDate := time.Unix(now, 0).Format("2006-01-02")
	checkDate := time.Unix(maxTime, 0).Format("2006-01-02")

	if currentDate != checkDate {
		todaySubmissions = 0
	} else {
		todaySubmissions = submissionCalendar[fmt.Sprintf("%d", maxTime)]
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
		TodaySubmissions: todaySubmissions,
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

	ret := db.First(&models.User{}, "user_slug = ?", userSlug)
	if ret.Error != nil {
		db.Create(&user)
	} else {

		ret = db.Model(&models.User{}).Where("user_slug = ?", userSlug).Updates(structToMap(user))
		if ret.Error != nil {
			return ret.Error
		}
	}

	ret = db.First(&models.UserCalendar{}, "user_slug = ?", userSlug)
	if ret.Error != nil {
		db.Create(&userCalendar)
	} else {
		ret = db.Model(&models.UserCalendar{}).Where("user_slug = ?", userSlug).Updates(userCalendar)
		if ret.Error != nil {
			return ret.Error
		}
	}

	ret = db.First(&models.RecentACSubmission{}, "user_slug = ?", userSlug)
	if ret.Error != nil {
		db.Create(&recentACSubmissions)
	} else {
		ret = db.Model(&models.RecentACSubmission{}).Where("user_slug = ?", userSlug).Updates(recentACSubmissions)
		if ret.Error != nil {
			return ret.Error
		}
	}

	ret = db.First(&models.Question{}, "question_frontend_id = ?", recentACSubmissions.QuestionFrontendID)
	if ret.Error != nil {
		db.Create(&question)
	} else {
		ret = db.Model(&models.Question{}).Where("question_frontend_id = ?", recentACSubmissions.QuestionFrontendID).Updates(question)
		if ret.Error != nil {
			return ret.Error
		}
	}

	// ret := db.Where(models.User{UserSlug: userSlug}).FirstOrCreate(&user)
	// if ret.Error != nil {
	// 	return ret.Error
	// }
	// ret = db.Where(models.User{UserSlug: userSlug}).FirstOrCreate(&userCalendar)
	// if ret.Error != nil {
	// 	return ret.Error
	// }
	// ret = db.Where(models.User{UserSlug: userSlug}).FirstOrCreate(&recentACSubmissions)
	// if ret.Error != nil {
	// 	return ret.Error
	// }
	// ret = db.FirstOrCreate(&question)
	// if ret.Error != nil {
	// 	return ret.Error
	// }

	return nil
}

func GetUserProfile(userSlug string) (*api.User, error) {
	dbUser := models.User{}
	dbUserCalendar := models.UserCalendar{}
	dbRecentACSubmission := models.RecentACSubmission{}
	dbQuestion := models.Question{}

	db := data.Get()
	ret := db.Where("user_slug = ?", userSlug).First(&dbUser)
	if ret.Error != nil {
		return nil, ret.Error
	}
	ret = db.Where("user_slug = ?", userSlug).First(&dbUserCalendar)
	if ret.Error != nil {
		return nil, ret.Error
	}
	ret = db.Where("user_slug = ?", userSlug).First(&dbRecentACSubmission)
	if ret.Error != nil {
		return nil, ret.Error
	}
	ret = db.Where("question_frontend_id = ?", dbRecentACSubmission.QuestionFrontendID).First(&dbQuestion)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return &api.User{
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
		TodaySubmissions: dbUser.TodaySubmissions,
		RecentACSubmission: struct {
			RecentACSubmission models.RecentACSubmission
			Question           models.Question
		}{
			RecentACSubmission: dbRecentACSubmission,
			Question:           dbQuestion,
		},
	}, nil
}

func GetUsersByRank() ([]*models.User, error) {
	dbUsers := []*models.User{}
	db := data.Get()
	ret := db.Order("today_submissions desc").Find(&dbUsers)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return dbUsers, nil
}
