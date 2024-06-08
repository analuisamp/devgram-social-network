package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

//To get a key to hashKey and blockKey
/*
func init() {
	hashkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashkey)

	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockkey)
}
*/

func main() {
	config.Load()
	cookies.Config()
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
