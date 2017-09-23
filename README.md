# testingo

It's OK to use private names inside unit tests.

    $ go test -v github.com/yansal/testingo
    === RUN   TestPrivate
    private
    --- PASS: TestPrivate (0.00s)
    === RUN   TestPublic
    Public
    private
    --- PASS: TestPublic (0.00s)
    PASS
    ok  	github.com/yansal/testingo	0.001s

Quoting from https://golang.org/cmd/go/: `When compiling packages, build ignores files that end in '_test.go'`.

Quoting from https://golang.org/pkg/testing/: `The file will be excluded from regular package builds`.

Note that `*_test.go` files are excluded from the compile command.

    $ go build -x github.com/yansal/testingo
    WORK=/var/folders/lc/ly7d57k1023569jxb4y2kq200000gn/T/go-build535115465
    mkdir -p $WORK/github.com/yansal/testingo/_obj/
    mkdir -p $WORK/github.com/yansal/
    cd /Users/yann/go/src/github.com/yansal/testingo
    /usr/local/Cellar/go/1.9/libexec/pkg/tool/darwin_amd64/compile -o $WORK/github.com/yansal/testingo.a -trimpath $WORK -goversion go1.9 -p github.com/yansal/testingo -complete -buildid 67bb8f6c70ff508a9fa1121d4edbb3f3840b2a94 -D _/Users/yann/go/src/github.com/yansal/testingo -I $WORK -pack ./testingo.go

Note that names defined in `*_test.go` files are not included in the build: e.g. `TestPrivate` is not present in the `nm`'s output.

    $ go install github.com/yansal/testingo && go tool nm $GOPATH/pkg/darwin_amd64/github.com/yansal/testingo.a
         U 
     63c T %22%22.Public
     6d4 T %22%22.init
     7d3 B %22%22.initdone·
     5ab T %22%22.private
     7b3 R %22%22.statictmp_0
     7c3 R %22%22.statictmp_1
         U fmt.Println
         U fmt.init
     99a R gclocals·33cdeccccebe80329f1fdbee7f5874cb
     988 R gclocals·69c1753bd5f81501d95132d08af04464
     990 R gclocals·e226d4ae4a7cad8835311c6a4683c14f
     777 ? go.info.%22%22.Public
     796 ? go.info.%22%22.init
     751 ? go.info.%22%22.private
     796 ? go.range.%22%22.Public
     7b3 ? go.range.%22%22.init
     771 ? go.range.%22%22.private
     771 R go.string."Public"
     74a R go.string."private"
         U runtime.algarray
     7d3 R runtime.gcbits.01
     81c R runtime.gcbits.03
         U runtime.morestack_noctxt
         U runtime.throwinit
     902 R type.*[1]interface {}
     87f R type.*[]interface {}
     7e4 R type.*interface {}
     982 R type..importpath.fmt.
     8ef R type..namedata.*[1]interface {}-
     86d R type..namedata.*[]interface {}-
     7d4 R type..namedata.*interface {}-
     93a R type.[1]interface {}
     8b7 R type.[]interface {}
     81d R type.interface {}
         U type.string



