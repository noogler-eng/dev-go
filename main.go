package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

type Todo struct {
	UserID    int    //`json:"userId"`
	ID        int    //`json:"id"`
	Title     string //`json:"title"`
	Completed bool   //`json:"completed"`
}

// fake api's
// https://jsonplaceholder.typicode.com/todos/1
func main() {

	fmt.Println("learning http's in go")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}

	// always closing the connection of res
	// res managment
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	fmt.Println("Response url:", res.Request.URL)
	fmt.Printf("Response type %T\n:", res) // *http.Response

	data, err := ioutil.ReadAll(res.Body)
	fmt.Printf("data type %T\n:", data) // bytes[]/[]uint8
	if err != nil {
		panic(err)
	}
	dataInString := string(data)
	fmt.Println("data", dataInString)
	var todo Todo
	json.Unmarshal(data, &todo)
	fmt.Printf("Data: %+v\n", todo)
	fmt.Printf("%v\n", todo.Title)

	// urls
	s_url := "http://user:pass@host.com:5432/path?k=v#f"
	// convert it into url from string
	u, err := url.Parse(s_url)
	if err != nil {
		panic(err)
	}
	fmt.Println("url parsed: ", u)
	fmt.Printf("%T\n", u) // *url.URL
	// url communication protocol
	fmt.Println(u.Scheme)          // http
	fmt.Println(u.User)            // user:pass
	fmt.Println(u.User.Username()) // user

	p, _ := u.User.Password()
	fmt.Println(p) // pass

	fmt.Println(u.Host)                        // host.com:5432
	host, port, _ := net.SplitHostPort(u.Host) // host.com  // 5432
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)     // /path
	fmt.Println(u.Fragment) // f

	fmt.Println(u.RawQuery)            // k=v
	m, _ := url.ParseQuery(u.RawQuery) //map[k:[v]]
	fmt.Println(m)
	fmt.Println(m["k"][0])

	// changing the url component
	u.Host = "football.com:1234"
	u.Path = "/newpath/sharad"
	u.RawQuery = "love>war"
	u.Fragment = "pyaar"
	fmt.Println("changed string", u)

	// create an new url
	newUrl := u.String();
	fmt.Printf("conversion of url to string %T\n", newUrl) 
	fmt.Println("new url :: ", newUrl)

}
