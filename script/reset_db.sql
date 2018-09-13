DROP DATABASE IF EXISTS go_vue;
CREATE DATABASE go_vue;

DROP USER IF EXISTS 'app' ;
FLUSH privileges;
CREATE USER 'app' IDENTIFIED BY 'password' ;
GRANT ALL ON go_vue.* TO app@'localhost';