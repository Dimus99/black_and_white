package main

var imgSuffixes = [3]string{
	".jpeg", ".jpg", ".png",
}
var workersCount = 3

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "log_entry"
)
