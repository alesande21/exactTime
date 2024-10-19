package exactTime

import (
	"fmt"
	"github.com/beevik/ntp"
	_ "github.com/beevik/ntp"
	"os"
	"time"
)

var host string = "0.beevik-ntp.pool.ntp.org"

func GetTime(params ...string) int {
	if len(params) != 0 {
	}
	ntpTime, err := ntp.Time(host)
	if err != nil {
		strErr := fmt.Sprintf("Ошибка получения времени: %s", err)
		fmt.Fprint(os.Stderr, strErr)
		return -1
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("#########################")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

	return 0
}
