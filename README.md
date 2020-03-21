# line-login

![line login](./images/login-with-qrcode.png)

Trying LINE Login on a web app using [golang](https://golang.org). See [demo site](https://line-login-demo-tw.herokuapp.com/).

## Heroku

Click on the Heroku button to easily deploy your app:

[![Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

Or alternatively you can follow the manual steps:

Clone this repository:

```sh
git clone https://github.com/go-training/line-login.git
```

Set the buildpack for your application

```sh
heroku config:add BUILDPACK_URL=https://github.com/go-training/line-login.git -a {Heroku app name}
```

Add Heroku git remote:

```sh
heroku git:remote -a {Heroku app name}
```

Deploy it!

```sh
git push heroku master
```

View the logs. For more information, see [View logs](https://devcenter.heroku.com/articles/logging#view-logs).

```sh
heroku logs --app {Heroku app name} --tail
```

Line config in heroku dashboard:

![Line Config](./images/line-config.png)
