# Invopop Merchandising Server

The main purpose of this server is to provide a small e-commerce platform
back-end. Defining an API to manage a virtual checkout basket. All the products and client's data is stored in memory; there is no DB connection of any kind.

## Usage

### Docker

The easiest way to try the server and client is with `docker`, clone the
repository, build the image and run it to get ther server running.

```
git clone https://github.com/dvdalilue/invopop.git
cd invopop

docker build -t invopop:v0.1 .
docker run --rm --name invopop -d -p 8080:8080 invopop:v0.1
```

*run without the `-d` flag if you like*

To start using the demo server client, do the following:

```
docker exec -it invopop /bin/sh
# now inside the container you can use the client
./client
```

### Go

Another option to try the server is to compile the code with `go`. There is a
`Makefile` to ease the building process. Run `make` and this will produce two
binaries `server` and `client`. Both can be run without passing any argument,
they will use the default configuration.

#### Server configuration

- `INVOPOP_SERVER_PORT`: server listening port (e.g. 8080)

#### Client configuration

- `INVOPOP_SERVER_URL`: url of the invopop server (e.g. http://localhost:8080)


## Test

To run the tests, run `make unit-test`