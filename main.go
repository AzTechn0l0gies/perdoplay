package main

import (
	"fmt"
	"os"

	"perdoplay/musicplay"
	"perdoplay/tray"
)

func main() {
	// проверяем наличие пути до музыкального файла
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: Вы не ввели путь до музыки")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Ошибка: Неизвестные аргументы")
		os.Exit(1)
	}
	go tray.MusicTray(int(musicplay.GetSecond(os.Args[1])))
	musicplay.Playsound(os.Args[1])

}
