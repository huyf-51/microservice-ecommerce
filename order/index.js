const express = require('express');
const grpcClient = require('./grpc/client')

const app = express();
require('dotenv').config();
const configs = require('./configs/db')
const port = process.env.PORT || 3004;

configs.connectDB()

app.get('/order', (req ,res) => {
    grpcClient.CreateOrderPayment({orderId: 111}, (err, res) => {
        if (err) {
            console.log(err);
        } else {
            console.log(res);
        }
    })
    res.json('order service')
})
app.listen(port, () => {
    console.log(`order service listen on port ${port}`);
});
