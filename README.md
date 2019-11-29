# drserver

![alt text][logo]

![alt text][status]

serve digital resources (specifically [github.com/timdrysdale/dr](https://github.com/timdrysdale/dr))

## Background

Digital resources such as descriptions of experiments, user interfaces, and hardware often require to be dynamically updated. They need finite lifetimes, and some resources can be reused (given to multiple requesters), while others cannot. The logic for this goes just beyond typical key-value databases. The business rules for this are implemented in ```dr```, offering digital resources with expiry, reuse, category, id, and description. In this way, selections can be made between single-use resources without needing to (destructively) access the actual resource until the actual request for the specific desired resource is made.

## Usage

```drserver``` is a command line utility that runs an http server with a RESTlike API, backed by an native golang in-memory data store. 

```
$ drserver --listen 8085
```


## Extension

This is an example of how to create a ```drserver``` utility from ```dr```. Other API and data stores can be implemented in ```dr``` if desired.


[logo]: ./img/logo.png "logo for drserver, clock and files"
[status]: https://img.shields.io/badge/alpha-do%20not%20use-orange "Alpha status, do not use" 