package props

import "errors"
import "os"
import "strings"

// Parses a java-style properties file and returns a map of those properties.
func Load(fileName string) (map[string]string, error) {
	if fileName == "" {
		return nil, errors.New("No properties file specified.")
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()

	propsMap := make(map[string]string)
	strs := strings.Split(string(buf), "\n")
	for _, str := range strs {
		if str == "" {
			break
		}

		if str[0] == '#' {
			// ignore it, it's a comment
			continue
		}

		s := strings.Split(str, "=")
		if len(s) != 2 {
			// just ignore it, invalid prop
			// TODO: consider logging something here
			continue
		}

		key := s[0]
		value := s[1]

		propsMap[key] = value
	}

	return propsMap, nil
}
