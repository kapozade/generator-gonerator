BINARY_NAME="<%=appname%>"
PACKAGE="<%=repository%>

all: build run

build: ## Compiles sources and builds executable
	@go mod init <%=appname%> || (echo "Holy \xF0\x9F\x92\xA9 ! Something bad happened during the download.." && exit 1)
	@go mod tidy || (echo "Holy \xF0\x9F\x92\xA9 ! Something bad happened during the download.." && exit 1)
	@echo "Compiled, cheers \xF0\x9F\x8D\xBA!"

run: ## Run baby run, faster than my bullet
	@go run main.go