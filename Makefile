GARBAGE=assmap \
		debug \
		.assmap \

default: build run

build:
	go build

run:
	./assmap

install:
	go install

clean:
	-rm -rf $(GARBAGE)