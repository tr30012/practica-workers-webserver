package storage

const (
	SQLStringCreateTableJobless = `CREATE TABLE IF NOT EXISTS Jobless (
		id            INTEGER PRIMARY KEY AUTOINCREMENT
							  UNIQUE
							  NOT NULL,
		last_name     STRING,
		first_name    STRING,
		patronymic    STRING,
		age           INTEGER,
		passport      STRING,
		passport_date DATE,
		region        STRING,
		address       STRING,
		phone         STRING,
		picture       BLOB,
		study_place   STRING,
		study_address STRING,
		study_type    STRING,
		registrar     STRING,
		reg_date      DATE,
		payment       DOUBLE,
		comment       STRING
	);`

	SQLStringCreateTableJobgivers = `CREATE TABLE IF NOT EXISTS Jobgivers (
		id       INTEGER PRIMARY KEY AUTOINCREMENT
						 UNIQUE
						 NOT NULL,
		type     STRING,
		name     INTEGER,
		giver    STRING,
		place    STRING,
		mobile   STRING,
		district STRING,
		money    DOUBLE,
		more     STRING
	);`

	SQLStringCreateTableArchive = `CREATE TABLE IF NOT EXISTS Archive (
		id         INTEGER PRIMARY KEY AUTOINCREMENT
						   NOT NULL
						   UNIQUE,
		jobless_id INTEGER REFERENCES Jobless (id),
		job_id     INTEGER REFERENCES Jobgivers (id),
		date       DATE,
		archivist  STRING
	);`

	SQLStringInsertJobless = `INSERT INTO Jobless(
		last_name,
		first_name, 
		patronymic,
		age,
		passport,
		passport_date,
		region,
		address,
		phone,
		study_place,
		study_address,
		study_type,
		registrar,
		reg_date,
		payment,
		comment) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	SQLStringInsertArchive = `INSERT INTO Archive(
		jobless_id,
		job_id,
		date,
		archivist) VALUES (?, ?, ?, ?)`

	SQLStringInsertJobfiver = `INSERT INTO Jobgivers(
		jobless_id,
		job_id,
		date,
		archivist) VALUES (?, ?, ?, ?)`
)
