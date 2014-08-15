package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/felixge/asciitable"
	"github.com/quantastic/qgo"
)

const (
	FullTime = "2006-01-02 15:04"
	TimeOnly = "15:04"
)

func main() {
	client := qgo.NewClient("http://localhost:8080")
	times, err := client.Times()
	if err != nil {
		panic(err)
	}
	table := asciitable.NewTable()
	table.AddRow("Category", "Start", "End", "Duration", "Id")
	table.AddSeparator()
	var prevDay int
	for _, t := range times {
		t.Start = t.Start.Local()
		end := "--:--"
		if t.End != nil {
			end = t.End.Local().Format(TimeOnly)
		}
		var start string
		if t.Start.Day() != prevDay && prevDay != 0 {
			start = t.Start.Format(FullTime)
		} else {
			start = strings.Repeat(" ", len(FullTime)-len(TimeOnly)) + t.Start.Format(TimeOnly)
		}
		table.AddRow(
			CategoryString(t.Category.Name),
			start,
			end,
			DurationString(t.Duration()),
			t.Id,
		)
		prevDay = t.Start.Day()
	}
	fmt.Printf("%s\n", table)
}

func DurationString(d time.Duration) string {
	hours := d / time.Hour
	d -= hours * time.Hour
	min := d / time.Minute
	return fmt.Sprintf("%02d:%02d", hours, min)
}

func CategoryString(c []string) string {
	return strings.Join(c, ":")
}
