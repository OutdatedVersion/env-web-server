clean:
	rm -rf build

build: clean
	go build -o build/server