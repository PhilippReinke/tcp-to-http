# HTTP simple

## Intro

The following examples illustrate how to respond and request with HTTP

### Responding

Open two shells and run

```sh
# shell 1
# Listen for TCP connection on port 8080.
# -l = listen
# -k = do not close when client closes connection
nc -kl 8080

# shell 2
# Send an HTTP request to localhost:8080
curl localhost:8080
```

The outcome is as follows:

```
GET / HTTP/1.1
Host: localhost:8080
User-Agent: curl/8.7.1
Accept: */*

```

Now respond with

```
HTTP/1.1 200 OK
content-length: 11

i like http

```

Congrats 🎉 you manually responded to an HTTP request via a TCP connection.

### Requesting

Lookup the IP of you favoruite webpage via
[DNS lookup](https://toolbox.googleapps.com/apps/dig/#A/) and establish a TCP
connection on port 80

```sh
# e.g. www.google.com has 64.233.164.103
nc 64.233.164.103 80
```

and enter

```
GET / HTTP/1.1
host: www.google.com

```

This returns something like

```
HTTP/1.1 200 OK
Expires: -1
Cache-Control: private, max-age=0
Content-Type: text/html; charset=ISO-8859-1
(... more)

42ab
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="de"><head><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google</title><script nonce="OAdrc6HsJtrEHoXc1ruthQ">(function(){var _g={kEI:'1tONaPXfNfqi1fIPlZm8sQI',kEXPI:'0,202854,2,38,4037151,78813,16105,201864,142932,238458,51586,5241682,36812642,25228681,138268,14109,4564,4381,3004,45182,8040,6754,23879,9139,4599,328,6226,1117,53092,9956,15048,8205,7430,58708,5,48901,5308,352,11059,7821,5870,5930,1784,5773,21644,5968,5556,10968,2808,453,2990,35,942,2478,5538,7946,8134,3973,5683,3605,593,6736,8899,1543,3366,5795,9470,649,678,3547,3,6881,338,1,5610,3821,1018,1,3459,2,217,696,2,3269,7,487,764,728,1032,939,6,633,5841,1728,3,5048,526,1455,934,1712,101,2,4,1,321,840,2,3249,1008,670,27,3,10684,963,4,2,2,2,1330,5031,2990,2714,690,3197,1365,573,150,5,2334,1279,408,183,4,1378,136,1279,1792,821,519,2126,714,407,573,380,754,2,4,531,812,373,315,127,878,796,5,1722,4,6,426,8,93,464,857,452,10,1,720,562,383,1499,114,1025,480,2,9,1,302,2393,505,924,817,2,669,7,130,5,747,161,169,227,604,708,157,115,
(... more)

```
