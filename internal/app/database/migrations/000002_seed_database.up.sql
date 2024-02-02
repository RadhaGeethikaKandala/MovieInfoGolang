INSERT INTO ratings (source, value)
VALUES
  ('Internet Movie Database', '8.2/10'),
  ('Rotten Tomatoes', '85%'),
  ('Metacritic', '70/100');



INSERT INTO movies (
  title, year, rated, released,
  runtime, genre, director, writer, actors, plot,
  language, country, awards, poster, metascore, imdbrating, imdbvotes,
  imdbid, type, dvd, boxoffice, production,
  website, response)
VALUES
  ('Batman Begins', '2005', 'PG-13', '15 Jun 2005'
    , '140 min', 'Action, Crime, Drama', 'Christopher Nolan'
    , 'Bob Kane, David S. Goyer, Christopher Nolan'
    , 'Christian Bale, Michael Caine, Ken Watanabe'
    , 'After witnessing his parents death, Bruce learns the art of fighting to confront injustice. When he returns to Gotham as Batman, he must stop a secret society that intends to destroy the city.'
    , 'English, Mandarin'
    , 'United States, United Kingdom'
    , 'Nominated for 1 Oscar. 14 wins & 79 nominations total'
    , 'https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg'
    , '70', '8.2', '1,559,730'
    , 'tt0372784', 'movie'
    , '09 Sep 2009', '$206,863,479'
    , 'N/A', 'N/A', 'true');


INSERT INTO moviesratings ( movie_id , rating_id )
VALUES
  (1,1),
  (1,2),
  (1,3);



