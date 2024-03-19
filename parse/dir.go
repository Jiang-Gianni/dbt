package parse

import (
	"bufio"
	"os"
	"regexp"
)

var sqlRegexp = regexp.MustCompile(".sql$")

// New returns a scanner after analyzing the input directory
func New(dir string) (*Scanner, error) {
	s := &Scanner{
		Queries: make(map[string]*QueryTest),
		MapList: make(map[string][]string),
		LineMap: make(map[string]int),
	}
	return s, s.ParseDir(dir)
}

// ParseDir scans the entire input directory and subdirectory. All the found ".sql" files are then analyzed.
func (s *Scanner) ParseDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {

		var entryName string
		if dir == "./" {
			entryName = "./" + entry.Name()
		} else {
			entryName = dir + "/" + entry.Name()
		}

		if sqlRegexp.MatchString(entryName) {
			f, err := os.Open(entryName)
			if err != nil {
				return err
			}
			defer f.Close()
			b := bufio.NewScanner(f)
			s.Run(entryName, b)
			continue
		}

		if entry.IsDir() {
			if err := s.ParseDir(entryName); err != nil {
				return err
			}
		}
	}
	return nil
}
