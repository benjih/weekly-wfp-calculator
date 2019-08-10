.PHONY: build

build:
	env GOOS=linux GOARCH=386 go build -o build/weekly-wfp-calculator-linux-386 main.go
	env GOOS=linux GOARCH=amd64 go build -o build/weekly-wfp-calculator-linux-amd64 main.go
	env GOOS=darwin GOARCH=amd64 go build -o build/weekly-wfp-calculator-mac-amd64 main.go
	env GOOS=windows GOARCH=386 go build -o build/weekly-wfp-calculator-windows-386.exe main.go
	env GOOS=windows GOARCH=amd64 go build -o build/weekly-wfp-calculator-windows-amd64.exe main.go