run:
	kubectl create configmap experiment-config --from-file nginx.conf
	kubectl apply -f staging-service.yml
	kubectl apply -f prod-service.yml
	kubectl apply -f ambassador.yml
.PHONY: run

delete:
	kubectl delete -f ambassador.yml; :
	kubectl delete -f prod-service.yml; :
	kubectl delete -f staging-service.yml; :
	kubectl delete configmap experiment-config; :
.PHONY: delete
