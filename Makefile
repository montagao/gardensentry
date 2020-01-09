PROJECT_NAME=gardensentry


api:
	swagger generate server -A ${PROJECT_NAME} -f ./api.yml
install:
	go install ./cmd/${PROJECT_NAME}-server/
