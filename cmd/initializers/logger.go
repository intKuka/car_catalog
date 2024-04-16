package initializers

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func HandleLogging() {
	logOptions := &slog.HandlerOptions{Level: slog.LevelDebug}

	Log = slog.New(
		slog.NewTextHandler(os.Stdout, logOptions),
	)
}
