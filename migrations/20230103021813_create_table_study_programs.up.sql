CREATE TABLE IF NOT EXISTS study_programs (
  id VARCHAR(200) PRIMARY KEY,
  departement_id VARCHAR(200) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (departement_id) REFERENCES departements(id)
)ENGINE=InnoDB;