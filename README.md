<br />
<div align="center">
  <img src="https://github.com/yohansky/BE-Hirejob/assets/69236028/5e71bf99-695c-4663-bdc1-aed6e17534c0" />
  <br />
  <h1>Peworld</h1>
  <a href="https://github.com/yohansky/Fe-Blanja-React">View Demo</a>
    .
    <a href="https://github.com/yohansky/Be-Blanja-fiber">Api Demo</a>
</div>

  ## Table of Contents

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Installation](#installation)
  - [Documentation](#documentation)
  - [Related Project](#related-project)

 ## About The Project

 Peworld Hirejob is a website specifically designed to help recruiters find and recruit quality talent according to their interests and needs. The platform provides various features and tools that make it easier for recruiters to identify, assess and contact potential job candidates. With Peworld Hirejob, recruiters can easily run an efficient and effective recruitment process, enabling them to find the most suitable individuals for available positions.

 ### Built With
 These are the libraries and service used for building this backend API

- [Golang](https://go.dev/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Json Web Token](https://jwt.io/)
- [Gorm.io](https://gorm.io/index.html)

## Installation
1. Clone this repository

```sh
git clone https://github.com/yohansky/BE-Hirejob
```

2. Change directory to BE_Hirejob

```sh
cd BE-Hirejob
```

3. Install all of the required modules

```sh
go mod tidy
```

4. Create PostgreSQL database, query are provided in [query.sql](./query.sql)

5. Create and configure `.env` file in the root directory, example credentials are provided in [.env.example](./.env.example)

```txt
- Please note that this server requires Google Drive API credentials and Gmail service account
- Otherwise API endpoint with image upload and account register won't work properly
```

6. Run this command to run the server

```sh
air
```

- Run this command for debugging and finding errors

```sh
golangci-lint run
```

### Documentation

- [Postman API colletion]()
- [PostgreSQL database query](./query.sql)

API endpoint list are also available as published postman documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/31330000/2sA3QniuZG)

### Related Project
:rocket: [`Backend Hirejob`](https://github.com/yohansky/BE-Hirejob)

:rocket: [`Frontend Hirejob`](https://github.com/yohansky/Fe-HireJob-Next)
