# Bitacora Forum

### What does this project consist of?
This project consists of a forum where users can share the experiences they have had in a specific country.

### Who is the target audience?
To all those people who want to share their stories in a foreign country, and to those who are curious to see how others have experienced it.

### How will the user interact with the application?
Users will have a country filter that they can use to search for the different associated forums. Users will not have to register to navigate the application, but will have to do so if they want to open a “thread” or post a comment.


### Various ideas:
- Include a visualization system, among other things, to create more filters and add a section on the home page with the most popular forums from around the world.
- Tag and follower system.

## Deploy development environment

Once the repository has been cloned and moved to its directory, add a ``conf/`` folder where we will configure the following files with their respective environment variables.

#### api.env
```env title=api.env
DSN=user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
```

#### mysql.env
```env
MYSQL_ALLOW_EMPTY_PASSWORD=yes
MYSQL_DATABASE=<db_name>
MYSQL_USER=<username>
MYSQL_PASSWORD=<password>
MYSQL_ROOT_PASSWORD=''
MYSQL_ROOT_HOST=%
```

#### phpmyadmin.env
```env
PMA_ARBITRARY=1
PMA_PORT=3306
PMA_HOST=mariadb
```

#### client.env
```env
# AT THE MOMENT EMPTY
```

Finally, run ``docker compose up`` and you are done. Maybe the first time the api may throw an error because you have no database user. Run ``docker exec -it mysql-db sh`` and once there, access the database and add it. You can also access phpmyadmin at ``http://localhost:8005`` and add it there. Now it should work.

To use the application “at its best” know that in ``api/scripts`` there are different scripts to insert false data in mysql.

There may be network conflicts (among others). Feel free to modify ``docker-compose.yml`` to resolve them.