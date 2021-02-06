Krai is learning Go by follow along the linkedin learning's tutorial: 
https://www.linkedin.com/learning/go-essential-training/defining-function

go run filname
go build filename
go get github.com/pkg/errors
time go run filename
go mod init cfg
go test
go test -v
go test -run testname -v
go test -v -bench .(all)
go test -v -bench . -run TTT (only benchmark)
go test -v -bench . -run TTT -cpuprofile=prof.out
go tool pprof prof.out
list functionname

curl -d@request.json http://localhost:8080/math
http test package

install wrk