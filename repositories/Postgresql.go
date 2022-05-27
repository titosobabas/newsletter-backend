package repositories

type PostgresqlRepository struct {
}

func NewPostgresqlRepository() *PostgresqlRepository {
	return &PostgresqlRepository{}
}

func (m *PostgresqlRepository) EmailExists(email string) (int, error) {
	return 1, nil
}

func (m *PostgresqlRepository) StoreData(email string) (int, error) {
	return 1, nil
}
