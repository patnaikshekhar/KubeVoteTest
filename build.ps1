$env:GOOS = "linux"

echo "Building Go Program"
go build

echo "Building Container"
docker-compose build

echo "Removing Existing Containers"
docker-compose rm -vf 

echo "Starting Stack"
docker-compose up