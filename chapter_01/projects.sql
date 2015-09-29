CREATE TABLE projects (id INTEGER NOT NULL PRIMARY KEY, name VARCHAR(50),
updated DATETIME DEFAULT CURRENT_TIMESTAMP, language VARCHAR(50),
short_description VARCHAR(150), likes INTEGER DEFAULT 0,
shares INTEGER DEFAULT 0, created DATETIME DEFAULT CURRENT_TIMESTAMP);

INSERT INTO projects (name, language, short_description) VALUES ("RestLib", "python", "Really dynamic and slow library");