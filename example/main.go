package main

import (
	"github.com/KlyuchnikovV/logist"
	"github.com/KlyuchnikovV/logist/internal/types"
)

func main() {
	logger, err := logist.New(
		logist.WithLevel(types.WarningLevel),
	)
	if err != nil {
		panic(err)
	}
	defer logger.Stop()

	logger.Info("info log")
}
