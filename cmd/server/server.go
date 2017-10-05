package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roscopecoltran/krakend-admin/db"
)

func startAdminUI() {

	mux := http.NewServeMux()
	if !db.Store.Admin.Disabled {

		db.Store.Admin.UI.SetSiteName("SniperKit Gateway - Administration Panel")
		// Add Dashboard
		db.Store.Admin.UI.MountTo("/admin", mux)
	}

	r := gin.Default()
	if !db.Store.Admin.Disabled {
		r.Any("/admin/*w", gin.WrapH(mux))
	}
	r.Run()

}
