start:
	docker-compose up -d
	docker-compose logs -f --tail 4000


clean:
	docker-compose down --rmi local
	docker volume prune
