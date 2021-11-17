> 本地调试中，只需要启动mysql即可；

``` yaml
# docker-compose.yml

version: '3.1'

services:
  db:
    image: mysql:5.7
    command:
    - --default-authentication-plugin=mysql_native_password
    - --collation-server=utf8mb4_unicode_ci
    - --character-set-server=utf8mb4
    restart: always
    environment:
      MYSQL_DATABASE: 'jelly'
      #MYSQL_USER: 'root'
      MYSQL_ROOT_PASSWORD: '123456'
    ports:
    - 3306:3306
    expose:
    - 3306
    volumes:
    - ./my-db:/var/lib/mysql
```