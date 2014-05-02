rm -rf staging/bin
mkdir -p staging/bin

cd src

go get github.com/rahulkj/server
go get github.com/mitchellh/gox
go build github.com/rahulkj/server

gox -osarch="linux/amd64" github.com/rahulkj/server

cp server_linux_amd64 ../staging/bin/server

cd ../staging
gcf push server -b http://github.com/ryandotsmith/null-buildpack.git -c ./bin/server
