package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type CanvasService struct {
	ProviderPlatformID int
	Client             *http.Client
	BaseURL            string
	Token              string
	AccountID          string
	BaseHeaders        *map[string]string
	ClientID           string
	RedirectURI        string
}

func NewCanvasService(provider *ProviderPlatform) *CanvasService {
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + provider.ApiKey
	headers["Accept"] = "application/json"
	return &CanvasService{
		Client:      &http.Client{},
		BaseURL:     provider.Url,
		Token:       provider.ApiKey,
		AccountID:   provider.AccountID,
		BaseHeaders: &headers,
	}
}

func (srv *CanvasService) GetUsers() ([]UnlockEdImportUser, error) {
	url := srv.BaseURL + "/api/v1/accounts/" + srv.AccountID + "/users"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range *srv.BaseHeaders {
		req.Header.Add(key, value)
	}
	resp, err := srv.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	users := make([]map[string]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}
	unlockedUsers := make([]UnlockEdImportUser, 0)
	for _, user := range users {
		name := strings.Split(user["name"].(string), " ")
		nameFirst, nameLast := "", ""
		if len(name) < 2 {
			nameFirst = name[0]
			nameLast = user["short_name"].(string)
		} else {
			nameFirst = name[0]
			nameLast = name[1]
		}
		unlockedUser := UnlockEdImportUser{
			ExternalUserID:   user["id"].(string),
			ExternalUsername: user["login_id"].(string),
			NameFirst:        nameFirst,
			NameLast:         nameLast,
			Email:            user["login_id"].(string),
			Username:         nameLast + nameFirst,
		}
		unlockedUsers = append(unlockedUsers, unlockedUser)
	}
	return unlockedUsers, nil
}

func (srv *CanvasService) GetContent() ([]UnlockEdImportContent, error) {
	url := srv.BaseURL + "/api/v1/accounts/" + srv.AccountID + "/courses?include[]=course_image&include[]=public_description"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range *srv.BaseHeaders {
		req.Header.Add(key, value)
	}
	resp, err := srv.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	courses := make([]map[string]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&courses)
	if err != nil {
		return nil, err
	}
	unlockedCourses := make([]UnlockEdImportContent, 0)
	for _, course := range courses {
		unlockedCourse := UnlockEdImportContent{
			ProviderPlatformID: srv.ProviderPlatformID,
			Name:               course["name"].(string),
			ExternalContentID:  course["id"].(string),
			Description:        course["description"].(string),
			CourseCode:         course["course_code"].(string),
			IsOpenEnrollment:   course["is_public"].(bool),
			IsOpenContent:      false,
			HasAssessments:     true,
			Subject:            course["course_code"].(string),
			ImgURL:             course["image_download_url"].(string),
		}
		unlockedCourses = append(unlockedCourses, unlockedCourse)
	}
	return unlockedCourses, nil
}
