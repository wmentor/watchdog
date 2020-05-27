package watchdog

import (
	"testing"
	"time"
)

func TestWD(t *testing.T) {

	wd := New(2, nil)
	defer wd.Close()

	for i := 0 ; i < 3 ; i++ {
		wd.Alive()
		time.Sleep(time.Second)
	}

}