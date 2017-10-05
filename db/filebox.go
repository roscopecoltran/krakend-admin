package db

import (
	"github.com/qor/filebox"
	"github.com/qor/roles"
	// "github.com/roscopecoltran/krakend-admin/config"
	// "github.com/roscopecoltran/krakend-admin/config/auth"
)

var Filebox *filebox.Filebox

func init() {
	Filebox = filebox.New("./shared/public/downloads")
	// Filebox.SetAuth(auth.AdminAuth{})
	dir := Filebox.AccessDir("/")
	dir.SetPermission(roles.Allow(roles.Read, "admin"))
}
