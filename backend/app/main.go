package main

import (
	"fmt"
	"os"
	"net/http"
	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
)

func main() {

	err := godotenv.Load(".env")
		
	if err != nil {
		fmt.Printf("Couldn't load : %v", err)
	} 
	
	key := os.Getenv("ENV")
	// fmt.Println(message)

	url := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=formatted_address,name,rating,opening_hours,geometry&key=" + key
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
  	fmt.Println(string(body))	

    // engine:= gin.Default()
    // engine.GET("/", func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{
    //         "message": "hello world hogehoge!!!",
    //     })
    // })
    // // engine.Run(":3000")	
}


// package main

// import (
//   "fmt"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Museum%20of%20Contemporary%20Art%20Australia&inputtype=textquery&fields=formatted_address,name,rating,opening_hours,geometry&key="
//   method := "GET"

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, nil)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }