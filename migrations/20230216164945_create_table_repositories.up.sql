CREATE TABLE IF NOT EXISTS repositories (
  id VARCHAR(200) PRIMARY KEY,
  collection_id VARCHAR(200) NOT NULL,
  departement_id VARCHAR(200) NOT NULL,
  category_id VARCHAR(200) NOT NULL,
  title VARCHAR(255) NOT NULL,
  abstract TEXT,
  improvement ENUM('1', '0') DEFAULT '0',
  related_title VARCHAR(255),
  update_desc TEXT,
  date_validated CHAR(10),
  status ENUM('pending', 'approved', 'denied') DEFAULT 'pending',
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (collection_id) REFERENCES collections(id),
  CONSTRAINT FOREIGN KEY (departement_id) REFERENCES departements(id),
  CONSTRAINT FOREIGN KEY (category_id) REFERENCES categories(id)
)ENGINE=InnoDB;