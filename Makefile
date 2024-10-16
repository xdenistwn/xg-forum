# run golang main.go
run:
	docker exec -it xgforum-golang sh -c "go run cmd/main.go"
tidy:
	docker exec -it xgforum-golang sh -c "go mod tidy"
get:
	docker exec -it xgforum-golang sh -c "go get $(url)"

# migrate
export MYSQL_URL='mysql://root:changeThisPassword@tcp(xgforum-mysql-db:3306)/xgforum'

migrate-create:
	docker exec -it xgforum-golang sh -c "migrate create -ext sql -dir database/migrations -seq $(name)"
migrate-up:
	docker exec -it xgforum-golang sh -c "migrate -database ${MYSQL_URL} -path database/migrations up $(n)"
migrate-down:
	docker exec -it xgforum-golang sh -c "migrate -database ${MYSQL_URL} -path database/migrations down $(n)"
migrate-fix:
	docker exec -it xgforum-golang sh -c "migrate -path database/migrations -database ${MYSQL_URL} force ${version}"