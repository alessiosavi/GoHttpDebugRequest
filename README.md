# GoHttpDebugRequest
Debug the GET POST request information by sending them to this server

## Usage

A simple Go server delegated to print the information related to the given HTTP request.

It spawn different endpoint for every method:

 - GET (/get)
 - HEAD (/head)
 - POST (/post)
 - PUT (/put)
 - DELETE (/delete)
 - CONNECT (/connect)
 - OPTIONS (/options)
 - TRACE (/trace)
 - PATCH (/patch)
 
 In order to run the tool, you need to proviede two input parameter:
 - Port where bind the service (-port)
 - Host where bind the service (-host)
 
 If the parameter are not provided, it use `9999` as default port and `localhost` as default host.
 
 ## Help
 
 Usage of ./GoHttpDebugRequest:  
 - -host string  
        Bind the server to the given host (default "localhost")  
 - -port int  
        Bind the server to the given port (default 9999)  
