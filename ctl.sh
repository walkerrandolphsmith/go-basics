command=$1

case $command in
  build)
    docker build -t my-golang-app .
    ;;
  up)
    docker run -it --rm --name my-running-app my-golang-app
    ;;
  *)
    echo "Command Unknown"
    exit 1
esac