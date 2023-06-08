CREATE SEQUENCE bot_seq
START 1
INCREMENT 1
OWNED BY bot.id;


-- CREATE TABLE bot(
--     id SERIAL,
--     question TEXT NOT NULL,
--     answer TEXT NOT NULL,
--     PRIMARY KEY(id)
-- );