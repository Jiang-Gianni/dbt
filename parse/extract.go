package parse

import (
	"fmt"
	"regexp"
)

func (s *Scanner) Extract(name string) (string, []any, error) {
	query := s.Queries[name].Query
	if query[0:2] != "--" {
		return query, nil, nil
	}
	f, args, err := ParseFunctionCall(query)
	return s.Queries[f].Query, args, err
}

func ParseFunctionCall(input string) (string, []any, error) {
	// Function should be like:
	// -- GetMember(4)
	// name: GetMember, args: [4]
	re := regexp.MustCompile(`^--.*?(\w+)\((.*?)\)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 3 {
		return "", nil, fmt.Errorf("invalid function call")
	}

	name := matches[1]
	argStr := matches[2]
	args := parseArguments(argStr)

	return name, args, nil
}

func parseArguments(argStr string) []any {
	var args []any
	var arg string
	var inString bool
	var inArray bool

	for _, char := range argStr {
		switch char {
		case ',':
			if !inString && !inArray {
				args = append(args, arg)
				arg = ""
			} else {
				arg += string(char)
			}
		case '\'':
			inString = !inString
		case '{', '}':
			if !inString {
				inArray = !inArray
			}
			arg += string(char)
		default:
			arg += string(char)
		}
	}
	args = append(args, arg)
	return args
}
