CREATE TABLE userData (
    userName VARCHAR(45) NOT NULL,
    email VARCHAR(45) NOT NULL,
    pword VARCHAR(45) NOT NULL,
    PRIMARY KEY (userName),
    UNIQUE (email)
)