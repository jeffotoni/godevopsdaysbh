package main

import (
	"encoding/json"
	"log"
	"time"
)

type Login struct {
	Id       int64  `json:id`
	nome     string `json:nome`
	password string
}

var (
	a   = 100
	b   = "devops"
	c   = []byte("jeff")
	gus = "public"
	d   = "private"
)

func main() {

	sl := Login{
		Id:       2000,
		nome:     "fams@",
		password: "xi9w939393" + "xi393939",
	}
	bjson, err := json.Marshal(sl)
	if err != nil {
		log.Println("error json: ", err.Error())
		return
	}

	println(string(bjson))

	// bjson := []byte(`{"Id":2000,"Nome":"Isabela...","Password":"xi9w939393xi393939"}`)
	// jsonstr := string(bjson)
	// println(jsonstr)

	// var err error
	// var sl Login
	// err = json.Unmarshal([]byte(jsonstr), &sl)
	// if err != nil {
	// 	log.Println("error json: ", err.Error())
	// 	return
	// }

	return

	x := 1000

	for i := 0; i < 1000; i++ {
		dev1(i, x, sl.Id, sl)
	}

	println("meu primeiro programa go devopsdays...!")
	time.Sleep(time.Second * 5)
}

func dev1(i, y int, id int64, slx Login) {

	println("dev1:", i, " :: ", y,
		" :: ", gus, " id: ", id, slx.Id)
	time.Sleep(time.Second * 2)
}
