CREATE TABLE jobs (
  id VARCHAR(36) PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  position_level_id INT NOT NULL,
  employment_type_id INT NOT NULL,
  status BOOL NOT NULL DEFAULT TRUE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_position_level FOREIGN KEY (position_level_id)
      REFERENCES position_levels (id)
      ON DELETE CASCADE,
  CONSTRAINT fk_employment_type FOREIGN KEY (employment_type_id)
      REFERENCES employment_types (id)
      ON DELETE CASCADE,
  INDEX idx_status_jobs (status)
);