all:
	clear
	go install -a
	go run ./test/test.go