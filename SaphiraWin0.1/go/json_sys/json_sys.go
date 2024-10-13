package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("error:11:Недостаточно аргументов.")
		return
	}

	if args[1] == "update_file" {
		err := updateJSONValue(args[2], args[3], args[4])

		if err != nil {
			fmt.Printf("error:21: %v\n", err)
		} else {
			fmt.Println("success:0")
		}
	}

	if (args[1] == "get_value") {
		content, err := os.ReadFile(args[2])
		if err != nil {
			fmt.Printf("error:21:Ошибка при чтении файла: %v\n", err)
			return
		}

		fmt.Print(getValue(string(content), args[3]))
	}

	if args[1] == "create_json" {
		createAndWriteFile(args[2])
	}
}

func updateJSONValue(filePath, key, newValue string) error {
	// Открываем JSON-файл
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл: %v", err)
	}

	// Парсим JSON в map
	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return fmt.Errorf("не удалось распарсить JSON: %v", err)
	}

	// Обновляем значение по ключу
	data[key] = newValue

	// Конвертируем обратно в JSON
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("не удалось конвертировать в JSON: %v", err)
	}

	// Записываем обновленный JSON обратно в файл
	if err := ioutil.WriteFile(filePath, updatedJSON, 0644); err != nil {
		return fmt.Errorf("не удалось записать в файл: %v", err)
	}

	return nil
}

func getValue(jsonStr, name string) string {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Print("error:22:Ошибка при разборе JSON:", err)
		return ""
	}

	// Получаем значение по ключу
	return result[name].(string)
}

func createAndWriteFile(filePath string) error {
	// Открываем файл с флагами для записи, создаем файл, если он не существует
	// и обнуляем его, если он уже существует.
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error:21:не удалось создать или открыть файл: %v", err)
	}
	defer file.Close()

 	_, err2 := file.WriteString("{}")
	if err2 != nil {
		return fmt.Errorf("error:22:не удалось записать в файл: %v", err)
	}

	return nil
}