package router

import (
	"google-translator/constant"
	"google-translator/controller"
	"net/http"
)

// This routes represent user
var translateRoutes = Routes{
	Route{"Translate to language", http.MethodPost, constant.TranslatePath, controller.Translate},
}
