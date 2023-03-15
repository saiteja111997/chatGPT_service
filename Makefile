# DEV 
build-dev:
	docker build -t chatgptservice -f containers/images/Dockerfile .

# PROD
build-prod:
	docker build --platform linux/amd64 -t chatgptservice -f containers/images/Dockerfile .
