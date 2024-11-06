package ytracker

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"ytlog/internal/config"
)

type YTracker struct {
	urlBase string
	cfg     *config.Config
}

type RequestBody struct {
	Query string `json:"query"`
}

func (yt *YTracker) setHeaders(r *http.Request) {
	r.Header.Add("Authorization", yt.cfg.OAuthToken)
	r.Header.Add("X-Org-ID", yt.cfg.OrgNumber)
}

func CreateYTracker(cfg *config.Config) *YTracker {
	return &YTracker{
		urlBase: "https://api.tracker.yandex.net/v2",
		cfg:     cfg,
	}
}

func (yt *YTracker) GetUsers() *[]User {
	url := yt.urlBase + "/users?perPage=500"
	request, err := http.NewRequest("GET", url, nil)
	yt.setHeaders(request)
	if err != nil {
		panic(err.Error())
	}
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		var users []User
		err = json.Unmarshal(bodyBytes, &users)
		if err != nil {
			panic(err.Error())
		}
		return &users
	}
	panic("YTracker not responding with 200 OK")
}

func unmarshall[T any](data *[]byte) []T {
	var items []T
	err := json.Unmarshal(*data, &items)
	if err != nil {
		panic(err.Error())
	}
	return items
}

func (yt *YTracker) doRequest(method string, url string, body *[]byte) (*[]byte, int) {

	var request *http.Request

	if body != nil {
		buffer := bytes.NewBuffer(*body)
		request, _ = http.NewRequest(method, url, buffer)
	} else {
		request, _ = http.NewRequest(method, url, nil)
	}

	yt.setHeaders(request)
	resp, _ := http.DefaultClient.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		panic("YTracker not responding with 200 OK")
	}

	pagesCount, _ := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
	bodyBytes, _ := io.ReadAll(resp.Body)
	return &bodyBytes, pagesCount
}

func (yt *YTracker) GetIssues() *[]Issue {
	url := yt.urlBase + "/issues/_search?page="

	body := RequestBody{
		Query: `Queue: ` + strings.Join(yt.cfg.Queues, ", ") + ` "Sort by": Updated DESC`,
	}

	bodyBytes, _ := json.Marshal(body)
	responseBody, pagesCount := yt.doRequest("POST", url+strconv.Itoa(1), &bodyBytes)
	result := unmarshall[Issue](responseBody)

	for pageNum := 2; pageNum <= pagesCount; pageNum++ {
		responseBody, _ := yt.doRequest("POST", url+strconv.Itoa(pageNum), &bodyBytes)
		issues := unmarshall[Issue](responseBody)
		result = append(result, issues...)
	}

	return &result
}

func (yt *YTracker) GetIssueChangelog(issueKey string) *[]IssueLog {
	url := yt.urlBase + "/issues/" + issueKey + "/changelog?page="

	responseBody, pagesCount := yt.doRequest("GET", url+strconv.Itoa(1), nil)
	result := unmarshall[IssueLog](responseBody)

	for pageNum := 2; pageNum <= pagesCount; pageNum++ {
		responseBody, _ := yt.doRequest("GET", url+strconv.Itoa(pageNum), nil)
		issuesLogs := unmarshall[IssueLog](responseBody)
		result = append(result, issuesLogs...)
	}

	return &result
}
