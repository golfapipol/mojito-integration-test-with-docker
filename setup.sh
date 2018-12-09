# SETUP
docker run -d -p 27017:27017 --volume /Users/golfapipol/sck/mojito:/tmp --name mongo mongo:3.2.10 
# Add Test Data
docker exec -it mongo mongo mongodb://localhost:27017/mojito /tmp/init-seed.js
# Test
go test ./...
# Remove Test Data & remove environment
docker rm -f mongo