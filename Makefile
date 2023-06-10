grpc_with_redis:
	rm -rf cmd/.env
	echo "DATABASE=redis" > cmd/.env
	echo "SERVICE=grpc" >> cmd/.env
	docker-compose up -d --build redis_db grpc_server

grpc_with_postgres:
	rm -rf cmd/.env
	echo "DATABASE=postgres" > cmd/.env
	echo "SERVICE=grpc" >> cmd/.env
	docker-compose up -d --build postgres_db grpc_server

tests:
	rm -rf cmd/.env
	echo "DATABASE=all" > cmd/.env
	echo "SERVICE=test" >> cmd/.env
	docker-compose up -d --build postgres_db redis_db test_service

rest_with_redis:
	rm -rf cmd/.env
	echo "DATABASE=redis" > cmd/.env
	echo "SERVICE=rest" >> cmd/.env
	docker-compose up -d --build redis_db rest_server

rest_with_postgres:
	rm -rf cmd/.env
	echo "DATABASE=postgres" > cmd/.env
	echo "SERVICE=rest" >> cmd/.env
	docker-compose up -d --build postgres_db rest_server