build:
	go build -o main

run: 
	./main

clean:
	rm ./main

dev:
	reflex -s make


