module main

// This tells go that module "examplelib" is at the path ../examplelib
// where ../examplelib is the path to the source.  If the path is
// relative (ie, starts with . or ..), the path is relative to this
// directory.
replace examplelib => ../examplelib

go 1.19

// To add this, run "go get examplelib"
require examplelib v0.0.0-00010101000000-000000000000 // indirect
