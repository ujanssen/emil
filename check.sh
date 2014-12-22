find . -type f -name "*.go" -exec golint {} \;
go test

go run doc/print_a_chessboard_for_doc.go 
go run doc/setup_a_chessboard.go 

