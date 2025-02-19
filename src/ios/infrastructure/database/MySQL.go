package database

import (
	"Angel/src/core"
	"Angel/src/ios/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar un nuevo registro de Ios
func (mysql *MySQL) Save(ios entities.Ios) error {
	query := "INSERT INTO ios (data_storage, model, batery) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, ios.Space, ios.Model, ios.Batery)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

// Obtener todos los registros de Ios
func (mysql *MySQL) GetAll() ([]entities.Ios, error) {
	query := "SELECT id_ios, data_storage, model, batery FROM ios"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var ioss []entities.Ios

	for rows.Next() {
		var ios entities.Ios
		if err := rows.Scan(&ios.ID, &ios.Space, &ios.Model, &ios.Batery); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		ioss = append(ioss, ios)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return ioss, nil
}

// Obtener un registro de Ios por su ID
func (mysql *MySQL) GetById(id int) (entities.Ios, error) {
	query := "SELECT id_ios, data_storage, model, batery FROM ios WHERE id_ios = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return entities.Ios{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var ios entities.Ios
	if rows.Next() {
		if err := rows.Scan(&ios.ID, &ios.Space, &ios.Model, &ios.Batery); err != nil {
			return entities.Ios{}, fmt.Errorf("error al escanear el teléfono: %v", err)
		}
		return ios, nil
	}

	return entities.Ios{}, fmt.Errorf("no se encontró el teléfono con ID %d", id)
}

// Actualizar un registro de Ios
func (mysql *MySQL) Update(ios entities.Ios) error {
	query := "UPDATE ios SET model = ?, batery = ?, data_storage = ? WHERE id_ios = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, ios.Model, ios.Batery, ios.Space, ios.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el teléfono con ID %d: %v", ios.ID, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún teléfono con ID %d para actualizar", ios.ID)
	}

	log.Printf("[MySQL] - Teléfono con ID %d actualizado", ios.ID)
	return nil
}

// Eliminar un registro de Ios
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM ios WHERE id_ios = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el teléfono con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún teléfono con ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Teléfono con ID %d eliminado", id)
	return nil
}
