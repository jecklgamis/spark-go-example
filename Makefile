build:
	go build -o bin/spark_connect_app cmd/spark_connect_app.go
	chmod +x bin/spark_connect_app
install-deps:
	go mod tidy
run:
	bin/spark_connect_app
