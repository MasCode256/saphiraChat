package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

func main () {
	http.HandleFunc("/", handler)

	fmt.Println("Подготовка UI...")
	// Запускаем сервер на localhost:8080
	err := http.ListenAndServe("127.0.0.1:1110", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Пример: получение значения параметра "name"
	if queryParams.Has("test") {
		// Если параметр существует, получаем его значение
		prompt := queryParams.Get("test")
		fmt.Fprintf(w, "Параметр 'test' найден: %s\n", prompt)
	} else if queryParams.Has("get") {
		prompt := queryParams.Get("get")
		if prompt == "contacts" {
			intValue := countFilesInDirectory("go\\data\\contacts") 
			intStr := strconv.Itoa(intValue)
    		fmt.Fprintln(w, intStr)
		}
	} else if queryParams.Has("name") {
		fmt.Println("Request result:", executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("name") + ".json name"))
		fmt.Fprintln(w, "" + executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("name") + ".json name"))
	} else if queryParams.Has("ip") {
		fmt.Println("Request result:", executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("ip") + ".json ip"))
		fmt.Fprintln(w, "" + executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("ip") + ".json ip"))
	} else if queryParams.Has("color") {
		fmt.Println("Request result:", executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("color") + ".json color"))
		fmt.Fprintln(w, "" + executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("color") + ".json color"))
	} else if queryParams.Has("key") {
		fmt.Println("Request result:", executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("key") + ".json key"))
		fmt.Fprintln(w, "" + executeCommand("go\\json_sys\\json_sys.exe get_value go\\data\\contacts\\" + queryParams.Get("key") + ".json key"))
	} else {
		fmt.Fprintln(w, "error:21:Параметры QK-0 не найдены.")
	}
}


func executeCommand(fullCommand string) string {
	parts := strings.Fields(fullCommand)
	if len(parts) == 0 {
		return ""
	}

	// Первая часть - это команда, остальные - аргументы
	command := parts[0]
	args := parts[1:]

	// Создаем команду с помощью exec.Command
	cmd := exec.Command(command, args...)

	// Используем bytes.Buffer для захвата стандартного вывода
	var out bytes.Buffer
	cmd.Stdout = &out

	// Выполняем команду
	err := cmd.Run()
	if err != nil {
		return ""
	}

	// Возвращаем вывод команды в виде строки
	return out.String()
}

func countFilesInDirectory(directory string) int {
	// Создаем полную команду в виде одной строки
	fullCommand := fmt.Sprintf("cmd /C dir %s /A-D /B", directory)

	// Выполняем команду
	output := executeCommand(fullCommand)

	// Разбиваем вывод на строки и считаем количество строк
	lines := strings.Split(strings.TrimSpace(output), "\n")
	return len(lines)
}