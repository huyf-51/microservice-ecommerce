const { Sequelize } = require('sequelize');

const sequelize = new Sequelize(process.env.DB_NAME, process.env.DB_USER, process.env.DB_PASS, {
    host: process.env.DB_HOST,
    dialect: 'postgres'
});

const connectDB = async () => {
    try {
        await sequelize.authenticate()
    } catch (error) {
        console.log(error);
    }
} 

module.exports = { connectDB, sequelize }