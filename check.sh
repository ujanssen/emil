find . -type f -name "*.go" -exec golint {} \;
go test

go run doc/build_endgame_db.go
