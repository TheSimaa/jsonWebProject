package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func getUsers()  {

	response, _ := http.Get("http://localhost:8000/api/users")

		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)

		fmt.Printf("%s\n", string(contents))


}

func main()  {

	getUsers()

}

