# teslatoken

`teslatoken` is a tiny Go program for the command line which, given your
Tesla account credentials, requests OAuth tokens from the API and prints them.

## How to get it

```
$ go get github.com/thcyron/teslatoken
```

## How to use it

```
$ teslatoken tesla@example.com
Password:
Access token: qts-d21182a6526485378d5c0cb0b0076329a771d5c1896f22694796fea6ef898696
Refresh token: 07ecd051f28cfdd0c47d3b7db904d0300b31317526ff8b2332b5470fd676382a
Expiry: Sun, 01 Nov 2020 18:27:47 CET
```
