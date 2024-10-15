chcp 65001
go build -o go/msg_sender/sender.exe go/msg_sender/sender.go
go build -o go/ui/ui.exe go/ui/ui.go
start /b "Local server for WebUI" "go/ui/ui.exe"
start /b "Local server for CS message sending" "go/msg_sender/sender.exe"
start /b "UI" "cmd /k npm start"