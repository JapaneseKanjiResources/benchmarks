package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestHtmlify(t *testing.T) {
	f, _ := ioutil.ReadFile("./keywords.txt")
	keywords := strings.Split(strings.TrimSpace((string(f))), "\n")
	re := regexp.MustCompile("(" + strings.Join(keywords, "|") + ")")

	for n := 1; n < 10; n++ {
		t.Logf("Test start for %d.txt", n)
		desc, err := ioutil.ReadFile("./descriptions/" + strconv.Itoa((n)) + ".txt")
		panicIf(err)

		expected, err := ioutil.ReadFile("./descriptions_html/" + strconv.Itoa((n)) + ".txt")
		panicIf(err)

		actual := Htmlify(string(desc), keywords)
		if actual != string(expected) {
			t.Errorf("Htmlify failed: %d", f)
		}

		actual = Htmlify2(string(desc), re)
		if actual != string(expected) {
			t.Errorf("Htmlify2 failed: %d", f)
		}

		actual = Htmlify3(string(desc), keywords)
		if actual != string(expected) {
			t.Errorf("Htmlify2 failed: %d", f)
		}
	}
}

func BenchmarkHtmlify(b *testing.B) {
	f, _ := ioutil.ReadFile("./keywords.txt")
	keywords := strings.Split(strings.TrimSpace((string(f))), "\n")

	for n := 1; n < 10; n++ {
		b.Logf("Benchmark start for %d.txt", n)
		desc, err := ioutil.ReadFile("./descriptions/" + strconv.Itoa((n)) + ".txt")
		panicIf(err)

		_ = Htmlify(string(desc), keywords)
	}
}

func BenchmarkHtmlify2(b *testing.B) {
	f, _ := ioutil.ReadFile("./keywords.txt")
	keywords := strings.Split(strings.TrimSpace((string(f))), "\n")
	re := regexp.MustCompile("(" + strings.Join(keywords, "|") + ")")

	for n := 1; n < 10; n++ {
		b.Logf("Benchmark start for %d.txt", n)
		desc, err := ioutil.ReadFile("./descriptions/" + strconv.Itoa((n)) + ".txt")
		panicIf(err)

		_ = Htmlify2(string(desc), re)
	}
}
func BenchmarkHtmlify3(b *testing.B) {
	f, _ := ioutil.ReadFile("./keywords.txt")
	keywords := strings.Split(strings.TrimSpace((string(f))), "\n")

	for n := 1; n < 10; n++ {
		b.Logf("Benchmark start for %d.txt", n)
		desc, err := ioutil.ReadFile("./descriptions/" + strconv.Itoa((n)) + ".txt")
		panicIf(err)

		_ = Htmlify3(string(desc), keywords)
	}
}
