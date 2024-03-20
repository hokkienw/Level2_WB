package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	urlFlag := flag.String("u", "", "URL сайта для скачивания")
	outputFlag := flag.String("o", "index.html", "Имя файла для сохранения")

	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("Не указан URL сайта")
		os.Exit(1)
	}

	response, err := http.Get(*urlFlag)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: %s\n", response.Status)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		os.Exit(1)
	}

	dir := "."
	if idx := strings.LastIndex(*outputFlag, "/"); idx != -1 {
		dir = (*outputFlag)[:idx]
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Ошибка при создании директории:", err)
			os.Exit(1)
		}
	}

	err = ioutil.WriteFile(*outputFlag, data, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		os.Exit(1)
	}

	fmt.Printf("Сайт успешно скачан и сохранен в файл %s\n", *outputFlag)
}
