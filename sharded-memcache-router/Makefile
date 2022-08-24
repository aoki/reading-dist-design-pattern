get:
	kubectl get all
.PHONY: get

run:
	kubectl apply -f memcached-shards.yml; :
	kubectl create configmap sharded-twem-config --from-file=./nutcracker.yml; :
	kubectl apply -f memcached-twemproxy-deploy.yml; :
.PHONY: run

delete:
	kubectl delete -f memcached-twemproxy-deploy.yml; :
	kubectl delete configmap sharded-twem-config; :
	kubectl delete -f memcached-shards.yml; :
.PHONY: delete
