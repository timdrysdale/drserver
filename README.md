# drserver

![alt text][logo]

![alt text][status]

serve digital resources (specifically [github.com/timdrysdale/dr](https://github.com/timdrysdale/dr))

## Background

Digital resources such as descriptions of experiments, user interfaces, and hardware often require to be dynamically updated. They need finite lifetimes, and some resources can be reused (given to multiple requesters), while others cannot. The logic for this goes just beyond typical key-value databases. The business rules for this are implemented in ```dr```, offering digital resources with expiry, reuse, category, id, and description. In this way, selections can be made between single-use resources without needing to (destructively) access the actual resource until the actual request for the specific desired resource is made.

## Usage

```drserver``` is a command line utility that runs an http server with a RESTlike API, backed by an native golang in-memory data store. 

```
$ drserver --listen 8086

```
Healthcheck ok?
```
$ curl -X GET http://localhost:8086/api/healthcheck -w "\n"
{"status":"ok"}
```
Load a resource with a short lifetime
```
$ curl -X POST -H "Content-Type: application/json" -d '{"ID":"0","Description":"apple","Category":"food","Resource":"secret!","Reusable":true,"TTL":5}' http://localhost:8086/api/resources/food/0
```
Check it is there
```
$ curl -X GET http://localhost:8086/api/resources -w "\n"
{"food":1}
```
Wait 5sec for resource to expire...
```
$ curl -X GET http://localhost:8086/api/resources -w "\n"
Storage is empty

```

Now load two objects into a category in a single call 
```
$ curl -X POST -H "Content-Type: application/json" -d '{"0":{"ID":"0","Description":"apple","Category":"food","Resource":"apple-secret!","Reusable":true,"TTL":60},"1":{"ID":"1","Description":"pear","Category":"food","Resource":"pear-secret!","Reusable":false,"TTL":60}}' http://localhost:8086/api/resources/food
$ curl -X GET http://localhost:8086/api/resources -w "\n"
{"food":2}
```

Compare resuable and non-resuable objects - non-reusable objects can only be obtained once

```
$ curl -X GET http://localhost:8086/api/resources/food/0 -w "\n"
{"Category":"food","Description":"apple","ID":"0","Resource":"apple-secret!","Reusable":true,"TTL":50}
$ curl -X GET http://localhost:8086/api/resources/food/0 -w "\n"
{"Category":"food","Description":"apple","ID":"0","Resource":"apple-secret!","Reusable":true,"TTL":49}
$ curl -X GET http://localhost:8086/api/resources/food/1 -w "\n"
{"Category":"food","Description":"pear","ID":"1","Resource":"pear-secret!","Reusable":false,"TTL":48}
$ curl -X GET http://localhost:8086/api/resources/food/1 -w "\n"
Resource not found
$ curl -X GET http://localhost:8086/api/resources/food/0 -w "\n"
{"Category":"food","Description":"apple","ID":"0","Resource":"apple-secret!","Reusable":true,"TTL":47}
```

Compare resources using the description - the resource (typically secret) is redacted from the list of items returned from GET to a ```/category```.
```
$ curl -X GET http://localhost:8086/api/resources/food -w "\n"
{"0":{"Category":"food","Description":"apple","ID":"0","Resource":"","Reusable":false,"TTL":33},"1":{"Category":"food","Description":"pear","ID":"1","Resource":"","Reusable":false,"TTL":33}}
$ curl -X GET http://localhost:8086/api/resources/food/1 -w "\n"
{"Category":"food","Description":"pear","ID":"1","Resource":"pear-secret!","Reusable":false,"TTL":27}
```


## Extension

This is an example of how to create a ```drserver``` utility from ```dr```. Other API and data stores can be implemented in ```dr``` if desired.


[logo]: ./img/logo.png "logo for drserver, clock and files"
[status]: https://img.shields.io/badge/alpha-do%20not%20use-orange "Alpha status, do not use" 