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
	Queries map[string]*QueryTest

	// Map: ["name", "test", "bench"] -> query name list
	MapList map[string][]string

	// name to line number
	LineMap map[string]int

	CurrentName string
	CurrentFile string
	CurrentLine int
}

type QueryTest struct {
	Query    string
	FileName string
	Line     int
}

type stateFn func(*Scanner) stateFn

func (s *Scanner) getNameTag() bool {
	for _, t := range TypeList {
		re := regexp.MustCompile(fmt.Sprintf(`^\s*--\s*%s:\s*(\S+)`, t))
		matches := re.FindStringSubmatch(s.Line)
		if matches != nil {
			s.MapList[t] = append(s.MapList[t], matches[1])
			s.Queries[matches[1]] = &QueryTest{
				FileName: s.CurrentFile,
				Line:     s.CurrentLine,
			}
			s.LineMap[matches[1]] = s.CurrentLine
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
	current := s.Queries[s.CurrentName].Query
	line := strings.Trim(s.Line, " \t")
	if len(line) == 0 {
		return
	}
	if len(current) > 0 {
		current = current + "\n"
	}
	current = current + line
	s.Queries[s.CurrentName].Query = current
}

func (s *Scanner) Run(filename string, io *bufio.Scanner) {
	s.CurrentFile = filename
	s.CurrentLine = 0
	for state := initialState; io.Scan(); {
		s.CurrentLine = s.CurrentLine + 1
		s.Line = io.Text()
		state = state(s)
	}
}
