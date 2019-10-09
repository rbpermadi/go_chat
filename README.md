# go_chat

## Description

go_chat is simple chat application . It's using **Go** as its programming language and websocket. In this example i am using  [gorilla websocket](https://github.com/gorilla/websocket) lib instead of using native golang.org/x/net.

## Onboarding Guide

### Prequisites

* [**Git**](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* [**Go (1.9.7 or later)**](https://golang.org/doc/install)
* [**Go dep**](https://github.com/golang/dep#installation)

### Setup

Please install/clone the [Prequisites](#prequisites) and make sure it works perfectly on your local machine.

After the [Prequisites](#prequisites) have been installed, please clone **go_chat** project into your local machine.

```
> cd github.com/rbpermadi
> git clone git@github.com:rbpermadi/go_chat.git
```

After **go_chat** already cloned, go to its directory then install some dependencies.

```
> cd go_chat
> make dep
```

Make a copy from [env.sample](env.example) and name it `.env`.

```
> cp env.sample .env
```

> _Important Note : **NEVER EVER** delete [env.sample](env.example) file from your local machine

To kill the server you just need to hold `Ctrl + C`

## Running Service

Run **go_chat** in your local machine

```
> cd go_chat
> make run
```

### Rest Services

* Get all current messages : `http://localhost:7171/messages` (GET)
* Create message : `http://localhost:7171/messages` (POST). sample payload : {"body":"test message"}

### Web Socket

To initiate websocket connection, call `ws:localhost:7171/chat`

### Sample Access

For example how to access, you can call `http://localhost:7171/` on browser after you run the apps.