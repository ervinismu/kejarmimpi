# kejarmimpi    
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)]()
[![GitHub release](https://img.shields.io/github/release/phw/peek.svg)](https://github.com/phw/peek/releases)
[![Twitter URL](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)]()
[![Translation Status](https://hosted.weblate.org/widgets/peek/-/svg-badge.svg)](https://hosted.weblate.org/engage/peek/?utm_source=widget)
[![label](https://img.shields.io/github/issues-raw/badges/shields/website.svg)]()  
Restfull api with gin-gonic framework golang language

# Deps
- Install golang
- Set gopath
- Install heroku CLI
- Install Postgresql
- Postman (optional) for testing

# Running in local
- Clone this appp :

      go get github.com/ervinismu/kejarmimpi

- Replace the database connection code with your database connection in the `db.go` file
- Run by : 
    
      go run main.go

- Open localhost:8080

# Running in Heroku
- Download this app :

      go get github.com/ervinismu/kejarmimpi

- Login akun : 

      Heroku login

- Create apps : 

      Heroku create

- Create database postgresql : 

      Heroku addons:create heroku-postgresql:hobby-dev

- Push : 

      Git push heroku master

- Open app : 

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

# Authors
- Reza andriyunantoreza@gmail.com
- Ervin social.ervin@gmail.com
- Yudha yudhaariestra@gmail.com


