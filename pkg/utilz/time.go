package utilz

import (
	"fmt"
	"time"
)

func MeasureTime(start time.Time, stageName string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took: %s\n", stageName, elapsed)
}
