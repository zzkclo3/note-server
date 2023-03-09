package main

import (
    "github.com/gin-gonic/gin"
    "note-api/routes"
    "note-api/config"
)

func main()  {
    router := gin.Default()
   
    config.ConnectDB()
    
    routes.UserRoute(router)
    router.Run("localhost:6000")
}