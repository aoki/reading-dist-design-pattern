run:
	kubectl apply -f dictionary-deploy.yml; :
	kubectl create configmap varnish-config --from-file=default.vcl; :
	kubectl apply -f varnish.yml; :
.PHONY: run

delete:
	kubectl delete -f varnish.yml; :
	kubectl delete configmap varnish-config; :
	kubectl delete -f dictionary-deploy.yml; :
.PHONY: delete
