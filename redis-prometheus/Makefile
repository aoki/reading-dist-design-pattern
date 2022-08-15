run:
	kubectl apply -f redis.yml
	kubectl create configmap prometheus-config --from-file conf/prometheus.yaml
	kubectl apply -f prometheus.yml
.PHONY: run

delete:
	kubectl delete -f prometheus.yml; :
	kubectl delete configmap prometheus-config; :
	kubectl delete -f redis.yml; :
.PHONY: delete
