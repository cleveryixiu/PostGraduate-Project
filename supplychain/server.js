var path = require('path')
var express = require('express')
var layout = require('express-layout')
var routes = require('./controller/routes')
var app = express()
var bodyParser = require('body-parser')
var flash = require('express-flash')
var cookieParser = require('cookie-parser')
var session = require('express-session')
var FileStore = require('session-file-store')(session)
var uuid = require('uuid/v4')

app.use(session({
  genid: (req) => {
    console.log('Inside the session middleware')
    console.log(req.sessionID)
    return uuid()
  },
  secret: 'secret',
  key: 'key',
  resave: false,
  saveUninitialized: false,
  cookie: { maxAge: 60 * 1000 * 60 * 5}
}))

app.use(function(req, res, next) {
  res.locals.session = req.session;
  console.log(req.session)
  console.log(res.locals.session )
  next();

  console.log("------------------------1-------------------------")
})

app.get('/login', function(req, res, next) {
  console.log("------------------------2-------------------------")
  console.log(req.sessionID)
  req.session.username = req.query.name;
  req.session.userId = req.query.id;
  res.redirect('/index');
})

const middlewares = [
  express.static(path.join(__dirname, 'public')),
  bodyParser.urlencoded({ extended: true }),
  cookieParser(),
  flash()
]
app.use(middlewares)
app.use(express.static(path.join(__dirname, 'views')));
app.set('view engine', 'ejs');
require('./controller/routes.js')(app);

app.use((req, res, next) => {
  res.status(404).send("Sorry can not find that!")
  console.log("------------------------3-------------------------")
})

app.use((err, req, res, next) => {
  console.log(err.stack)
  res.status(500).send('Something broke!')

  console.log("------------------------4-------------------------")
})



var port = process.env.PORT || 8080;

app.listen(port,function(){
  console.log("Live on port: " + port);
});
