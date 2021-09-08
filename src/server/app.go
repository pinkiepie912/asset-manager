package server

func Init() {
	// Load Config
	router := Router()
	router.Run(":8080")
}