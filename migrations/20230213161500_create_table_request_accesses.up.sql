CREATE TABLE IF NOT EXISTS request_accesses (
  id VARCHAR(200) PRIMARY KEY,
  pemustaka_id VARCHAR(200) NOT NULL,
  support_evidence TEXT NOT NULL,
  status ENUM('pending', 'accepted', 'denied') DEFAULT 'pending',
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (pemustaka_id) REFERENCES pemustakas(id)
)ENGINE=InnoDB;