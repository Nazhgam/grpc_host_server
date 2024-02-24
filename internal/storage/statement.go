package storage

const (
	createQuizTable = `
		CREATE TABLE IF NOT EXISTS quiz(
			id BIGSERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL
		);
	`
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
	insertQuiz = `
			INSERT INTO quiz (title, description)
			VALUES ($1, $2) RETURNING id;
	`
	insertHost = `
			INSERT INTO questions (question, first_wrong_answer, second_wrong_answer, third_wrong_answer, correct_answer)
			VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`
)

func (s *storage) createTables() error {
	_, err := s.db.Exec(createQuizTable)
	return err
}
