const express = require('express');
const cors = require('cors')

const app = express();
require('dotenv').config();
const port = process.env.PORT || 3005;

// app.use(cors({origin: [process.env.]}))

app.get('/product', (req, res) => {
    res.json('hello huy')
})
 
app.listen(port, () => {
    console.log(`product service listen on port ${port}`);
});
