GARBAGE=assmap \
		debug \
		.assmap \

default: build run

build:
	go build

run:
	./assmap

clean:
	-rm -rf $(GARBAGE)