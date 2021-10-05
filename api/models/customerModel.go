package models

import (
	"time"
)

type Customer struct {
  Cliente_ID     		int			`json:"clienteId" gorm:"primary_key;auto_increment;not_null"`
  Nombre_Usuario  		string 		`json:"nombreUsuario" gorm:"unique"`
  Contraseña 			string 		`json:"contraseña"`
  Nombre 				string 		`json:"nombre"`
  Apellidos 			string 		`json:"apellidos"`
  Correo_Electronico 	string 		`json:"correoElectronico" gorm:"unique"`
  Edad 					int  		`json:"edad"`
  Estatura 				float64 	`json:"estatura"`
  Peso 					float64 	`json:"peso"`
  IMC 					float64 	`json:"IMC"`
  GEB 					float64 	`json:"GEB"`
  ETA 					float64 	`json:"EMC"`
  Fecha_Creacion 		time.Time 	`json:"fechaCreacion"`
  Fecha_Actualizacion 	time.Time 	`json:"fechaActualizacion"`
}