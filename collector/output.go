package collector

import (
	"fmt"
	"os"
)

// WritePoints formats the given points as InfluxDB line protocol and writes
// them to stdout. Each point is written on its own line. Stdout is synced
// after writing to ensure telegraf receives the data promptly.
func WritePoints(points []Point) error {
	for _, p := range points {
		if len(p.Fields) == 0 {
			continue
		}
		_, err := fmt.Fprintln(os.Stdout, FormatLineProtocol(p))
		if err != nil {
			return err
		}
	}
	os.Stdout.Sync()
	return nil
}
