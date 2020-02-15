# Setup server on ubuntu 18.04 with go 1.12
sudo apt update
sudo apt dist-upgrade -y
sudo apt install wireshark
sudo apt autoremove
# Add these if there isn't a non-root user
# sudo useradd rj -m
# cd /home/rj
# Via https://tecadmin.net/install-go-on-ubuntu/
wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
tar -xvf go1.12.7.linux-amd64.tar.gz
sudo mv go /usr/local
rm go1.12.7.linux-amd64.tar.gz
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
mkdir go
go get github.com/gorilla/mux
# Download bpfcheck.go
# go run bpfcheck.go
echo "alias checkbpf='sudo /usr/local/go/bin/go run $HOME/go/checkbpf.go & echo $! > ~/go/checkbpf.pid'" >> ~/.bashrc
echo "alias recheck='sudo ps ux | grep [c]heckbpf | awk {"print \$2"} | xargs -I {} sudo kill {}; checkbpf'" >> ~/.bashrc
source ~/.bashrc

