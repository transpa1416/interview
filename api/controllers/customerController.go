package controllers

import (
	"net/http"
	"strconv"
	"time"

	conn "interview/api/connections"
	"interview/api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/**
 * Method: POST
 * endPoint: /NutriNET/Cliente
 */
func CreateCustomer(c *gin.Context) {
	var customerModel models.Customer
	if err := c.ShouldBindJSON(&customerModel); err != nil {
	  	c.JSON(http.StatusBadRequest, gin.H{
		  	"Cve_Error": -1,
			"Cve_Mensaje": err,
			"data": customerModel,
		})
	  return
	}

	hashPassword, err := hashPassword(customerModel.Contraseña)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Cve_Error": -1,
			"Cve_Mensaje": err,
			"data": customerModel,
	  })
	return
	}
  
	customer := models.Customer{
		Nombre_Usuario: customerModel.Nombre_Usuario,
		Contraseña: hashPassword,
		Nombre: customerModel.Nombre,
		Apellidos: customerModel.Apellidos,
		Correo_Electronico: customerModel.Correo_Electronico,
		Edad: customerModel.Edad,
		Estatura: customerModel.Estatura,
		Peso: customerModel.Peso,
		IMC: customerModel.IMC,
		GEB: customerModel.GEB,
		ETA: customerModel.ETA,
		Fecha_Creacion: time.Now(),
		Fecha_Actualizacion: time.Now()}
	
	if err := conn.DB.Create(&customer); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Cve_Error": -1,
			"Cve_Mensaje": err,
			"data": customer,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Cve_Error": 0,
			"Cve_Mensaje": nil,
			"data": customer,
		})
	}
}

/**
 * Method: GET
 * endPoint: /NutriNET/Cliente
*/
func FindCustomers(c *gin.Context) {
	var customerModel models.Customer
	idTxt := c.Query("clienteId")
	
	if idTxt != "" {
		id, errId := strconv.Atoi(idTxt)
		if errId != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Cve_Error": -1,
				"Cve_Mensaje": errId,
				"data": nil,
		  })
		  return
		}

		if err := conn.DB.Where("Cliente_ID = ?", id).First(&customerModel).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Cve_Error": -1,
				"Cve_Mensaje": errId,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Cve_Error": 0,
			"Cve_Mensaje": nil,
			"data": customerModel,
		})
	} else {
	
		var customersModel []models.Customer
		conn.DB.Find(&customersModel)
	
		c.JSON(http.StatusOK, gin.H{
			"Cve_Error": 0,
			"Cve_Mensaje": nil,
			"data": customersModel,
		})
	}
}

/**
 * Method: PUT
 * endPoint: /NutriNET/Cliente/?id
*/
func UpdateCustomer(c *gin.Context) {

	var customerModel models.Customer

	if err := conn.DB.Where("Cliente_ID = ?", c.Query("clienteId")).First(&customerModel).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{
		"Cve_Error": -1,
		"Cve_Mensaje": err,
		"data": nil,
	})
	  return
	}
  
	//Ingresamos valores
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	conn.DB.Model(&customerModel).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": customerModel})
}

//Función privada para la creación de la ocntraseña
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}