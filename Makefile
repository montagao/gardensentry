PROJECT_NAME=gardensentry


api:
	swagger generate server -t gen -A ${PROJECT_NAME} -f ./api.yml --exclude-main
install:
	go install ./cmd/${PROJECT_NAME}-server/
clean:
	rm -rf gen/*
