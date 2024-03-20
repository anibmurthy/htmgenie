package mapper

import (
	"fmt"
	"regexp"
	"strings"
)

const URL = "url"
const REGEX = `\[([\w\s\d\\/#.<>]+)\]\(((?:\/|https?:\/\/)[\w\d./?=#]+)\)`

func Url(in string) string {
	// TODO: Need to refine this a bit more to consider other `link text` possibilities.
	re := regexp.MustCompile(REGEX)
	matches := re.FindAllStringSubmatch(in, -1)

	for _, match := range matches {
		if len(match) == 3 {
			// match found is captured in the order of occurance.
			// match[2] = "Link text" in ["Link text found here"]
			// match[3] = "url" in ("url found here")
			in = replaceUrl(in, match[1], match[2])
		}
	}

	return in
}

func replaceUrl(in string, text string, url string) string {
	old := fmt.Sprintf("[%s](%s)", text, url)

	return strings.Replace(in, old, fmt.Sprintf(`<a href="%s">%s</a>`, url, text), 1)
}
