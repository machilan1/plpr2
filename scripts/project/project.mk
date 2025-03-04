## newmovie: make new movie
.PHONY: newmovie
newmovie:
	@echo 'create a movie of env BODY...'
	curl -i -d ${MOVIE} http://localhost:30000/v1/movies

## querymovie: find movie, movieID=$1
.PHONY: querymovie
querymovie:
	@echo 'find a movie of...'
	curl  'localhost:30000/v1/movies/${movieID}'