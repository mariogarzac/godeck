CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(60),
    password VARCHAR(60)
);

CREATE TABLE sound_library( 
    sound_id SERIAL PRIMARY KEY,
    user_id integer,
    game_date DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
