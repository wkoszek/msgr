package main


// Context holds a current application state
type Context struct {
	From		 string
	To		 string
	Config		*Config
	ArgWhere	*string
	ArgCode		*bool
	ArgProfileName	*string
}
