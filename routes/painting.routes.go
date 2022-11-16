package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func IniPainting(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**CONFIG RESPONSE*/
	resp := make(map[string]string)
	resp["message"] = "iniciando Pating"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}



	/**REQUEST DE FRASES PARA EMPEZAR A PINTAR   ----> SE HA DE VERIFICAR QUE DIBUJOS HA HECHO Y HAY QUE PEDIRLE UNO NUEVO*/

	/**RESPONSE*/
	w.Write(jsonResp)
}

func SavePaint(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**POST PARA GUARDAR EL DIBUJO */

	r.ParseMultipartForm(0<<50)

	file,handler,err := r.FormFile("myFile");

	if err != nil {
		log.Println(err);
		return
	}

	defer file.Close()
	
	log.Print("Uploaded File: $+v\n", handler.Filename)

	tempFile, err := ioutil.TempFile("temp-images","upload-*.png")

	if err != nil {
		log.Println(err)
		return
	}

	defer tempFile.Close()


	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(fileBytes)
	tempFile.Write(fileBytes)


	/**DEVUELVE FRASE*/


	/**CONFIG RESPONSE*/
	resp := make(map[string]string)
	resp["message"] = "DIBUJO GUARDADO"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write([]byte(jsonResp))
	// fmt.Fprintf(w,"subiendo archivo")
}