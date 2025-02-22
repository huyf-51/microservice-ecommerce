const express = require('express');

const app = express();
const configs = require('./configs/db')
require('dotenv').config();
const cookieParser = require('cookie-parser')
const port = process.env.PORT || 3005;
// const apm = require('elastic-apm-node').start({
//     serviceName: 'my-service-name',
  
//     secretToken: '',
  
//     serverUrl: 'http://localhost:8200',
  
//     environment: 'my-environment'
//   })
app.use(cookieParser())

configs.connectDB()

app.get('/product', (req, res) => {
    console.log('cookie>>>', req.cookies);
    console.log('hello iam product service');
    res.json('hello huy')
})

app.listen(port, () => {
    console.log(`product service listen on port ${port}`);
});
