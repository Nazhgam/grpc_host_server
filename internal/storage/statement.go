package storage

const (
	createHostTable = `
		CREATE TABLE IF NOT EXISTS questions(
			id BIGSERIAL PRIMARY KEY,
			question TEXT NOT NULL,
			first_wrong_answer TEXT NOT NULL,
			second_wrong_answer TEXT NOT NULL,
			third_wrong_answer TEXT NOT NULL,
			correct_answer TEXT NOT NULL
		);
	`
	insertHost = `
			INSERT INTO questions (question, first_wrong_answer, second_wrong_answer, third_wrong_answer, correct_answer)
			VALUES (?, ?, ?, ?, ?) RETURNING id;
	`
)

func (s *storage) createTables() error {
	_, err := s.db.Exec(createHostTable)
	return err
}
