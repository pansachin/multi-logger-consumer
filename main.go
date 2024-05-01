package main

import (
	"github.com/pansachin/multi-logger-poc/plog"
	"github.com/pansachin/multi-logger-poc/slog"
	"github.com/pansachin/multi-logger-poc/zerolog"
)

func main() {
	slog.Log()
	plog.Log()
	zerolog.Log()
}
