find . -type f -name "*.go" -exec golint {} \;
go test

go run doc/setup_a_chessboard.go 

