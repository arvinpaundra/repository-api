CREATE TABLE IF NOT EXISTS collections (
  id VARCHAR(200) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  visibility ENUM('Dosen', 'Mahasiswa', 'Semua') DEFAULT 'Semua',
  created_at TIMESTAMP,
  updated_at TIMESTAMP
)ENGINE=InnoDB;