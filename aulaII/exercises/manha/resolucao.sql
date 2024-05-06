-- mostre o título e o nome do gênero de todas as séries.

SELECT 
	series.title as 'series name',
  genres.name as 'genre' 
FROM series
INNER JOIN genres
ON series.genre_id = genres.id; 

-- mostre o título dos episódios, o nome e sobrenome dos atores que trabalham em cada um deles.

SELECT
	ep.title,
  ac.first_name,
  ac.last_name
FROM episodes as ep
INNER JOIN actor_episode as ae
ON ae.episode_id = ep.id
INNER JOIN actors as ac
ON ac.id = ae.actor_id;

-- mostre o título de todas as séries e o número total de temporadas que cada uma delas possui.

SELECT 
	series.title,
  COUNT(seasons.id) as 'seasons quantity'
FROM series
INNER JOIN seasons
ON seasons.serie_id = series.id
GROUP BY series.title;

-- mostre o nome de todos os gêneros e o número total de filmes de cada um, desde que seja maior ou igual a 3.

SELECT 
	genres.name,
  COUNT(movies.id) as 'movies count'
FROM genres
INNER JOIN movies
ON movies.genre_id = genres.id
GROUP BY genres.name
HAVING COUNT(movies.id) >= 3;

-- mostre apenas o nome e sobrenome dos atores |  que atuam em todos os filmes de Star Wars e que estes não se repitam.

SELECT DISTINCT
	ac.first_name,
  ac.last_name
FROM actors AS ac
INNER JOIN actor_movie AS am
ON am.actor_id = ac.id
INNER JOIN movies AS mo
ON mo.id = am.movie_id
WHERE mo.id IN(3,4);