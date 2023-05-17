package fileutils

import (
	"bytes"
	"os"
)

type Replacement struct {
	From string
	To   string
}

func ReplaceInFile(filepath string, replacements []Replacement) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	output := file

	for _, replacement := range replacements {
		output = bytes.ReplaceAll(output, []byte(replacement.From), []byte(replacement.To))
	}

	if err = os.WriteFile(filepath, output, 0666); err != nil {
		return err
	}

	return nil
}
