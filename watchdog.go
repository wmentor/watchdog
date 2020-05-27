package watchdog

import (
	"context"
	"os"
	"time"
)

type WatchDog struct {
	deadline int64
	ttl      int64
	stop     context.CancelFunc
	killFunc func()
}

func New(ttl int64, killFunc func()) *WatchDog {

	wd := &WatchDog{
		deadline: time.Now().Unix() + ttl,
		ttl:      ttl,
		killFunc: killFunc,
	}

	ctx, fn := context.WithCancel(context.Background())

	wd.stop = fn

	go func() {

		for {

			select {

			case <-time.After(time.Second):

				if wd.deadline < time.Now().Unix() {
					if wd.killFunc != nil {
						wd.killFunc()
						break
					} else {
						os.Exit(1)
					}
				}

			case <-ctx.Done():
				return

			}

		}

	}()

	return wd
}

func (wd *WatchDog) Alive() {
	wd.deadline = time.Now().Unix() + wd.ttl
}

func (wd *WatchDog) Close() {
	wd.Alive()
	wd.stop()
}
