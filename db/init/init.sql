-- used to create the application database and user for the docker image
CREATE DATABASE example;
CREATE USER example_user WITH PASSWORD 'password';
GRANT CONNECT ON DATABASE example TO example_user;
GRANT pg_read_all_data TO example_user;
GRANT pg_write_all_data TO example_user;
