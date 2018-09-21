package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/ulule/ancre/logging"
	"github.com/ulule/ancre/serializer"
)

// Info open the given file and display information.
func Info(logger logging.Logger, filepath string) error {
	if path.Ext(filepath) != ".ots" {
		return fmt.Errorf("%s is not an ots file", filepath)
	}

	logger.Info("opening file", logging.String("path", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	ots, err := serializer.TimestampFileFromReader(file)
	if err != nil {
		return err
	}

	return ots.Infos(os.Stdout)
}
