CREATE TABLE IF NOT EXISTS departements (
  id VARCHAR(200) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  code VARCHAR(5) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
)ENGINE=InnoDB;