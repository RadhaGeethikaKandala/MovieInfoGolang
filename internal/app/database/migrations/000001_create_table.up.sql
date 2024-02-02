
CREATE TABLE IF NOT EXISTS ratings (
  id SERIAL PRIMARY KEY NOT NULL,
  source VARCHAR(255),
  value VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS movies (
  id SERIAL PRIMARY KEY NOT NULL,
  title VARCHAR(255),
  year VARCHAR(255),
  rated VARCHAR(255),
  released VARCHAR(255),
  runtime VARCHAR(255),
  genre VARCHAR(255),
  director VARCHAR(255),
  writer VARCHAR(255),
  actors VARCHAR(255),
  plot VARCHAR(255),
  language VARCHAR(255),
  country VARCHAR(255),
  awards VARCHAR(255),
  poster VARCHAR(255),
  metascore VARCHAR(255),
  imdbrating VARCHAR(255),
  imdbvotes VARCHAR(255),
  imdbid VARCHAR(255),
  type VARCHAR(255),
  dvd VARCHAR(255),
  boxoffice VARCHAR(255),
  production VARCHAR(255),
  website VARCHAR(255),
  response VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS moviesratings (
  movie_id INTEGER NOT NULL,
  rating_id INTEGER NOT NULL,
  constraint pk PRIMARY KEY (movie_id, rating_id),
  constraint fk_movies foreign key (movie_id) REFERENCES movies (id),
  constraint fk_ratings foreign key (rating_id) REFERENCES ratings (id)
);


