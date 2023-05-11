CREATE TABLE IF NOT EXISTS roles (
  id VARCHAR(200) PRIMARY KEY,
  role VARCHAR(255) NOT NULL,
  visibility ENUM('Pemustaka', 'Petugas') NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
)ENGINE=InnoDB;