CREATE TABLE IF NOT EXISTS staffs (
  id VARCHAR(200) PRIMARY KEY,
  user_id VARCHAR(200) NOT NULL,
  role_id VARCHAR(200) NOT NULL,
  fullname VARCHAR(255) NOT NULL,
  nip CHAR(18),
  telp CHAR(13),
  address TEXT,
  gender ENUM('Pria', 'Wanita', '') DEFAULT '',
  birth_date VARCHAR(10),
  is_active ENUM('1', '0') DEFAULT '0',
  avatar VARCHAR(255),
  signature VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT FOREIGN KEY (role_id) REFERENCES roles(id)
)ENGINE=InnoDB;