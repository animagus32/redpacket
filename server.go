package main

import "github.com/go-martini/martini"
import "redpacket/controllers"

func main(){
	m := martini.Classic()

	m.Get("/",func() string{
		return "Hello World!"
	});

	m.Group("/api/v1",func(r martini.Router){
		r.Post("/redpacket/dispatch",controllers.Dispatch)
		r.Get("/redpacket/grab",controllers.Grab)
		r.Get("/redpacket/list",controllers.Index)

		r.Get("/user/balance",controllers.UserBalance)

	});


	m.Run()

}
