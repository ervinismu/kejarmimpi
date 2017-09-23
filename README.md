<<<<<<< HEAD
# km-gingonic  
Restfull api with gin-gonic framework golang language

# Deps
* Install golang
* Set gopath
* Install heroku CLI
* Install Postgresql
* Postman (optional) for testing

# Running in local
* Clone this repo
* Replace the database connection code with your database connection in the `db.go` file
* Run by : `go run server.go`
* Open localhost:8080

# Running in Heroku
* Clone this repo
* Login akun : `Heroku login`
* Create apps : `Heroku create`
* Create database postgresql : `Heroku addons:create heroku-postgresql:hobby-dev`
* Push : `Git push heroku master`
* Open app : `Heroku open`

# Testing in Url
* Register  
`<your_url>/register` with method POST in Postman
* Login  
`<your_url>/login` with method POST in Postman
* Logout  
`<your_url>/logout` with method GET in Postman
* Get All Post  
`<your_url>/post` with method GET in Postman
* Create Post  
`<your_url>/post` with method POST in Postman
* Delete Post  
`<your_url>/register` with method DELETE in Postman
* View Profile Data  
`<your_url>/register` with method GET in Postman
* Insert Profile Data  
`<your_url>/register` with method POST in Postman

# License
MIT

# Authors
* Reza andriyunantoreza@gmail.com
* Ervin social.ervin@gmail.com
* Yudha yudhaariestra@gmail.com
=======
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
![Peek recording itself](https://github.com/ervinismu/kejarmimpi/blob/master/asset/local.gif)
- Clone this appp :

      go get github.com/ervinismu/kejarmimpi

- Replace the database connection code with your database connection in the `db.go` file
- Run by : 
    
      go run main.go

- Open localhost:8000

# Running in Heroku
![Peek recording itself](https://github.com/ervinismu/kejarmimpi/blob/master/asset/Peek%202017-09-22%2010-49.gif)
- Download this app :

      go get github.com/ervinismu/kejarmimpi

- Login akun : 

      Heroku login

- Create apps : 

      heroku create -b https://github.com/heroku/heroku-buildpack-go.git

- Remote apps in Heroku :

       heroku git:remote -a <your-apps-name>

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
>>>>>>> 53b1a263fc64ede3f0f1aa69986c94fbb420bf65


