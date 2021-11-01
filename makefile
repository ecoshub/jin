NAME="jin"

all: test build

build:
	@echo "building..."
	@go mod tidy
	@go build -o "$(NAME)" .
	@echo "done."
	@echo "executable created. "$(NAME)""

test:
	@go test -v .

help:
	@echo '-------------'
	@echo '|  welcome  |'
	@echo '-------------'
	@echo ''
	@echo 'example usage:'
	@echo '    $$ make build'
	@echo ''
	@echo 'options:'
	@echo '    - help:  print this help dialog.'
	@echo '    - build: build the package. create an executable in working dir.'
	@echo '    - test:  execute tests in verbose mode.'


.DEFAULT_GOAL := help
