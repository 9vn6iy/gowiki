.PHONY : clean all

all : wiki 

wiki : wiki.go 
	go build wiki.go 

clean:
	rm -rf wiki 


