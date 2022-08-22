run:
	kubectl apply -f dictionary-deploy.yml; :
	kubectl create configmap varnish-config --from-file=default.vcl; :
	kubectl apply -f varnish.yml; :
	kubectl create secret tls ssl --cert=server.crt --key=server.key; :
	kubectl create configmap nginx-conf --from-file=nginx.conf; :
	kubectl apply -f nginx-deploy.yml
.PHONY: run

delete:
	kubectl delete -f nginx-deploy.yml; :
	kubectl delete configmap nginx-conf; :
	kubectl delete secret ssl; :
	kubectl delete -f varnish.yml; :
	kubectl delete configmap varnish-config; :
	kubectl delete -f dictionary-deploy.yml; :
.PHONY: delete
