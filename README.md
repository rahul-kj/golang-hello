hello
=====

This is sample project for learning Golang.

Project needs gox for cross compilation

Important commands that you want to run from src dir:

```
go get github.com/rahulkj/server
go get github.com/mitchellh/gox
go build github.com/rahulkj/server
```

If you want to build something specific to your OS, you can use the following command:
```
gox -osarch="linux/amd64" github.com/rahulkj/server
```

else to build for all platforms
```
gox github.com/rahulkj/server
```

Once done you will have the binaries generated in src folder.
Copy the binary and rename it to server and place it in the following location:

```
staging/bin/
````

Steps to deploy to PCF:
```
cd staging
gcf api https://api.run.pivotal.io
gcf login
gcf push server -b http://github.com/ryandotsmith/null-buildpack.git -c ./bin/server
```

That's it!!

