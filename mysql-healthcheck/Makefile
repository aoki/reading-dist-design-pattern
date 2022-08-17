image:
	docker build -t mysql-healthcheck:latest .

run:
	docker run -it --rm -p 9090:8080 mysql-healthcheck -query="SELECT * from COLUMNS" -user="root" -password="root" -database=information_schema

