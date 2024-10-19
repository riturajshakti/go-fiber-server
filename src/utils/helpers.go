package utils

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Contains(line, "=") {
			keyValue := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])

			if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
				value = strings.Trim(value, "\"")
			} else if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
				value = strings.Trim(value, "'")
			}

			switch key {
			case "PORT":
				Env.Port = value
			case "HOST":
				Env.Host = value
			}
		}
	}
	return scanner.Err()
}
