package timer

import "time"

func CurrentTime() (ctime string) {

	timeUnix := time.Now().Unix()
	ctime = time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	return

}
