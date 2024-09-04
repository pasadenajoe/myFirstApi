package model

// --------------------------------------------------
// The following is my structure in Postgres.
// Hints:
// - In order to make available the fields of the structures
//   all the fields have to be in <upper case>
//   In the database, postgreSQL the fields are in <lower case>
// - Due return values of the CRUD function have to be in
//   json format, I tried to simplify the names.  The
//   abbreviatons are easy to understand for spanish native speakers.
// - Est_num is the table id, a sequence number
// - Est_ced is the equivalente to Id Card. This is a Unique Number
//   given by the country to each person.
// * * * * * * * * * *
// La siguiente es la estructura que se usa en Posgres
// Importante:
// - a fin de poder exponer los nombre de la misma, es necesario
//   que los campos estén en mayúsculas.
//   En la base de datos, los campos están en minúscula.
// - Debido a que los valores de la función CRUD se devuelven
//	 en formato json, traté de simplificar los nombres y darle su
//   equivalente.  Las abreviaturas usadas son fácilmente entendidas
//   por los hispano hablantes.
//   Est_num equivale a la llave principal de la tabla, un consecutivo.
//   Est_ced equivale al número de çédula, que es un número único dado
//   a cada persona en el pais.
// --------------------------------------------------
type Estudiante struct {
	Est_num          int64  `json:"Id"`
	Est_ced          string `json:"Cedula"`
	Est_p_nom        string `json:"P_Nom"`
	Est_s_nom        string `json:"S_Nom"`
	Est_p_apel       string `json:"P_Apel"`
	Est_s_apel       string `json:"S_Apel"`
	Est_fecha_nac    string `json:"FecNac"`
	Est_tipo_sangre  string `json:"TipoSangre"`
	Est_genero       string `json:"Genero"`
	Est_nacionalidad string `json:"Nacionalidad"`
}
