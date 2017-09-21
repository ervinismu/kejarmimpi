# kejarmimpi    
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)]()
[![GitHub release](https://img.shields.io/github/release/phw/peek.svg)](https://github.com/phw/peek/releases)
[![Twitter URL](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)]()
[![Translation Status](https://hosted.weblate.org/widgets/peek/-/svg-badge.svg)](https://hosted.weblate.org/engage/peek/?utm_source=widget)
[![label](https://img.shields.io/github/issues-raw/badges/shields/website.svg)]()  
Restfull api with gin-gonic framework golang language

# Deps
- [x] Install golang
- [x] Set gopath
- [x] Install heroku CLI
- [x] Install Postgresql
- [x] Postman (optional) for testing

# Running in local
- [x] Clone this repo
- [x] Replace the database connection code with your database connection in the `db.go` file
- [x] Run by : 
    
      go run server.go

- [x] Open localhost:8080

# Running in Heroku
- [x] Clone this repo
- [x] Login akun : 

      Heroku login

- [x] Create apps : 

      Heroku create

- [x] Create database postgresql : 

      Heroku addons:create heroku-postgresql:hobby-dev

- [x] Push : 

      Git push heroku master

- [x] Open app : 

      Heroku open

# Testing in Url
Testing your url with Postman or curl

| Name                | Url                   | Method   |
| --------------------|:---------------------:|:--------:|
| Register            | `<your_url>/register` |   **POST**   |
| Login               | `<your_url>/login`    |   **POST**   |
| Logout              | `<your_url>/logout`   |   **GET**    |
| Get All Post        | `<your_url>/post`     |   **GET**    |
| Create Post         | `<your_url>/post`     |   **POST**   |
| Delete Post         | `<your_url>/post`     |   **DELETE** |
| View Profile Data   | `<your_url>/profile`  |   **GET**    |
| Insert Profile Data | `<your_url>/profile`  |   **POST**  |

# License
MIT

# Authors
* Reza andriyunantoreza@gmail.com
* Ervin social.ervin@gmail.com
* Yudha yudhaariestra@gmail.com


