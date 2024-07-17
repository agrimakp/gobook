### Fetching a URL

Go Provides a collection of packages, grouped under **net**, that make it easy to send and recieve information though the internet,make low level network connections, and set up servers, for which Go's concurrency features are particularly useful.

http.Get function makes an HTTP request and, if there is no error, returns the result in the response struct resp

The body field of resp contains the server response as a readable stream.

ioutil.ReadAll reads the entire response, result is stored in b.

The body stream is closed to avoid leaking resourses