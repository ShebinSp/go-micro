* If any permission errors to access ./db-data/postgres or /var/lib/postgresql/data/
* follow the commands:
sudo mkdir -p ./db-data/postgres/
sudo chown -R $(id -u):$(id -g) ./db-data/postgres/
sudo chmod -R 700 ./db-data/postgres/
