// hello world for go
package main

import (
	databases "modules/gin/Databases"
	"modules/gin/Router"
)

func main() {
	databases.ConnectDB()
	Router.InitRouter()
}
