get:
	kubectl get all
.PHONY: get

run:
	kubectl apply -f app-deploy.yml; :
	kubectl create configmap nginx-config --from-file nginx.conf; :
	kubectl apply -f nginx-service.yml; :
.PHONY: run

delete:
	kubectl delete -f nginx-service.yml; :
	kubectl delete configmap nginx-config; :
	kubectl delete -f app-deploy.yml; :
.PHONY: delete
