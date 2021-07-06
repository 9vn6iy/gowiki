.PHONY : clean all

all : wiki server

wiki : wiki.go 
	go build wiki.go 

server: http-server.go
	go build http-server.go

clean:
	rm -rf *.txt http-server wiki 


