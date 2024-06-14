package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Create
// Delete
// Read
// Update

// 1. creation of server
// 2. database connection to server
// 3. endpoints on server

// this should be match from the incomming response
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func postReq() {
	todo := Todo{
		UserId:    1,
		Id:        123,
		Title:     "love triangle",
		Completed: true,
	}

	// convert todo into json
	dataInJson, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	dataInString := string(dataInJson)
	stringReader := strings.NewReader(dataInString)
	res, err_1 := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", stringReader)
	if err_1 != nil {
		panic(err_1)
	}
	defer res.Body.Close()
	fmt.Println("^^^^^")
	data, err_2 := ioutil.ReadAll(res.Body)
	if err_2 != nil {
		panic(err_2)
	}
	fmt.Println(res.StatusCode, res.Status)
	fmt.Println(string(data))
	fmt.Println("^^^^^")
}

func putReq() {

	// making of an own request
	// 1. making an NewReqest by putting method
	// 2. setting up the headers like content-type
	// 3. making of an client
	// 4. send req by client
	// 4. close the connection

	todo := Todo{
		UserId:    23,
		Title:     "love triangle",
		Completed: true,
	}

	dataInJson, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	dataInString := string(dataInJson)
	dataReader := strings.NewReader(dataInString)

	// updating todo with id = 1
	req, err_1 := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", dataReader)
	if err_1 != nil {
		panic(err_1)
	}
	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, err_2 := client.Do(req)
	if err_2 != nil {
		panic(err_2)
	}
	defer res.Body.Close()

	data, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("******")
	// fmt.Println(res)
	fmt.Println(res.Status)
	fmt.Println(string(data))
	fmt.Println("******")
}

func main() {
	fmt.Println("learning crud...")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// var todo_1 Todo
	// err_0 := json.NewDecoder(res.Body).Decode(&todo_1)
	// fmt.Println("--------")
	// if err_0 != nil {
	// 	fmt.Println(err_0)
	// }
	// fmt.Println(todo_1)
	// fmt.Println(todo_1.title, todo_1.completed)
	// fmt.Println("--------")

	// fmt.Println("data incomming", res)
	// returns data in bytes
	data, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		panic(err1)
	}

	if strconv.Itoa(http.StatusOK) != strings.Split(res.Status, " ")[0] {
		fmt.Println("error in server side")
		return
	}

	// fmt.Println("data in bytes: ", data)
	fmt.Println("data in string: ", string(data))

	// decoding into struct
	var todo Todo
	err3 := json.Unmarshal(data, &todo)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(todo, todo.UserId, todo.Title, todo.Completed)

	postReq()
	putReq()
}
