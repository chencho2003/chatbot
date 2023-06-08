CREATE TABLE userdata (
userName NOT NULL,
email varchar(45) NOT NULL,
password_hash VARCHAR(255) NOT NULL,
PRIMARY KEY (userName),
UNIQUE (Email)
)