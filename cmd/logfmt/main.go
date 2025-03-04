// This program takes the structured log output and makes it readable.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	reset = "\033[0m"

	// Colors
	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

const (
	timeFormat = "2006/01/02 15:04:05.000"
	timeZone   = "Asia/Taipei"
)

var service string

func init() {
	flag.StringVar(&service, "service", "", "filter which service to see")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)
	_ = syscall.Kill(os.Getppid(), syscall.SIGINT)
}

func main() {
	flag.Parse()
	var b strings.Builder

	service := strings.ToLower(service)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		m := make(map[string]any)
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			if service == "" {
				fmt.Println(s)
			}
			continue
		}

		// If a service filter was provided, check.
		if service != "" && strings.ToLower(m["service"].(string)) != service {
			continue
		}

		// Always having a trace id present in the logs is nice.
		traceID := "00000000000000000000000000000000"
		if v, ok := m["trace_id"]; ok {
			traceID = fmt.Sprintf("%v", v)
		}

		// Build out the know portions of the log in the preferred order.
		b.Reset()
		b.WriteString(fmt.Sprintf("%s: %s: %s: %s: %s: %s: ",
			m["service"],
			formatTime(m["time"]),
			m["file"],
			colorizeLevel(m["level"]),
			traceID,
			colorize(m["msg"], magenta),
		))

		// Add the rest of the keys ignoring the ones we already
		// added for the log.
		for k, v := range m {
			switch k {
			case "service", "time", "file", "level", "trace_id", "msg":
				continue
			}

			// It's nice to see the key[value] in this format
			// especially since map ordering is random.
			b.WriteString(fmt.Sprintf("%s[%+v]: ", k, v))
		}

		// Write the new log format, removing the last :
		out := b.String()
		fmt.Println(out[:len(out)-2])
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func colorize(s any, color int) string {
	if color == 0 {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\033[%dm%s%s", color, s, reset)
}

func colorizeLevel(lvlStr any) string {
	s := fmt.Sprintf("%v", lvlStr)

	switch s {
	case slog.LevelDebug.String():
		return colorize(s, darkGray)
	case slog.LevelInfo.String():
		return colorize(s, cyan)
	case slog.LevelWarn.String():
		return colorize(s, lightYellow)
	case slog.LevelError.String():
		return colorize(s, red)
	default:
		return colorize(s, 0)
	}
}

func formatTime(tStr any) string {
	t, err := time.Parse(time.RFC3339Nano, tStr.(string))
	if err != nil {
		return fmt.Sprintf("%v", tStr)
	}

	loc, _ := time.LoadLocation(timeZone)
	return t.In(loc).Format(timeFormat)
}
