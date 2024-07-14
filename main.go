package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Baticaly/leetcode-cli/leetcode"
	"github.com/Baticaly/leetcode-cli/markdown"
	"github.com/fatih/color"
)

func main() {
	errorColor := color.New(color.FgRed).PrintfFunc()
	successColor := color.New(color.FgGreen).PrintfFunc()
	infoColor := color.New(color.FgBlue).PrintfFunc()
	warningColor := color.New(color.FgYellow).PrintfFunc()

	infoColor("Reading problem slugs from README.md...\n")
	problemSlugs, err := markdown.ParseReadme("README.md")
	if err != nil {
		errorColor("Error reading README.md: %v\n", err)
		return
	}

	if len(problemSlugs) == 0 {
		warningColor("No problems to fetch.\n")
		return
	}

	infoColor("Fetching problems from LeetCode...\n")
	problems, err := leetcode.FetchProblems()
	if err != nil {
		errorColor("Error fetching problems: %v\n", err)
		return
	}

	problemMap := make(map[string]leetcode.Problem)
	for _, problem := range problems {
		problemMap[problem.Stat.TitleSlug] = problem
	}

	for _, slug := range problemSlugs {
		problem, exists := problemMap[slug]
		if !exists {
			warningColor("Problem with slug '%s' not found.\n", slug)
			continue
		}
		if problem.PaidOnly {
			warningColor("Problem with slug '%s' is paid only.\n", slug)
			continue
		}

		infoColor("Fetching details for problem '%s'...\n", slug)
		details, err := leetcode.FetchProblemDetails(slug)
		if err != nil {
			errorColor("Error fetching details for problem '%s': %v\n", slug, err)
			continue
		}

		problemPath := filepath.Join("problems", fmt.Sprintf("%s.%s", details.Data.Question.QuestionID, details.Data.Question.Title))
		if err := os.MkdirAll(problemPath, os.ModePerm); err != nil {
			errorColor("Error creating directory '%s': %v\n", problemPath, err)
			continue
		}

		if err := leetcode.SaveProblem(details, problemPath); err != nil {
			errorColor("Error saving problem: %v\n", err)
		} else {
			successColor("Problem '%s' saved successfully.\n", slug)
		}
	}

	successColor("Specified problems fetched and saved successfully.\n")
}
