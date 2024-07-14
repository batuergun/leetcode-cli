package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const (
	leetcodeAPI       = "https://leetcode.com/api/problems/all/"
	leetcodeGraphQLAPI = "https://leetcode.com/graphql"
)

type Problem struct {
	Stat struct {
		QuestionID int    `json:"question_id"`
		Title      string `json:"question__title"`
		TitleSlug  string `json:"question__title_slug"`
	} `json:"stat"`
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
	PaidOnly bool `json:"paid_only"`
}

type LeetCodeResponse struct {
	StatStatusPairs []Problem `json:"stat_status_pairs"`
}

type GraphQLRequest struct {
	Query     string `json:"query"`
	Variables struct {
		TitleSlug string `json:"titleSlug"`
	} `json:"variables"`
}

type GraphQLResponse struct {
	Data struct {
		Question struct {
			Title       string `json:"title"`
			Content     string `json:"content"`
			Difficulty  string `json:"difficulty"`
			QuestionID  string `json:"questionId"`
			TitleSlug   string `json:"titleSlug"`
		} `json:"question"`
	} `json:"data"`
}

func FetchProblems() ([]Problem, error) {
	resp, err := http.Get(leetcodeAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leetCodeResp LeetCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&leetCodeResp); err != nil {
		return nil, err
	}

	return leetCodeResp.StatStatusPairs, nil
}

func FetchProblemDetails(slug string) (GraphQLResponse, error) {
	query := `
		query getQuestionDetail($titleSlug: String!) {
			question(titleSlug: $titleSlug) {
				title
				content
				difficulty
				questionId
				titleSlug
			}
		}
	`
	var gqlReq GraphQLRequest
	gqlReq.Query = query
	gqlReq.Variables.TitleSlug = slug

	reqBody, err := json.Marshal(gqlReq)
	if err != nil {
		return GraphQLResponse{}, err
	}

	resp, err := http.Post(leetcodeGraphQLAPI, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return GraphQLResponse{}, err
	}
	defer resp.Body.Close()

	var gqlResp GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&gqlResp); err != nil {
		return GraphQLResponse{}, err
	}

	return gqlResp, nil
}

func SaveProblem(problem GraphQLResponse) error {
	dir := fmt.Sprintf("problems/%s.%s", problem.Data.Question.QuestionID, problem.Data.Question.Title)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	readmeContent := fmt.Sprintf("## %s. %s (%s)\n\n", problem.Data.Question.QuestionID, problem.Data.Question.Title, problem.Data.Question.Difficulty)
	readmeContent += fmt.Sprintf("Link: [LeetCode](https://leetcode.com/problems/%s/)\n\n", problem.Data.Question.TitleSlug)
	readmeContent += problem.Data.Question.Content

	return os.WriteFile(filepath.Join(dir, "README.md"), []byte(readmeContent), 0644)
}
