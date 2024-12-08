new-db :
	cp -r template $(word 2, $(MAKECMDGOALS))

migrate-fresh :
	dbmate drop
	./migrate-fresh.sh