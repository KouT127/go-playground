package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"os"
"time"
)

type Order struct {
	Id       int `json:"id"`
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}
type User struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Gender string   `json:"gender"`
	Orders []*Order `json:"orders"`
}

func main() {
	exportJson("data")
	importJson("data")
}

func importJson(dir string) []*User {
	var (
		err   error
		user  *User
		users []*User
	)
	fmt.Println("start file reading")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		d := fmt.Sprintf("%s/%s", dir, file.Name())
		f, err := os.Open(d)
		if err != nil {
			fmt.Printf("No such file or directory %s", file.Name())
			return nil
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err = json.Unmarshal(b, &user); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	fmt.Println("finish file read")
	return users
}

func exportJson(dir string) {
	var (
		bytes  []byte
		err    error
		orders []*Order
	)
	orders = make([]*Order, 0)
	for i := 1; i <= 10; i++ {
		order := &Order{
			Id:       i,
			Price:    i * 100,
			Quantity: i * 10,
		}
		orders = append(orders, order)
	}
	user := User{
		Id:     "1",
		Name:   "Name",
		Gender: "1",
		Orders: orders,
	}
	if bytes, err = json.Marshal(user); err != nil {
		panic(err)
	}
	file := fmt.Sprintf("%s/sample%s.json", dir, time.Now().Format("200601021504"))
	if err = ioutil.WriteFile(file, bytes, 0644); err != nil {
		panic(err)
	}
}
