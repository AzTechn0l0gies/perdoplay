package musicplay

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

func Playsound(pathh string) {
	f, err := os.Open(pathh)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// декодировали его
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Инициализация говорилки
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Стример содержит саму музыку спикер способ её воспроизвести
	speaker.PlayAndWait(streamer)
}

func GetSecond(pathh string) int64 {
	f, err := os.Open(pathh)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// декодировали его
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	duration := format.SampleRate.D(streamer.Len()).Seconds()
	return int64(duration)
}
