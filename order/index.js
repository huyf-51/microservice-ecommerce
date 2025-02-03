const express = require('express');

const app = express();
require('dotenv').config();
const port = process.env.PORT || 3004;

app.get('/order', (req ,res) => {
    res.json('order service')
})
app.listen(port, () => {
    console.log(`order service listen on port ${port}`);
});
