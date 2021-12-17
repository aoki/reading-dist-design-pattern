image:
	docker build -t go-server:latest .

run:
	docker run -it --rm -p 8080:8080 go-server

sidecar:
	docker run -d --rm -p 8181:8080 --name go-server go-server
	docker run -d --rm --pid=container:go-server -p 8080:8080 brendanburns/topz:db0fa58 /server -addr=0.0.0.0:8080
	open http://localhost:8080/topz
