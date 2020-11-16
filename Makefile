configure:
	go mod init
	goimports -w ./*.go

compile:
	goimports -w ./*.go
	go vet ./
	golint
	go test
	go install