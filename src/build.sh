go build github.com/rahulkj/server/
gox -osarch="linux/amd64" github.com/rahulkj/server/
cp server_linux_amd64 ../staging/bin/server
