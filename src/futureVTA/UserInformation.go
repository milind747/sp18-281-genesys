package main

import (
	"net/http"
	"github.com/unrolled/render"
	)


type UserDetails struct {
	uid		string		
	username	string		
	parentid	string		
	phone		string		
	email		string		
}

