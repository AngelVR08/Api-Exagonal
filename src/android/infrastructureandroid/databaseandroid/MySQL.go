package databaseandroid

import (
	"Angel/src/android/domainandroid/entities"
	"Angel/src/core"
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

func (mysql *MySQL) Save(android entities.Android) error {
	query := "INSERT INTO android (brand, model, data_storage) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, android.Brand, android.Model, android.Space)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Android, error) {
	query := "SELECT id_android, brand, model, data_storage FROM android"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var androids []entities.Android

	for rows.Next() {
		var android entities.Android
		if err := rows.Scan(&android.ID, &android.Brand, &android.Model, &android.Space); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		androids = append(androids, android)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return androids, nil
}

func (mysql *MySQL) Update(android entities.Android) error {
	query := "UPDATE android SET brand = ?, model = ?, data_storage = ? WHERE id_android = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, android.Brand, android.Model, android.Space, android.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el teléfono con ID %d: %v", android.ID, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún teléfono con ID %d para actualizar", android.ID)
	}

	log.Printf("[MySQL] - Teléfono con ID %d actualizado", android.ID)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM android WHERE id_android = ?"
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
