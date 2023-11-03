package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lonzzi/leetcode-daily-rank/pkg/models"
	"io"
	"net/http"
)

func sendGraphQLRequest(apiURL string, query string, variables map[string]interface{}) ([]byte, error) {
	requestBody := models.GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GraphQL request failed with status code: %d", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetUserProfileCalendar(userSlug string) (*models.UserProfileCalendarResp, error) {
	apiURL := "https://leetcode.cn/graphql/noj-go"
	query := `
    query userProfileCalendar($userSlug: String!, $year: Int) {
		userCalendar(userSlug: $userSlug, year: $year) {
		  streak
		  totalActiveDays
		  submissionCalendar
		  activeYears
		  monthlyMedals {
			name
			obtainDate
			category
			config {
			  icon
			  iconGif
			  iconGifBackground
			}
			progress
			id
			year
			month
		  }
		  recentStreak
		}
	  }`

	variables := map[string]interface{}{
		"userSlug": userSlug,
	}

	response, err := sendGraphQLRequest(apiURL, query, variables)
	if err != nil {
		return nil, err
	}

	var resp models.UserProfileCalendarResp
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetRecentAcSubmissions(userSlug string) (*models.RecentACSubmissionsResp, error) {
	apiURL := "https://leetcode.cn/graphql/noj-go"
	query := `
    query recentAcSubmissions($userSlug: String!) {
		recentACSubmissions(userSlug: $userSlug) {
		  submissionId
		  submitTime
		  question {
			title
			translatedTitle
			titleSlug
			questionFrontendId
		  }
		}
	  }`

	variables := map[string]interface{}{
		"userSlug": userSlug,
	}

	response, err := sendGraphQLRequest(apiURL, query, variables)
	if err != nil {
		return nil, err
	}

	var resp models.RecentACSubmissionsResp
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetUserProfilePublicProfile(userSlug string) (*models.UserProfileResp, error) {
	apiURL := "https://leetcode.cn/graphql"
	query := `
    query userProfilePublicProfile($userSlug: String!) {
		userProfilePublicProfile(userSlug: $userSlug) {
		  haveFollowed
		  siteRanking
		  profile {
			userSlug
			realName
			aboutMe
			userAvatar
			gender
			websites
			skillTags
			ipRegion
			birthday
			location
			useDefaultAvatar
			github
			school: schoolV2 {
			  schoolId
			  logo
			  name
			}
			company: companyV2 {
			  id
			  logo
			  name
			}
			job
			globalLocation {
			  country
			  province
			  city
			  overseasCity
			}
			socialAccounts {
			  provider
			  profileUrl
			}
			skillSet {
			  langLevels {
				langName
				langVerboseName
				level
			  }
			  topics {
				slug
				name
				translatedName
			  }
			  topicAreaScores {
				score
				topicArea {
				  name
				  slug
				}
			  }
			}
		  }
		  educationRecordList {
			unverifiedOrganizationName
		  }
		  occupationRecordList {
			unverifiedOrganizationName
			jobTitle
		  }
		}
	  }`

	variables := map[string]interface{}{
		"userSlug": userSlug,
	}

	response, err := sendGraphQLRequest(apiURL, query, variables)
	if err != nil {
		return nil, err
	}

	var resp models.UserProfileResp
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
