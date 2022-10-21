test:
	./etc/scripts/run-test.sh

# create linux 64 bit binary.
linux-bin:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/guardrails-linux-amd64 .

# create osx 64 bit binary.
darwin-bin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/guardrails-amd64-darwin .

# create windows 64 bit binary.
windows-bin:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-w -s" -o ./bin/guardrails.exe .

# install linux bin to ${GOPATH}/bin so it can be executed globally.
install-linux-bin:
	make linux-bin && chmod +x ./bin/guardrails-linux-amd64 && cp ./bin/guardrails-linux-amd64 ${GOPATH}/bin/guardrails
