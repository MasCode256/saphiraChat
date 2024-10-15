package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// handler функция для обработки POST-запросов
func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		fmt.Println("Получены данные: ", http.StatusMethodNotAllowed)
        return }

    // Чтение тела запроса
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		fmt.Println("Получены данные: ", http.StatusInternalServerError)
        return }
    defer r.Body.Close()

	queryParams := r.URL.Query()

	if queryParams.Has("send") {
		fmt.Println("Отправка сообщения клиенту №" + queryParams.Get("send"))
	}

    fmt.Fprintf(w, "Получены данные: %s\n", string(body))
	fmt.Println("Получены данные: ", string(body))
}

func main() {
    http.HandleFunc("/", postHandler)

    fmt.Println("Подготовка систем отправки...")
    log.Fatal(http.ListenAndServe(":1112", nil))
}
