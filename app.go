package main

import "asset-manager.com/asset-manager/src/server"

func init() {

}

func main() {
	// load config from env
	//environment := os.Getenv("ENV")
	// db.Init
	server.Init()
}