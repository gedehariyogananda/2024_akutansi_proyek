### DB Standalone dbmate for akutansi_proyek

****PREREQUISITES****
- [Dbmate](https://github.com/amacneil/dbmate)
- Makefile (optional)
- Postgres

run from makefile, or manually refered to ./Makefile

****HOW TO USE****

**project level**
- create new section

  `make new-db users`


- initial migration

  `make migrate-fresh`


**section level**
- create new table

    `make create companies`


- modify table

    `make modify companies`


- reload migration

    `make migrate-reload`


- rollback migration

    `make down`


- commit migration

    `make up`


happy migrate ❤️