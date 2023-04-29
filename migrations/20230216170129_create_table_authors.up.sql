CREATE TABLE IF NOT EXISTS authors (
  id VARCHAR(200) PRIMARY KEY,
  repository_id VARCHAR(200) NOT NULL,
  pemustaka_id VARCHAR(200) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (repository_id) REFERENCES repositories(id) ON DELETE CASCADE,
  CONSTRAINT FOREIGN KEY (pemustaka_id) REFERENCES pemustakas(id)
)ENGINE=InnoDB;