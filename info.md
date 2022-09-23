# URLs
* https://github.com/techschool/simplebank
* https://discord.com/invite/BksFFXu

# Docker Commands
* sudo docker pull postgres:12-alpine
* sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
* sudo docker exec -it postgres12 psql -U root (type \q to quit)

# Install Tableplus
* https://tableplus.com/blog/2019/10/tableplus-linux-installation.html

# Install Migrate
* https://github.com/golang-migrate/migrate/releases
* Downlaod migrate.linux-386.deb
    * sudo dpkg -i migrate.linux-386.deb
    * create dbMigration
    * cd dbMigration && migrate create -ext sql -dir . -seq init_schema

# Install sqlc
    sudo snap install sqlc