package repositories

import (
	"database/sql"
	"fmt"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (m *MysqlRepository) EmailExists(email string) (int, error) {
	// An album to hold data from the returned row.
	var existe int
	row := m.db.QueryRow("SELECT COUNT(1) existe FROM newsletter_emails WHERE email = ?", email)
	row.Scan(&existe)
	if existe > 0 {
		return 1, fmt.Errorf("this email already exists and cannot be duplicated")
	}
	// email does not exist
	return 0, nil
}

func (m *MysqlRepository) StoreData(email string) (int, error) {
	_, err := m.db.Exec("INSERT INTO newsletter_emails (email) VALUES (?)", email)
	if err != nil {
		return 0, err
	}
	/*id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil*/
	return 1, nil
}
