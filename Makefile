build-binary:	
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o service1/.builds/binary service1/main.go
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o service2/.builds/binary service2/main.go
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o service3/.builds/binary service3/main.go
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o service4/.builds/binary service4/main.go
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o service5/.builds/binary service5/main.go
