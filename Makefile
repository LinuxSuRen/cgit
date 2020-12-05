build:
	CGO_ENABLE=0 go build -ldflags "-w -s" -o bin/cgit

copy: build
	sudo cp bin/cgit /usr/local/bin/cgit
