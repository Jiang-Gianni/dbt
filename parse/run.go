package parse

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

const (
	Name  = "name"
	Test  = "test"
	Bench = "bench"
)

var TypeList = [3]string{Name, Test, Bench}

type Scanner struct {
	Line string

	// Map: name/alias -> query
	Queries map[string]string

	// Map: ["name", "test", "bench"] -> query list
	MapList map[string][]string

	CurrentName string
	CurrentType string
}

type stateFn func(*Scanner) stateFn

func (s *Scanner) getNameTag() bool {
	for _, t := range TypeList {
		re := regexp.MustCompile(fmt.Sprintf(`^\s*--\s*%s:\s*(\S+)`, t))
		matches := re.FindStringSubmatch(s.Line)
		if matches != nil {
			s.MapList[t] = append(s.MapList[t], matches[1])
			s.CurrentType = t
			s.CurrentName = matches[1]
			return true
		}
	}
	return false
}

func initialState(s *Scanner) stateFn {
	if s.getNameTag() {
		return queryState
	}
	return initialState
}

func queryState(s *Scanner) stateFn {
	if !s.getNameTag() {
		s.appendQueryLine()
	}
	return queryState
}

func (s *Scanner) appendQueryLine() {
	current := s.Queries[s.CurrentName]
	line := strings.Trim(s.Line, " \t")
	if len(line) == 0 {
		return
	}
	if len(current) > 0 {
		current = current + "\n"
	}
	current = current + line
	s.Queries[s.CurrentName] = current
}

func (s *Scanner) Run(io *bufio.Scanner) map[string]string {
	for state := initialState; io.Scan(); {
		s.Line = io.Text()
		state = state(s)
	}
	return s.Queries
}
