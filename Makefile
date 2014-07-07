all:
	GOOS=linux go build
clean:
	rm sample_app
push:
	git commit -a -m 'this is silly'; git push origin master
