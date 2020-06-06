package main


// Context holds a current application state
type Context struct {
	Config		*Config

	ArgFrom		*string
	ArgTo		*string
	ArgSubject	*string
	ArgWhere	*string
	ArgProfileName	*string
	ArgCode		*bool
	ArgHTML		*bool
	ArgDry		*bool
}
