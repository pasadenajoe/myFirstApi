package rutas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pasadenajoe/myFirstApi/datasources"
	"github.com/pasadenajoe/myFirstApi/model"
)

// --------------------------------------------------
// Routes associated to the root of the Api.
// * * * * * * * * * *
// Rutas asociadas a la raiz del Api.
// --------------------------------------------------
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api test, Hello!"))
}

// --------------------------------------------------
// EstGetByIdHandler, seek the "student" by the primary key.
// * * * * * * * * * *
// EstGetByIdHandler, devuelve el "estudiante" haciendo uso de
// la llave primaria
// --------------------------------------------------
func EstGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := (params["id"])
	id_number, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("El número del cliente presenta error en su etiqueta o en su valor."))
		return
	}

	registro, err := datasources.EstudianteById(r.Context(), int64(id_number))

	switch err {

	case nil:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(registro)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No se encontró registro coincidente"))
		return
	}

}

// --------------------------------------------------
// EstGetByCedHandler is used to return a unique record
// based on the id card number of the student.
// * * * * * * * * * *
// EstGetByCedHandler es el manejador que devuelve el
// estudiante basado en el número de cédula (único).
// --------------------------------------------------
func EstGetByCedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ced := (params["ced"])
	registro, err := datasources.EstudianteByCedula(r.Context(), ced)

	switch err {

	case nil:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(registro)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No se encontró registro coincidente"))
		return
	}

}

// --------------------------------------------------
// EstInsertHandler is the method used to insert a record
// in the database.
// * * * * * * * * * *
// EstInsertHandler es el método usado para la inserción del
// registro en la base de datos.
// --------------------------------------------------
func EstInsertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EstInsertHandler ----> <<entering>>")

	var e *model.Estudiante
	var err error

	fmt.Println("EstInsertHandler ----> ")

	// --------------------------------------------------
	// Important:  In this case, to the struct defined as
	// Estudiante (student), I set also de Json Tags, therefore
	// the json that is sending in the body have to match
	// those tags, and not the struct name.
	// In other words.  Est_ced is not to be send but
	// Cedula, which is the json tag.
	// Sometimes DisallowUnknownFields() is used to easily
	// see which fields are not matching exactly. An error will
	// show which field is causing the conflict.
	// * * * * * * * * * *
	// Importante: En este caso, la estructura definida como
	// Estudiante, se le asignaron etiquetas Json, de allí que
	// el json enviado en el cuerpo debe utilizar dichas etiquetas
	// y no los campos de la estructura.
	// En otras palabras, Est_ced no debe ser enviado sino
	// Cedula, es cual es la etiqueta json del campo.
	// A veces el método DisallowUnknowsFields() es usado
	// para identificar fácilmente qué campo está creando conflicto
	// y no cuadra con la estructura Json.
	// --------------------------------------------------
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&e)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	e, err = datasources.EstudianteInsert(r.Context(), e)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(e)
	w.WriteHeader(http.StatusOK)
}

// --------------------------------------------------
// Not Implemented.
// * * * * * * * * * *
// No implementado
// --------------------------------------------------
func EstDeleteHandler(w http.ResponseWriter, r *http.Request) {

}

// --------------------------------------------------
// EstUpdateHandler, is the method for updating the
// record of the student.
// * * * * * * * * * *
// EstUpdateHandler, es el método para actualizar el
// registro del estudiante.
// --------------------------------------------------
func EstUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var estudianteAct model.Estudiante
	var err error
	var b []byte

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&estudianteAct)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b = []byte(err.Error() + " Error decodificando")
		w.Write(b)
		return
	}

	e, registrosAfectados, err := datasources.EstudianteUpdate(r.Context(), &estudianteAct)

	if err != nil || registrosAfectados == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		b = []byte(err.Error() + " Después de EstudianteUpdate")
		w.Write(b)
		return
	}

	json.NewEncoder(w).Encode(e)
	w.WriteHeader(http.StatusOK)
}
