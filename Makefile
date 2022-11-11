current_dir = $(shell pwd)

.PHONY: generator-ca-root
generator-ca-root:
	@openssl genrsa -out cert/ca.key 2048
	@openssl req -new -x509 -key cert/ca.key -out cert/ca.crt -days 365 -subj "/C=CN/ST=hubei/L=wuhan/O=lbh/OU=demo/CN=micro"

.PHONY: generator-ca
generator-ca:
ifneq (cert/ca.key_cert/ca.crt,$(wildcard cert/ca.key)_$(wildcard cert/ca.crt))
	$(error "ca.key or ca.crt is not exist. please execution 'make generator-ca-root' first")
endif
	@echo "subjectAltName=DNS:$(option).micro.com,IP:127.0.0.1,IP:10.3.73.160,IP:0.0.0.0" > cert/$(option).conf
	@openssl genrsa -out cert/$(option).key 2048
	@openssl req -new -key cert/$(option).key -out cert/$(option).csr -subj "/C=CN/ST=hubei/L=wuhan/O=lbh/OU=demo/CN=$(option)"
	@openssl x509 -req -days 365 -sha256 -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -extfile cert/$(option).conf -in cert/$(option).csr -out cert/$(option).crt

.PHONY: build-tls-transport
build-tls-transport: option=transport
build-tls-transport: generator-ca

.PHONY: build-tls-broker
build-tls-broker: option=broker
build-tls-broker: generator-ca

.PHONY: build-tls-etcd
build-tls-etcd: option=etcd
build-tls-etcd: generator-ca
	@docker run -d --name etcd-tls \
		 -p 12379:2379 \
		 --mount type=bind,source=$(current_dir)/cert/etcd.crt,destination=/etcd/cert/server.crt \
		 --mount type=bind,source=$(current_dir)/cert/etcd.key,destination=/etcd/cert/server.key \
		 quay.io/coreos/etcd:v3.5.0 \
		 /usr/local/bin/etcd \
			--name s1 \
		  	--data-dir /etcd-data \
			--listen-client-urls https://0.0.0.0:2379 \
			--advertise-client-urls https://0.0.0.0:2379 \
            --listen-peer-urls http://0.0.0.0:2380 \
            --initial-advertise-peer-urls http://0.0.0.0:2380 \
            --initial-cluster s1=http://0.0.0.0:2380 \
            --initial-cluster-token tkn \
            --initial-cluster-state new \
            --log-level info \
            --logger zap \
            --log-outputs stderr \
            --cert-file=/etcd/cert/server.crt \
            --key-file=/etcd/cert/server.key \

.PHONY: build-redis
build-redis:
	@docker run -d --name redis -p 6379:6379 redis:7.0.4

.PHONY: build-mongodb
build-mongodb:
	@docker run -d --name mongodb -p 27017:27017 \
 		-e MONGO_INITDB_ROOT_USERNAME=admin \
		-e MONGO_INITDB_ROOT_PASSWORD=123456 \
		mongo:6.0