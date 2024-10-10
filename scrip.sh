docker compose build
docker compose create
docker compose start
docker compose down

docker compose down --rmi all



docker build -t koriebruh/apitest .

docker images #size lebih kecil cek

docker container create --name api1 -p 8080:8080 koriebruh/apitest

docker start api1
