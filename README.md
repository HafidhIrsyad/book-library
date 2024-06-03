<!-- PROJECT LOGO -->
<br />
<p align="center">

<h3 align="center">Book Library Management</h3>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
	- [Feature](#feature)
	- [Built With](#built-with)
	- [Usage](#usage)
- [Getting Started](#getting-started)
	- [Prerequisites](#prerequisites)
	- [Installation](#installation)


<!-- ABOUT THE PROJECT -->
## About The Project

### Feature

* CRUD Book
* CRUD Auhtor
* CRUD Category

### Built With

* [Go as Programming Language](https://golang.org/)
* [PostgreSQL as Database](https://www.postgresql.org/)
* [Chi as HTTP Router](https://go-chi.io/#/README)
* [Viper for reading .env file](https://github.com/spf13/viper)
* [GORM as ORM](https://gorm.io/)
* [Zerolog as logging mechanism](https://github.com/rs/zerolog?tab=readme-ov-file)

### Usage
* [Postman Collections](https://drive.google.com/file/d/1kaatYll4cOrl6Y5FcbKalam9nxln8_cv/view?usp=sharing)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running, follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Install Golang, PostgreSQL, and Postman for testing
* create an `.env` file or copy from file `.env.example`

```bash
APP_PORT=8080
DRIVER_NAME=postgres
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_MAX_OPEN_CONNECTION=25
DB_MAX_IDLE_CONNECTION=10
DB_CONNECTION_MAX_LIFE_TIME=300
```

### Installation

1. Clone the repo
```sh
git clone git@github.com:HafidhIrsyad/book-library.git
```
2. Install module with get
```sh
go mod tidy
```
3. Run
```sh
source .env
go run main.go
```
4. Access via url (postman)
```JS
http://localhost:8080/api/v1/
```
