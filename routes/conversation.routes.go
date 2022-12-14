package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "strconv"

	"github.com/eloy411/project-M12-BACK/config"
	"github.com/eloy411/project-M12-BACK/models"
)

var testCalcul int


func GetConversation(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/** RECIBIMOS LOS DATOS DEL USUARIO*/

	var user models.User
	
	
	err := json.NewDecoder(r.Body).Decode(&user)

	if(user.Edad >=3 && user.Edad <= 5){
		testCalcul = 100
	}else if(user.Edad >=6 && user.Edad <= 8){
		testCalcul = 200
	}else if(user.Edad >=9 && user.Edad <= 12){
		testCalcul = 300
	}

	testCalcul = testCalcul + (user.Gravedad * 10) + user.Numtest
	fmt.Println(testCalcul)
	/** REQUEST PARA PEDIR LAS CONVERSACIONES SEGUN DIA, CONFIANZA, NIVEL DE ENFERMEDAD, NUM CONVERSACIONES*/
	var result []models.Preguntas
	config.DB.Table("preguntas").Select("*").Where("Id_Test = ?",testCalcul+1).Scan(&result)

	jsonResp, err := json.Marshal(&result)   

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	/**RESPONSE*/

	// fmt.Println(jsonResp)
	w.Write(jsonResp)

	
}


func RegisterResponses(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/** REGISTRAR LAS RESPUESTAS */

	var respuestas models.ResponsesTestsDaily
	var user models.User


	err := json.NewDecoder(r.Body).Decode(&respuestas)
	fmt.Println(respuestas)
	fmt.Println(user)

 	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	config.DB.Create(&respuestas)
	config.DB.Model(&user).Where("Id_User = ?",respuestas.IdUser).Update("NumTest",respuestas.Numtest)

  
	/**RESPONDER AL CLIENTE*/
 
	
	resp := make(map[string]string)
	resp["message"] = "almacenando respuestas"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}


	w.Write(jsonResp)
} 