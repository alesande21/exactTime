package exactTime

import (
	"fmt"
	"github.com/beevik/ntp"
	_ "github.com/beevik/ntp"
)

func GetTime() (string, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return time.String(), nil
}
