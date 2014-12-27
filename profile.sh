go build doc/profile_endgame_db.go 
go tool pprof --text profile_endgame_db cpu.profile 