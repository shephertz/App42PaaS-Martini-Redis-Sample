package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/hoisie/redis"
	"github.com/martini-contrib/render"
	"net/http"
)

var (
	client redis.Client
)

type Profile struct {
	Username string
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() {
  client.Addr = "192.168.3.241:37474"
  client.Db = 0
  client.Password = "ag1mes2c27b39gxnywz7fasbyznrwlf1"
  
  //client.Addr = "127.0.0.1:6379"
  //client.Db = 0
}

func main() {
	m := martini.Classic()
	SetupDB()

	// reads "templates" directory by default
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Post("/users", func(ren render.Render, r *http.Request) {

		fmt.Println(r.FormValue("name"))
		
		user := r.FormValue("name")
		err := client.Rpush("username", []byte(user))

	  fmt.Println("Redis Error:", err)
		
		username, _ := client.Lrange("username", 0, 10000)
	  fmt.Println("User--->", username)

	  profiles := []Profile{}
	  for i, v := range username {
		  println("Name", i,":",string(v))

		  profile := Profile{}
		  profile.Username = string(v)
		  profiles = append(profiles, profile)
	  }

		fmt.Println("Profile details", profiles)
		ren.HTML(200, "users", profiles)
	})

	m.Run()

}
