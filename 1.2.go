package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	// 低效版本
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("低效版本耗时%.6f秒\n", time.Since(start).Seconds())
	fmt.Printf("结果为%s\n", s)
	start = time.Now()
	alsoS := strings.Join(os.Args[1:], " ")
	fmt.Printf("高效版本耗时%.6f秒\n", time.Since(start).Seconds())
	fmt.Printf("结果为%s\n", alsoS)
}
