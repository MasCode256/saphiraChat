package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
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
	} else if queryParams.Has("new") {
		prompt := queryParams.Get("new")
		if prompt == "contact" {
			executeCommand("go\\json_sys\\json_sys.exe create_json go\\data\\contacts\\" + int_to_str(countFilesInDirectory("go\\data\\contacts")) + ".json")

			executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + int_to_str(countFilesInDirectory("go\\data\\contacts") - 1) + ".json name Новый_контакт")
			executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + int_to_str(countFilesInDirectory("go\\data\\contacts") - 1) + ".json ip undefined_ip")
			executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + int_to_str(countFilesInDirectory("go\\data\\contacts") - 1) + ".json color #ffffff")
			executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + int_to_str(countFilesInDirectory("go\\data\\contacts") - 1) + ".json key unset")

			fmt.Fprintf(w, "sucess:0")
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
	} else if queryParams.Has("set_name") {
		executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + before(queryParams.Get("set_name"), '/') + ".json name " + after(queryParams.Get("set_name"), '/'))

		fmt.Fprintln(w, "sucess:0")

	} else if queryParams.Has("set_ip") {
		executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + before(queryParams.Get("set_ip"), '/') + ".json ip " + after(queryParams.Get("set_ip"), '/'))

		fmt.Fprintln(w, "sucess:0")

	} else if queryParams.Has("set_key") {
		executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + before(queryParams.Get("set_key"), '/') + ".json key " + after(queryParams.Get("set_key"), '/'))

		fmt.Fprintln(w, "sucess:0")

	} else if queryParams.Has("set_color") {
		executeCommand("go\\json_sys\\json_sys.exe update_file go\\data\\contacts\\" + before(queryParams.Get("set_color"), '/') + ".json color #" + after(queryParams.Get("set_color"), '/'))

		fmt.Fprintln(w, "sucess:0")

	} else if queryParams.Has("delete_contact") {
	 	removeFile("go\\data\\contacts\\" + queryParams.Get("delete_contact") + ".json")

		 for i := str_to_int(queryParams.Get("delete_contact")) + 1; i <= countFilesInDirectory("go\\data\\contacts"); i++ {
			src := "go\\data\\contacts\\" + int_to_str(i) + ".json"
			dst := "go\\data\\contacts\\" + int_to_str(i-1) + ".json"
			fmt.Println("Copying", src, "to", dst)
			
			// Убедитесь, что файл назначения не существует или удалите его
			os.Remove(dst)
		
			err := CopyFile(src, dst)
			if err != nil {
				fmt.Println("Ошибка копирования файла:", err)
				return
			}
		}
		

		fmt.Fprintln(w, "sucess:0:Успешное удаление контакта.")
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

func int_to_str(num int) string {
	ret := strconv.Itoa(num)
	return ret
}

func str_to_int(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

func removeFile(filePath string) error {
	// Используем os.Remove для удаления файла.
	err := os.Remove(filePath)
	if err != nil {
		return err}
	return nil
}


func CopyFile(src, dst string) error {
    // Открываем исходный файл
    sourceFile, err := os.Open(src)
    if err != nil {
        return err }
    defer sourceFile.Close()

    if _, err := os.Stat(dst); err == nil {
        return fmt.Errorf("файл назначения уже существует")
    } else if !os.IsNotExist(err) {
        return err }

    // Создаем файл назначения
    destinationFile, err := os.Create(dst)
    if err != nil {
        return err }
    defer destinationFile.Close()

    // Копируем данные из исходного файла в файл назначения
    _, err = io.Copy(destinationFile, sourceFile)
    if err != nil {
        return err }

    // Опционально: копируем права доступа
    sourceInfo, err := os.Stat(src)
    if err != nil {
        return err
    }
    err = os.Chmod(dst, sourceInfo.Mode())
    if err != nil {
        return err }

    return nil
}


func before(str string, c byte) string {
	// Используем strings.IndexByte для поиска первого вхождения символа c
	index := strings.IndexByte(str, c)
	if index == -1 {
// Если символ не найден, возвращаем всю строку
		return str
	}
	// Возвращаем подстроку от начала до найденного индекса
	return str[:index]
}


func after(str string, c byte) string {
	// Используем strings.IndexByte для поиска первого вхождения символа c
	index := strings.IndexByte(str, c)
	if index == -1 {
		// Если символ не найден, возвращаем пустую строку
		return ""
	}
	// Возвращаем подстроку от символа, следующего за найденным индексом, до конца строки
	return str[index+1:]
}

