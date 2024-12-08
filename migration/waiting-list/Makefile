reload :
	cp ../.env .env
	dbmate down
	dbmate up

create :
	dbmate new create_$(word 2, $(MAKECMDGOALS))_table

modify :
	dbmate new modify_$(word 2, $(MAKECMDGOALS))_table

up :
	cp ../.env .env
	dbmate up

down :
	cp ../.env .env
	dbmate down