.. Autogenerated by Gollum RST generator (docs/generator/*.go)

HTTPRequest
===========


The HTTPRequest producer sends messages as HTTP requests to a given webserver.


* If RawData = true, the incoming messages are expected to contain complete
  HTTP requests in "wire format", such as::

    POST /foo/bar HTTP/1.0\n
    Content-type: text/plain\n
    Content-length: 24
    \n
    Dummy test\n
    Request data\n

  In this mode, the message's contents is parsed as an HTTP request and
  sent to the destination server (virtually) unchanged. If the message
  cannot be parsed as an HTTP request, an error is logged. Only the scheme,
  host and port components of the "Address" URL are used; any path and query
  parameters are ignored. The "Encoding" parameter is ignored.


* If RawData = false, a POST request is made to the destination server
  for each incoming message, using the complete URL in "Address". The
  incoming message's contents are delivered in the POST request's body
  and Content-type is set to the value of "Encoding"




Parameters
----------

**Address**
defines the URL to send http requests to. Set by default
to "http://localhost:80". If the value doesn't contain "://",
it is prepended with "http://", so short forms like "localhost:8088"
are accepted.


**RawData**
chooses how to interpret and relay the incoming data


**Encoding**
defines the payload encoding when RawData is set to false.
Set to "text/plain; charset=utf-8" by default.


Parameters (from SimpleProducer)
--------------------------------

**Enable**
switches the consumer on or off. By default this value is set to true.


**ID**
allows this producer to be found by other plugins by name. By default this
is set to "" which does not register this producer.


**Channel**
sets the size of the channel used to communicate messages. By default
this value is set to 8192.


**ChannelTimeoutMs**
sets a timeout in milliseconds for messages to wait if this
producer's queue is full.
A timeout of -1 or lower will try the fallback route without notice.
A timeout of 0 will block until the queue is free. This is the default.
A timeout of 1 or higher will wait x milliseconds for the queues to become
available again. If this does not happen, the message will be send to the
retry channel.


**ShutdownTimeoutMs**
sets a timeout in milliseconds that will be used to detect
a blocking producer during shutdown. By default this is set to 1 second.
Decreasing this value may lead to lost messages during shutdown. Increasing
this value will increase shutdown time.


**Router**
contains either a single string or a list of strings defining the
message channels this producer will consume. By default this is set to "*"
which means "listen to all routers but the internal".


**FallbackStream**
defines the stream used for messages that cannot be delivered
e.g. after a timeout (see ChannelTimeoutMs). By default this is "".


**Formatter**
sets a formatter to use. Each formatter has its own set of options
which can be set here, too. By default this is set to format.Forward.
Each producer decides if and when to use a Formatter.


**Filter**
sets a filter that is applied before formatting, i.e. before a message
is send to the message queue. If a producer requires filtering after
formatting it has to define a separate filter as the producer decides if
and where to format.


Example
-------

.. code-block:: yaml

	 - "producer.HTTPRequest":
	   RawData: true
	   Encoding: "text/plain; charset=utf-8"
	   Address: "http://localhost:80"
	


