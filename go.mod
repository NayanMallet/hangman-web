module hangman-web

go 1.19

replace hangman-classic => ../hangman-classic

require hangman-classic v0.0.0-00010101000000-000000000000

require (
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.1.0 // indirect
)
