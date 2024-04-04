-- postgresql
CREATE TABLE IF NOT EXISTS characters (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  job VARCHAR(255) NOT NULL,
  coworker VARCHAR(255) NOT NULL,
  employer VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS report_card (
  id SERIAL PRIMARY KEY,
  character_id INT NOT NULL,
  year INT NOT NULL,
  FOREIGN KEY (character_id) REFERENCES characters(id)
);

CREATE TABLE IF NOT EXISTS skills (
  id SERIAL PRIMARY KEY,
  report_card_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  d4s INT NOT NULL,
  rerolls INT NOT NULL,
  FOREIGN KEY (report_card_id) REFERENCES report_card(id)
);

CREATE TABLE IF NOT EXISTS relationships (
  id SERIAL PRIMARY KEY,
  character_id INT NOT NULL,
  points INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  inspiration BIT NOT NULL,
  boon_or_bane VARCHAR(255) NOT NULL,
  FOREIGN KEY (character_id) REFERENCES characters(id)
);

CREATE TABLE IF NOT EXISTS extra_curriculars (
  id SERIAL PRIMARY KEY,
  character_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  d4 BIT NOT NULL,
  skills VARCHAR(255) NOT NULL,
  member VARCHAR(255) NOT NULL,
  FOREIGN KEY (character_id) REFERENCES characters(id)
);