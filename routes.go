package main

//
//import (
//	"net/http"
//
//	"github.com/gorilla/mux"
//)
//
//type route struct {
//	Name        string
//	Method      string
//	Pattern     string
//	HandlerFunc http.HandlerFunc
//}
//
//type routes []route
//
//var operationRoutes = map[string]routes{
//	"data": {
//		route{
//			"GetdataFromListURL",
//			"POST",
//			"list",
//			GetData,
//		},
//	},
//}
//
//func newRouter() *mux.Router {
//	for node, routes := range operationRoutes {
//		for _, route := range routes {
//			var handler http.Handler
//
//			handler = route.HandlerFunc
//		}
//	}
//}
