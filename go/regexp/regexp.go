package main

import (
	"crypto/sha1"
	"fmt"
	"html"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	keywords, _ := ioutil.ReadFile("./keywords.txt")

	for n := 1; n < 7001; n++ {
		content, err := ioutil.ReadFile("./descriptions/" + strconv.Itoa(n) + ".txt")
		panicIf(err)

		result := Htmlify3(string(content), strings.Split(strings.TrimSpace(string(keywords)), "\n"))
		err = ioutil.WriteFile("./descriptions_html/"+strconv.Itoa(n)+".txt", []byte(result), os.ModePerm)
		panicIf(err)
	}
}

func Htmlify(content string, keywords []string) string {
	re := regexp.MustCompile("(" + strings.Join(keywords, "|") + ")")
	kw2sha := make(map[string]string)
	content = re.ReplaceAllStringFunc(content, func(kw string) string {
		kw2sha[kw] = "isuda_" + fmt.Sprintf("%x", sha1.Sum([]byte(kw)))
		return kw2sha[kw]
	})
	content = html.EscapeString(content)
	for kw, hash := range kw2sha {
		u, err := url.Parse("http://localhost/keyword/" + pathURIEscape(kw))
		panicIf(err)
		link := fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(kw))
		content = strings.Replace(content, hash, link, -1)
	}
	return strings.Replace(content, "\n", "<br />\n", -1)
}

func Htmlify2(content string, re *regexp.Regexp) string {
	content = html.EscapeString(content)

	content = re.ReplaceAllStringFunc(content, func(kw string) string {
		u, err := url.Parse("http://localhost/keyword/" + pathURIEscape(kw))
		panicIf(err)
		return fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(kw))
	})
	return strings.Replace(content, "\n", "<br />\n", -1)
}

func Htmlify3(content string, keywords []string) string {
	reps := []string{}

	for _, kw := range keywords {
		reps = append(reps, kw)
		u, err := url.Parse("http://localhost/keyword/" + pathURIEscape(kw))
		panicIf(err)
		reps = append(reps, fmt.Sprintf("<a href=\"%s\">%s</a>", u, kw))
	}

	r := strings.NewReplacer(reps...)

	content = html.EscapeString(content)
	content = r.Replace(content)

	return strings.Replace(content, "\n", "<br />\n", -1)
}

func pathURIEscape(s string) string {
	return (&url.URL{Path: s}).String()
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
