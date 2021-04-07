package util

import (
	"fmt"
	"time"
)

func Runningtime(s string) (string, time.Time) {
	fmt.Println("Start:	", s)
	return s, time.Now()
}

func Track(s string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println("End:	", s, "took", endTime.Sub(startTime))
}
