image:
	docker build -t go-server:latest .

run:
	docker run -it --rm -p 8080:8080 go-server
