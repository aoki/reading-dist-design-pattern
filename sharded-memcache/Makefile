run:
	kubectl apply -f memcached-shards.yml; :
	kubectl create configmap twem-config --from-file=./nutcracker.yml; :
	kubectl apply -f memcached-ambassador-pod.yml; :
.PHONY: run

delete:
	kubectl delete -f memcached-ambassador-pod.yml; :
	kubectl delete configmap twem-config; :
	kubectl delete -f memcached-shards.yml; :
.PHONY: delete 
