# build -> build application
build:
	go build -o main ./main.go

# run -> application
run:
	./main

# dev -> run build then run it
dev: 
	watchexec -r -c -e go -- just build run

# docs -> generate swagger docs
docs:
	swag init

# docs-csv -> generate swagger docs then convert it into csv
docs-csv:
	just docs
	python3 exporter.py ./docs/swagger.json