module main

replace examplelib => ../examplelib

go 1.19

// To add this, run "go get examplelib"
require examplelib v0.0.0-00010101000000-000000000000 // indirect
