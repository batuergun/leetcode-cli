package main

import (
	"fmt"
	"leetcode-cli/leetcode"
	"leetcode-cli/markdown"
)

func main() {
	problemSlugs, err := markdown.ParseReadme("README.md")
	if err != nil {
		fmt.Println("Error reading README.md:", err)
		return
	}

	if len(problemSlugs) == 0 {
		fmt.Println("No problems to fetch.")
		return
	}

	problems, err := leetcode.FetchProblems()
	if err != nil {
		fmt.Println("Error fetching problems:", err)
		return
	}

	problemMap := make(map[string]leetcode.Problem)
	for _, problem := range problems {
		problemMap[problem.Stat.TitleSlug] = problem
	}

	for _, slug := range problemSlugs {
		problem, exists := problemMap[slug]
		if !exists {
			fmt.Printf("Problem with slug '%s' not found.\n", slug)
			continue
		}
		if problem.PaidOnly {
			fmt.Printf("Problem with slug '%s' is paid only.\n", slug)
			continue
		}

		details, err := leetcode.FetchProblemDetails(slug)
		if err != nil {
			fmt.Printf("Error fetching details for problem '%s': %v\n", slug, err)
			continue
		}

		if err := leetcode.SaveProblem(details); err != nil {
			fmt.Println("Error saving problem:", err)
		}
	}

	fmt.Println("Specified problems fetched and saved successfully.")
}
