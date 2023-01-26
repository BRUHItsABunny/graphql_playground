CREATE SEQUENCE movie_id_seq;

CREATE TABLE movies(
    id BIGINT NOT NULL DEFAULT nextval('movie_id_seq') PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    release_date TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER SEQUENCE movie_id_seq OWNED BY movies.id;