@echo off
chcp 65001
echo Для запуска этого скрипта у вас должны быть установлены компилятор Go, фреймворк electron и менеджер пакетов npm
echo Этот скрипт запустит компиляцию исходного кода, установленного вместе с программой
echo Начало компиляции...

go build -o go/msg_sender/sender.exe go/msg_sender/sender.go
go build -o go/ui/ui.exe go/ui/ui.go

echo Компиляция завершена. Запуск систем...

start /b "Local server for WebUI" "go/ui/ui.exe"
start /b "Local server for CS message sending" "go/msg_sender/sender.exe"
start /b "UI" "cmd /k npm start"