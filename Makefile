mac:
	go build -o mjhp_mac main/main.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mjhp_linux main/main.go