run :
	@echo "----------------------------------Starting Program---------------------------------"
	@go run cmd/server/rest/main.go

build :
	@echo "----------------------------------Build Program------------------------------------"
	@go build -o ./main cmd/server/rest/main.go