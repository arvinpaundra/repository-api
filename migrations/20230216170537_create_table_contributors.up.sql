CREATE TABLE IF NOT EXISTS contributors (
  id VARCHAR(200) PRIMARY KEY,
  repository_id VARCHAR(200) NOT NULL,
  pemustaka_id VARCHAR(200) NOT NULL,
  contributed_as ENUM('Pembimbing 1', 'Pembimbing 2', 'Penguji 1', 'Penguji 2', 'Pembimbing Magang') NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (repository_id) REFERENCES repositories(id) ON DELETE CASCADE,
  CONSTRAINT FOREIGN KEY (pemustaka_id) REFERENCES pemustakas(id)
)ENGINE=InnoDB;