package repositories

import "newsletter-backend/types"

type PostgresqlRepository struct {
}

func NewPostgresqlRepository() *PostgresqlRepository {
	return &PostgresqlRepository{}
}

func (m *PostgresqlRepository) GetEmails() ([]types.Newsletter, error) {
	return nil, nil
}
func (m *PostgresqlRepository) EmailExists(email string) (int, error) {
	return 1, nil
}

func (m *PostgresqlRepository) StoreData(email string) (int, error) {
	return 1, nil
}
