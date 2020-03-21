


function build(){
  mkdir -p build
  export GOOS=$1
  export GOARCH=$2
  mkdir -p build/$GOOS
  FN=$(cd `dirname $0` && pwd)/build/${GOOS}/reconciler
  go build -o "$FN" main.go
  echo $FN
}


DARWIN_BUILD=`build darwin amd64`
LINUX_BUILD=`build linux amd64`

fn=~/josh-env/darwin/logs/brew-cask.txt
brew cask list | $DARWIN_BUILD $fn