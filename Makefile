run:
	go run main.go

build:
	go build -o .build/my-gram

test:
	postman login --with-api-key PMAK-63bd040aa60c457c57c34da1-10769ff052a6d148bfed15954d29d2e7bc
	postman collection run 18453632-f1efe350-fefe-46c4-9fef-2ef88f3f7d45 -e 18453632-0c7a0a34-af9b-4675-98c8-2f3fd41ab4a5