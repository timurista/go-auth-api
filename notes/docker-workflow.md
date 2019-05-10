## Virtualization
Docker workflow

Vm OS, virtual hardware, application, which runs on real hardware
less volume of private copy. Isolated NW-interface, process space, file system

Container is running state of image
## Docker ps
forward of local terminal -ti

troubleshooting:
```
Step 5/8 : RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a
 ---> Running in 9649dca56d8b
main.go:15:2: cannot find package "github.com/davecgh/go-spew/spew" in any of:
        /usr/local/go/src/github.com/davecgh/go-spew/spew (from $GOROOT)
        /go/src/github.com/davecgh/go-spew/spew (from $GOPATH)
main.go:16:2: cannot find package "github.com/dgrijalva/jwt-go" in any of:
        /usr/local/go/src/github.com/dgrijalva/jwt-go (from $GOROOT)
        /go/src/github.com/dgrijalva/jwt-go (from $GOPATH)
main.go:17:2: cannot find package "github.com/gorilla/mux" in any of:
        /usr/local/go/src/github.com/gorilla/mux (from $GOROOT)
        /go/src/github.com/gorilla/mux (from $GOPATH)
driver/driver.go:9:2: cannot find package "github.com/lib/pq" in any of:
        /usr/local/go/src/github.com/lib/pq (from $GOROOT)
        /go/src/github.com/lib/pq (from $GOPATH)
main.go:18:2: cannot find package "github.com/subosito/gotenv" in any of:
        /usr/local/go/src/github.com/subosito/gotenv (from $GOROOT)
        /go/src/github.com/subosito/gotenv (from $GOPATH)
main.go:19:2: cannot find package "golang.org/x/crypto/bcrypt" in any of:
        /usr/local/go/src/golang.org/x/crypto/bcrypt (from $GOROOT)
        /go/src/golang.org/x/crypto/bcrypt (from $GOPATH)
The command '/bin/sh -c cd ${SOURCES} && CGO_ENABLED=0 go install -a' returned a non-zero code: 1
```
