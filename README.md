# golang gotchas
Anything I didn't know about golang and is somehow special to golang for me with programming background in Java, Python, Groovy, C, C++ and misc other programming languages

Maybe it helps others starting to use golang :-)

1. [acceptLanguage](https://github.com/framps/golang_gotchas/blob/master/acceptlanguage/main.go) - Sample using golang.org/x/text/language to select the correct language from HTTP Accept-Language
2. [broadcastChan](https://github.com/framps/golang_gotchas/blob/master/broadcastChan/main.go) - Sample how to implement simple broadcasts
3. [compareMaps](https://github.com/framps/golang_gotchas/blob/master/compareMaps/main.go) -Typical pitfall when comparing map elements
4. [copyStructs](https://github.com/framps/golang_gotchas/blob/master/copyStructs/copyStructs.go) - Sample code using github.com/jinzhu/copier to copy structs
5. [debugHTTPClient](https://github.com/framps/golang_gotchas/blob/master/debugHTTPClient/main.go) - Sample code which uses httputil to get low level debugging of http clients
6. [errorMarshaling](https://github.com/framps/golang_gotchas/blob/master/errorMarshaling/main.go) - Sample code hoe marshaling of errors works
7. [fetchUrls](https://github.com/framps/golang_gotchas/blob/master/fetchURLs/fetchURLs.go) - Use gofuncs to execute get requests against a number of urls and calculate responsetime statistics
8. [github](https://github.com/framps/golang_gotchas/blob/master/github/googleGithubLib.go) - Retrieve all public repositories of a github organization with a home grown client or a client using the google go-github library
9. [gorohttp](https://github.com/framps/golang_gotchas/blob/master/gorohttp/goro.go) - Simple sample code how to kick off http requests as fast as possible
10. [httpStress](https://github.com/framps/golang_gotchas/blob/master/httpStress/main.go) - More complex sample code how to kick off http requests as fast as possible
11. [httpClient](https://github.com/framps/golang_gotchas/blob/master/httpclient/httpclient.go) - Sample how to use timeouts in a httpClient
12. [i18n](https://github.com/framps/golang_gotchas/blob/master/i18n/main.go) - Sample code using github.com/nicksnyder/go-i18n/i18n to enable localization
11. [jsonMarshalling](https://github.com/framps/golang_gotchas/blob/master/jsonMashalling/jsonMarshalling.go) - Print the different marshal/unmarshal results of different structures into/from JSON
12. [logging](https://github.com/framps/golang_gotchas/blob/master/logging/logging.go) - Sample code using go.uber.org/zap to create structured logs and use log hooks
13. [pointers](https://github.com/framps/golang_gotchas/blob/master/pointers/main.go) - Sample code to demonstrate a common error using pointers in golang when you are used to use C or C++
14. [pollURLs](https://github.com/framps/golang_gotchas/blob/master/pollURLs/pollURLs.go) - Golang pollURL sample code
15. [sigTerm](https://github.com/framps/golang_gotchas/blob/master/sigTerm/main.go) - Catch SIGTERM or SIGINT to gracefully shutdown program
16. [templates](https://github.com/framps/golang_gotchas/blob/master/templates/main.go) - Sample code using golang /text/template
17. utils
    1. [json.go](https://github.com/framps/golang_gotchas/blob/master/utils/json.go) - pretty print JSON
18. [waitGroupTimeout](https://github.com/framps/golang_gotchas/blob/master/waitGroupTimeout/waitGroupTimeout.go) - Sample which implements a waitGroup which can timeout
