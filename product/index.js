const express = require('express');

const app = express();
const configs = require('./configs/db')
require('dotenv').config();
const port = process.env.PORT || 3005;

configs.connectDB()

app.get('/product', (req, res) => {
    res.json('hello huy')
})
 
app.listen(port, () => {
    console.log(`product service listen on port ${port}`);
});
