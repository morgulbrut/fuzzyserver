# fuzzyserver

Dead simple webserver in go that just throw back random http status codes.

Listens on port 1337 as default, port can be defined using the `-port` flag.

* `localhost:1337` gives back a random status code from the whole range.
* `localhost:1337/100` gives back a random status code from the 100 range.
* `localhost:1337/200` gives back a random status code from the 200 range.
* `localhost:1337/300` gives back a random status code from the 300 range.
* `localhost:1337/400` gives back a random status code from the 400 range.
* `localhost:1337/500` gives back a random status code from the 500 range.
