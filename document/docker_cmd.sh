# sh
docker run -idt --name shrimp_blog -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

docker run -idt --name shrimp_blog_redis -p 6379:6379 redis:latest