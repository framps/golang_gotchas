# golang gotchas
Anything I didn't know about golang and is somehow special to golang for me

1. acceptLanguage - Sample using golang.org/x/text/language to select the correct language from HTTP Accept-Language
2. compareMaps -Typical pitfall when comparing map elements
3. copyStructs - Sample code using github.com/jinzhu/copier to copy structs
4. debugHTTPClient - Sample code which uses httputil to get low level debugging of http clients
5. errorMarshaling - Sample code hoe marshaling of errors works
6. fetchUrls - Use gofuncs to execute get requests against a number of urls and calculate responsetime statistics
7. github - Retrieve all public repositories of a github organization
8. gorohttp - Sample how to kick off http requests as fast as possible
9. httpClient - Sample how to use timeouts in a httpClient
10. i18n - Sample code using github.com/nicksnyder/go-i18n/i18n to enable localization
11. jsonMarshalling - Print the different marshal/unmarshal results of different structures into/from JSON
12. logging - Sample code using go.uber.org/zap to log code execution
13. pointers - Sample code for a common error using pointers in golang when you are used to use C or C++
14. pollURLs - Golang pollURL sample code
15. sigTerm - Catch SIGTERM or SIGINT to gracefully shutdown program
16. templates - Sample code using golang /text/template
17. utils
    1. json.go - pretty print JSON
18. waitGroupTimeout - Sample which implements a waitGroup which can timeout
