.PHONY: db-init
db-init:
	docker-compose down
	docker-compose up -d
	docker cp data.sql cockroach-25P02:/cockroach/.
	docker-compose exec -T cockroach cockroach sql --insecure < data.sql