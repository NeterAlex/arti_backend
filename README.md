# Arti
An article-editing application with the separation of front-end and back-end, developed in `typescript` and `go`.

This is the back-end project, and for front-end, it's **[here](https://github.com/NeterAlex/arti)**.

## 🧬 Introduction
It's a testing project with the purpose of learning web and fullstack development, and is being constantly improved.

In the **front-end**, `react` & `next.js` are used as a base framework, `chakra-ui` and `iconpark` built a practical user interface，
`tanstack-query` & `axios` are tasked with handling the data interaction between the front-end and back-end.

In the **back-end**, `gin` is used as a web framework with `go-jwt` for user authentication, 
`gorm` is used to interact with the `SQLite` database.

## ⭐️ Features
+ Basic article display and edit.
+ Markdown support.
+ Dark mode.
+ Responsive UI for mobile and desktop.
+ User authentication using token.

## 🗂 Run
Download the latest release and write a server configuration, then run `arti_server.exe`.

## 🪤 Configuration
> the config file `app.ini` should be placed in `conf` folder, and the follows below are its format.
```
#define the run-mode of gin (release/debug)
RUN_MODE = debug

[app]
#secret for jwt to generate token
JWT_SECRET = $foo

[server]
HTTP_PORT = 8080
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = sqlite
```
