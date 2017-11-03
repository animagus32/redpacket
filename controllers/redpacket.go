package controllers


import (
	"net/http"
)

func Dispatch(res http.ResponseWriter,req *http.Request){
	res.Write([]byte("Dispatch"));
}

func Grab(res http.ResponseWriter,req *http.Request)  {
	res.Write([]byte("Grab"));

}

func UserBalance(res http.ResponseWriter,req *http.Request)  {
	res.Write([]byte("UserBalance"));

}

func Index(res http.ResponseWriter,req *http.Request)  {
	res.Write([]byte("Index"));

}