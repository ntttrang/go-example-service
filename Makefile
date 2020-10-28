build:
	go build -o t-log-service main.go

run:
	./t-log-service

clean:
	rm t-log-service

start: build run