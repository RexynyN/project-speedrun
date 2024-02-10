package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Now()

	parser := "2024-02-10 17:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", parser)

	fmt.Println(t1)
	fmt.Println(t2)
	// https://stackoverflow.com/questions/17631900/something-like-python-timedelta-in-golang

	diff := t1.Sub(t2)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())

	// fmt.Println(-t2.Second() + t1.Second())
	// fmt.Println(printDiff(-t2.Second() + t1.Second()))

	// for {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(diff)
	// }
}

func printDiff(diff int) string {
	// hours, minutes, seconds := int(diff.Seconds()/3600), int(diff.Seconds()/60)%60, int(diff.Seconds())%60
	hours, minutes, seconds := int(diff/3600), int(diff/60)%60, int(diff)%60
	days := int(hours / 24)

	return fmt.Sprintf("%d Dias, %d Horas, %d Minutos, %d Segundos", days, hours, minutes, seconds)
}
