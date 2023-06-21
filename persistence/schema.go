package persistence

const Schema = `
	CREATE TABLE IF NOT EXISTS student (
		id SERIAL PRIMARY KEY,
		first_name varchar,
		last_name varchar,
		email varchar,
		enrollment_year int
	);
`
