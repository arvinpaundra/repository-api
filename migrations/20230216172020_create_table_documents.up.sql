CREATE TABLE IF NOT EXISTS documents (
  id VARCHAR(200) PRIMARY KEY,
  repository_id VARCHAR(200) NOT NULL,
  validity_page VARCHAR(255) NOT NULL,
  cover_and_list_content VARCHAR(255) NOT NULL,
  chp_one VARCHAR(255) NOT NULL,
  chp_two VARCHAR(255) NOT NULL,
  chp_three VARCHAR(255) NOT NULL,
  chp_four VARCHAR(255) NOT NULL,
  chp_five VARCHAR(255) NOT NULL,
  bibliography VARCHAR(255) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT FOREIGN KEY (repository_id) REFERENCES repositories(id) ON DELETE CASCADE
)ENGINE=InnoDB;