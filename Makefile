default: build

test:
	go test ./...

build:
	go build -o bin/yatas-markdown

update:
	go get -u 
	go mod tidy

install: build
	mkdir -p ~/.yatas.d/plugins/github.com/StanGirard/yatas-markdown/local/
	mv ./bin/yatas-markdown ~/.yatas.d/plugins/github.com/StanGirard/yatas-markdown/local/

release: test
	standard-version
	git push --follow-tags origin main 