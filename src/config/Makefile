.PHONY: generator-pem
generator-pem:
	@openssl genrsa -out private_key.pem 2048
	@openssl rsa -in private_key.pem -pubout -out public_key.pem