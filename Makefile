migration:
	@read -p "migration name: " name;\
		migrate create -ext sql -dir migrations -seq $$name
up:
	migrate -path migrations \
		-database "mysql://user:prosklad@tcp(127.0.0.1:3306)/prosklad" -verbose up
down:
	migrate -path migrations \
		-database "mysql://user:prosklad@tcp(127.0.0.1:3306)/prosklad" -verbose down
force:
	@read -p "Enter version to force :" version;\
		migrate -path migrations \
			-database "mysql://user:prosklad@tcp(127.0.0.1:3306)/prosklad" -verbose force $$version

