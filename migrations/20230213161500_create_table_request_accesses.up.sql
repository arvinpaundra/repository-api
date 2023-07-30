CREATE TABLE IF NOT EXISTS request_accesses (
  id VARCHAR(200) PRIMARY KEY,
  pemustaka_id VARCHAR(200) NOT NULL,
  support_evidence VARCHAR(255) NOT NULL,
  status ENUM('pending', 'accepted', 'denied') DEFAULT 'pending',
  reasons TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (pemustaka_id) REFERENCES pemustakas(id)
)ENGINE=InnoDB;