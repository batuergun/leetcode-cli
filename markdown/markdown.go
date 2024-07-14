package markdown

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseReadme(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var slugs []string
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`- \[ \] \d+\.(.+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			slug := strings.TrimSpace(matches[1])
			slug = strings.ReplaceAll(slug, " ", "-")
			slug = strings.ToLower(slug)
			slugs = append(slugs, slug)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return slugs, nil
}
