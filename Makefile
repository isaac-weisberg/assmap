GARBAGE=assmap \
	debug \

default: build run

build:
	go build

run:
	./assmap

clean:
	-rm -rf $(GARBAGE)