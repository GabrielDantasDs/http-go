## How http server work
- web server is just a computer that stores a application, like a website, more specifically the application files and where it run
- a web server control how web users interact with server files
- A **Http server** is as software that understand URLS and HTTP protocol
- Whanever a browser need a file that is storage in serve, the broswer send a request to server
- To publish a website, you need neither a static or a dynamic web server

**Static web server**: Send files as-is to browser
**Dynamic web server**: Server update files before send to browser, files is processed by a extra software like a **application**

- HTTP follows a classical [client-server-model](##client-server-model)

## Communicating through HTTP

- HTTP: Hypertext transfer protocol
- Specifies how transfer hypertext (linked web documents) between two computers
- Textual, stateless protocol (dont keep any info between "messages")
- Clients make request, server respond
- Client must provied file URL
- Server must awser every HTTP request, at least with an error message


### Communicating steps
1. Upon receiving a request, an HTTP server checks if the request URL matches an existing file
2. If so, server send the file content back to browser, if not the server will check if it should send generate a file dynamically
3. If neither these options, so the web server returns a error message, base on HTTP status


## Client-server-model
Is a distributed application structure that partitions tasks or workloads between the providers of a resource or service, called servers, annd services requesters called clients


## HTTP Headers
- Used to send metadata about resource or/and a HTTP message.


## HTTP Methods
- Indicate the purpose of the request and what is expected if the request is succcessful