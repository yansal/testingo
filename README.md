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

Names defined in `*_test.go` files are not included in the build: e.g. `TestPrivate` is not present in the output.

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

Quoting from https://golang.org/pkg/testing/: `The file will be excluded from regular package builds`.
