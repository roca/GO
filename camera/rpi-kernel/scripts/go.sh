# sudo apt-get install -y golang-go

wget https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin:${HOME}/go/bin

sudo apt-get install libopencv-dev

go get -u -d gocv.io/x/gocv
make deps
make download
make build

# export GOPATH="/go"
# export GOARCH="arm"
# export GOOS="linux"
# export GOARM=7

