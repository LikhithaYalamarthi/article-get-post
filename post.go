package main
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
 )
 
 func main() {
 //Encode the data
	requestBody, err := json.Marshal(map[string]string{
	   "name":  "Likhitha",
	   "email": "Likhitha@example.com",
	})
	if err != nil {
		log.Fatalln(err)
	}
	
 //Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://inshorts.com/post", "application/json", bytes.NewBuffer(requestBody))
 //Handle Error
	if err != nil {
	   log.Fatalln(err)
	}
	defer resp.Body.Close()
 //Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	
	log.Println(string(body))
 }