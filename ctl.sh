command=$1

case $command in
  build)
    # docker build -t my-golang-app .
    docker-compose build api
    ;;
  up)
    docker-compose up
    # docker run -p 127.0.0.1:8080:8080 -it --rm --name my-running-app my-golang-app go run main.go
    ;;
  *)
    echo "Command Unknown"
    exit 1
esac