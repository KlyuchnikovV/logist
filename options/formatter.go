package options

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/KlyuchnikovV/logist/internal"
	"github.com/KlyuchnikovV/logist/internal/types"
)

func FormatJSON(entry types.Entry) (string, error) {
	bytes, err := json.Marshal(entry)
	if err != nil {
		return "", fmt.Errorf("can't format entry as JSON: %w", err)
	}

	return string(bytes), nil
}

func FormatString(delimiter rune, fields ...types.Key) internal.Formatter {
	var err error
	if delimiter == 0 || delimiter == '\n' {
		err = fmt.Errorf("delimiter is empty")
	}

	if len(fields) == 0 {
		err = fmt.Errorf("fields are empty")
	}

	return func(e types.Entry) (string, error) {
		if err != nil {
			return "", err
		}

		var result = strings.Builder{}

		for _, key := range fields {
			field, ok := e[key]
			if !ok {
				continue
			}

			result.WriteString(fmt.Sprintf("%s%c", field, delimiter))
		}

		return strings.Trim(result.String(), string(delimiter)), nil
	}
}
