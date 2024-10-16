ALTER TABLE users
ADD username VARCHAR(64) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);