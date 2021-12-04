# go-jwt-auth-api

A simple API using Go Fiber to provide JWT authentication

```
curl -X POST http://localhost/login \
   -H "Content-Type: application/json" \
   -d '{"user":"rodolfo","pass":"bandeira"}' \
```

```
curl --location --request GET 'http://localhost:3000/restricted' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSw\
iZGV0YWlscyI6IkFueSBkZXRhaWwgZm9yIHRoZSB1c2VyIGNhbiBiZSBhZGRlZCBoZXJlLiA6KSIsImV4cCI6MTYz\
ODY1MzA0MCwibmFtZSI6IlJvZG9sZm8gQmFuZGVpcmEifQ.df-CWh_coFlMaShQCSv3WdscpL_Q_EEUhGJIvfsQ9Wd\
JHf9NzWwj6OmGTlHWKne-UXGcHHmKxIy8arzGAH5eZldiYTd_8STsFxXldXpFWWtsQYheOA1J7s4OtzN2LRiTNuz\
si5jf3ZByrrTUtZJEfuxzZQ24UAF22EZ5DE8P7FjAjGP1NxD44LXSAqrhZkPRESjze77q-y2HvvlUYD8K3pkgvvt\
1we6bx3dvLmOpHy_aXv76bEXX1whQ5JVDIJwSHMtmjxha2G118afZxH89Fy0xI182nfvPpKxIh6Ib-H2ZT0gcXxF\
ShimobhGoHyxkScqYUdrsqBeGI1chyEI_Qg'
```
