CREATE TABLE IF NOT EXISTS departements (
  id VARCHAR(200) PRIMARY KEY,
  study_program_id VARCHAR(200) NOT NULL,
  name VARCHAR(255) NOT NULL,
  code VARCHAR(5) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (study_program_id) REFERENCES study_programs(id)
)ENGINE=InnoDB;