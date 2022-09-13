current_dir = $(shell pwd)

.PHONY: generator-ca-root
generator-ca-root:
	@openssl genrsa -out cert/ca.key 2048
	@openssl req -new -x509 -key cert/ca.key -out cert/ca.crt -days 365 -subj "/C=CN/ST=hubei/L=wuhan/O=lbh/OU=demo/CN=micro"

.PHONY: generator-ca-etcd
generator-ca-etcd:
	@echo "subjectAltName=DNS:etcd.micro.com,IP:127.0.0.1" > cert/etcd.conf
	@openssl genrsa -out cert/etcd.key 2048
	@openssl req -new -key cert/etcd.key -out cert/etcd.csr -subj "/C=CN/ST=hubei/L=wuhan/O=lbh/OU=demo/CN=etcd"
	@openssl x509 -req -days 365 -sha256 -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -extfile cert/etcd.conf -in cert/etcd.csr -out cert/etcd.crt

.PHONY: generator-ca-transport
generator-ca-transport:
	@echo "subjectAltName=DNS:*.server.com,IP:10.3.73.160" > cert/transport.conf
	@openssl genrsa -out cert/transport.key 2048
	@openssl req -new -key cert/transport.key -out cert/transport.csr -subj "/C=CN/ST=hubei/L=wuhan/O=lbh/OU=demo/CN=transport"
	@openssl x509 -req -days 365 -sha256 -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -extfile cert/transport.conf -in cert/transport.csr -out cert/transport.crt

.PHONY: build-tls-etcd
build-tls-etcd:
	@docker run -d --name etcd-tls \
		 -p 2379:2379 \
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
