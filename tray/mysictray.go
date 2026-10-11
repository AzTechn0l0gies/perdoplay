package tray

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

func MusicTray(n int) {
	bar := progressbar.Default(int64(n))
	for i := 0; i < n; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}
