package datasources

import (
	"context"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/pasadenajoe/myFirstApi/model"
)

// --------------------------------------------------
// EstudianteUpdate, receives
// - ctx: context
// - estInput: the Student structure
// returns
// - Est: the Student structure after the update
// - filas: number of record affected
// - err: as error.
// * * * * * * * * * *
// EstudianteUpdate, recibe
// - ctx: contexto
// - estInput: estructura estudiante
// retorna
// - Est: La estructura después de la actualización
// - filas: número de registros afectados.
// - err: error.
// --------------------------------------------------
func EstudianteUpdate(ctx context.Context, estInput *model.Estudiante) (Est *model.Estudiante, filas int64, err error) {
	d := connectToDb()
	var estAct model.Estudiante
	var estudianteExistente *model.Estudiante
	var estudianteCed *model.Estudiante

	// --------------------------------------------------
	// The record is supposed to be sent with the id of the
	// student, to be seek.
	// * * * * * * * * * *
	// En principio el registro debería venir completo, con
	// número de cliente.
	// Dicho cliente se busca en la base de datos por ese
	// criterio.
	// --------------------------------------------------
	estudianteExistente, err = EstudianteById(ctx, estInput.Est_num)
	if err != nil {
		return estInput, 0, err
	}

	// --------------------------------------------------
	// In case that "cedula" (id number) might tried to be
	// updated, first has to be evaluated if the "new" id number
	// does not clash with an existing record.
	// * * * * * * * * * *
	// Pudiera producirse el error de que la cédula se esté
	// intentando actualizar, en tal caso, hay que verificar
	// que dicho campo no exista ya en la base de datos.
	// --------------------------------------------------
	if estInput.Est_ced != estudianteExistente.Est_ced {

		estudianteCed, err = EstudianteByCedula(ctx, estInput.Est_ced)
		if err != nil {
			return estInput, 0, err
		}

		if err == nil && estudianteCed.Est_ced != "" {
			err1 := fmt.Errorf("Se está intentando actualizar a un estudiante con un número de cédula que ya existe")
			return estInput, 0, err1
		}

	}

	estAct.Est_num = estudianteExistente.Est_num

	// --------------------------------------------------
	// Basically the validations are the same for all the fields.
	// In case the json field sent in the body has a
	// different value of the stored database value, the
	// new value is assigned to the update structure
	// This makes partial json sending possible, thus the json
	// only has the fields to be updated, and not
	// all the fields of the structure.
	// Ex. if I want to update de second name, I just send the
	// id and the second name, and just the second name will be
	// updated.
	// * * * * * * * * * *
	// Básicamente las validaciones son las mismas para todos los
	// campos. En caso de que el campo de la estructura enviada
	// no tenga el mismo valor que el almacenado en la base de datos,
	// el nuevo valor se asigna a la estructura de actualización.
	// Esto facilita el envío parcial del json de actualización
	// con los campos que solo se quieren actualizar, sin la necesidad de
	// enviarlos todos.
	// Ej. si quiero actualizar el segundo nombre, solo envío el número
	// del estudiante y el segundo nombre.
	// --------------------------------------------------
	if estInput.Est_ced != estudianteExistente.Est_ced &&
		estInput.Est_ced != "" {
		estAct.Est_ced = estInput.Est_ced
	} else {
		estAct.Est_ced = estudianteExistente.Est_ced
	}

	if estInput.Est_fecha_nac != estudianteExistente.Est_fecha_nac &&
		estInput.Est_fecha_nac != "" {
		estAct.Est_fecha_nac = estInput.Est_fecha_nac
	} else {
		estAct.Est_fecha_nac = estudianteExistente.Est_fecha_nac
	}

	if estInput.Est_genero != estudianteExistente.Est_genero &&
		estInput.Est_genero != "" {
		estAct.Est_genero = estInput.Est_genero
	} else {
		estAct.Est_genero = estudianteExistente.Est_genero
	}

	if estInput.Est_nacionalidad != estudianteExistente.Est_nacionalidad &&
		estInput.Est_nacionalidad != "" {
		estAct.Est_nacionalidad = estInput.Est_nacionalidad
	} else {
		estAct.Est_nacionalidad = estudianteExistente.Est_nacionalidad
	}

	if estInput.Est_p_apel != estudianteExistente.Est_p_apel &&
		estInput.Est_p_apel != "" {
		estAct.Est_p_apel = estInput.Est_p_apel
	} else {
		estAct.Est_p_apel = estudianteExistente.Est_p_apel
	}

	if estInput.Est_p_nom != estudianteExistente.Est_p_nom &&
		estInput.Est_p_nom != "" {
		estAct.Est_p_nom = estInput.Est_p_nom
	} else {
		estAct.Est_p_nom = estudianteExistente.Est_p_nom
	}

	if estInput.Est_s_apel != estudianteExistente.Est_s_apel &&
		estInput.Est_s_apel != "" {
		estAct.Est_s_apel = estInput.Est_s_apel
	} else {
		estAct.Est_s_apel = estudianteExistente.Est_s_apel
	}

	if estInput.Est_s_nom != estudianteExistente.Est_s_nom &&
		estInput.Est_s_nom != "" {
		estAct.Est_s_nom = estInput.Est_s_nom
	} else {
		estAct.Est_s_nom = estudianteExistente.Est_s_nom
	}

	if estInput.Est_tipo_sangre != estudianteExistente.Est_tipo_sangre &&
		estInput.Est_tipo_sangre != "" {
		estAct.Est_tipo_sangre = estInput.Est_tipo_sangre
	} else {
		estAct.Est_tipo_sangre = estudianteExistente.Est_tipo_sangre
	}

	qUpdate := "UPDATE estud SET est_ced = $1, est_p_nom = $2, "
	qUpdate = qUpdate + "est_s_nom = $3, est_p_apel = $4, "
	qUpdate = qUpdate + "est_s_apel = $5, est_fecha_nac = $6, "
	qUpdate = qUpdate + "est_tipo_sangre = $7, est_genero=$8, est_nacionalidad=$9 "
	qUpdate = qUpdate + "where est_num = " + strconv.FormatInt(estudianteExistente.Est_num, 10)

	Resultado, err1 := d.Exec(qUpdate,
		estAct.Est_ced, estAct.Est_p_nom,
		estAct.Est_s_nom, estAct.Est_p_apel,
		estAct.Est_s_apel, estAct.Est_fecha_nac,
		estAct.Est_tipo_sangre, estAct.Est_genero,
		estAct.Est_nacionalidad)

	if err1 != nil {
		log.Fatal(err1)

	}

	filasAfectadas, errAct := Resultado.RowsAffected()
	if errAct != nil {
		log.Fatal(errAct)
	}

	return &estAct, filasAfectadas, err1
}

// --------------------------------------------------
// EstudianteInsert is the function that insert record
// * * * * * * * * * *
// EstudianteInsert es la función que inserta el registro
// --------------------------------------------------
func EstudianteInsert(ctx context.Context, e *model.Estudiante) (Est *model.Estudiante, err error) {

	d := connectToDb()

	// --------------------------------------------------
	// Remarkable is the fact, that the sql string contains
	// the nextVal('') postgreSQL function.
	// Another important thing is the fact that variables in
	// postgreSQL use the form $<<number>>
	// * * * * * * * * * *
	// Es importante destacar el hecho de que la cadena
	// de inserción contiene el llamado de la función nextval('').
	// Otro aspecto importante es el hecho de que las variables
	// se definen $<<número>>
	// --------------------------------------------------
	qInsert := "INSERT INTO estud (est_num, est_ced,est_p_nom,est_s_nom,est_p_apel,est_s_apel,est_fecha_nac,"
	qInsert = qInsert + "est_tipo_sangre,est_genero,est_nacionalidad) "
	qInsert = qInsert + "VALUES (nextVal('estud_secuencia'), $1,$2,$3,$4,$5,$6,$7,$8,$9) "
	qInsert = qInsert + "RETURNING est_num	"

	var est_num int64
	est_num = 0

	err = d.QueryRow(qInsert, &e.Est_ced, &e.Est_p_nom, &e.Est_s_nom, &e.Est_p_apel, &e.Est_s_apel,
		&e.Est_fecha_nac, &e.Est_tipo_sangre, &e.Est_genero, e.Est_nacionalidad).Scan(&est_num)

	d.Close()
	e.Est_num = est_num

	return e, err
}

// --------------------------------------------------
// EstudianteByCedula return a struct of Estudiante (Student)
// using "cedula" as main parameter to seek the record.
// * * * * * * * * * *
// EstudianteByCedula retorna la estructura de Estudiante
// usando la cédula como parámetro para buscar el registro.
// --------------------------------------------------
func EstudianteByCedula(ctx context.Context, cedula string) (e *model.Estudiante, err error) {

	var estud model.Estudiante

	d := connectToDb()

	consulta := "SELECT * FROM estud WHERE est_ced=$1"

	row := d.QueryRow(consulta, cedula)

	err = row.Scan(
		&estud.Est_num,
		&estud.Est_ced,
		&estud.Est_p_nom,
		&estud.Est_s_nom,
		&estud.Est_p_apel,
		&estud.Est_s_apel,
		&estud.Est_fecha_nac,
		&estud.Est_tipo_sangre,
		&estud.Est_genero,
		&estud.Est_nacionalidad,
	)

	defer d.Close()
	return &estud, err
}

// --------------------------------------------------
// EstudianteById return a struct of Estudiante (Student)
// using "id" (Est_num) as main parameter to seek the record.
// * * * * * * * * * *
// EstudianteById retorna la estructura de Estudiante
// usando el Est_num como parámetro para buscar el registro.
// --------------------------------------------------
func EstudianteById(ctx context.Context, id int64) (e *model.Estudiante, err error) {

	var estud model.Estudiante

	d := connectToDb()

	consulta := "SELECT * FROM estud WHERE est_num=$1"

	row := d.QueryRow(consulta, id)

	err = row.Scan(
		&estud.Est_num,
		&estud.Est_ced,
		&estud.Est_p_nom,
		&estud.Est_s_nom,
		&estud.Est_p_apel,
		&estud.Est_s_apel,
		&estud.Est_fecha_nac,
		&estud.Est_tipo_sangre,
		&estud.Est_genero,
		&estud.Est_nacionalidad,
	)

	defer d.Close()
	return &estud, err
}
