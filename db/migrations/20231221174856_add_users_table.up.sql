CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(60),
    password VARCHAR(60)
);

CREATE TABLE sound_library( 
    sound_id SERIAL PRIMARY KEY,
    user_id integer,
    sound_path varchar(80),
    file_ext varchar(5),
    file_size int,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
