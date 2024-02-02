INSERT INTO ratings (source, value)
VALUES
  ('Internet Movie Database', '8.2/10'),
  ('Rotten Tomatoes', '85%'),
  ('Metacritic', '70/100'),
  ('Internet Movie Database', '7.8/10'),
  ('Rotten Tomatoes', '85%'),
  ('Metacritic', '72/100');


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
    , 'N/A', 'N/A', 'true'),
  ('The Batman', '2022', 'PG-13', '04 Mar 2022'
    , '176 min', 'Action, Crime, Drama', 'Matt Reeves'
    , 'Matt Reeves, Peter Craig, Bob Kane'
    , 'Robert Pattinson, Zoë Kravitz, Jeffrey Wright'
    , 'When a sadistic serial killer begins murdering key political figures in Gotham, Batman is forced to investigate the citys hidden corruption and question his familys involvement.'
    , 'English, Spanish, Latin, Italian'
    , 'United States'
    , 'Nominated for 3 Oscars. 35 wins & 171 nominations total'
    , 'https://m.media-amazon.com/images/M/MV5BM2MyNTAwZGEtNTAxNC00ODVjLTgzZjUtYmU0YjAzNmQyZDEwXkEyXkFqcGdeQXVyNDc2NTg3NzA@._V1_SX300.jpg'
    , '72', '7.8', '765,024'
    , 'tt1877830', 'movie'
    , '19 Apr 2022', '$369,345,583'
    , 'N/A', 'N/A', 'true');



INSERT INTO moviesratings ( movie_id , rating_id )
VALUES
  (1,1),
  (1,2),
  (1,3),
  (2,4),
  (2,5),
  (2,6);


