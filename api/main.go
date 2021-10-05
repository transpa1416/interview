package main

import (
	conn "interview/api/connections"
	"interview/api/migrations"
	"interview/api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  conn.GetMySQLConn()
  //Se utiliza para crear la tabla
  migrations.MigrateCustomer()
  
  routers.LoadCustomerPaths(r)

  r.Run()
}