package fields

import "github.com/KlyuchnikovV/logist/internal/types"

func String(key, value string) types.Field {
	return types.Field{
		Name:  key,
		Value: value,
	}
}
