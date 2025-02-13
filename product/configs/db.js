const { Sequelize } = require('sequelize');
const fs = require('fs-extra')

const sequelize = new Sequelize(process.env.DB_NAME, process.env.DB_USER, fs.readFileSync(process.env.DB_PASS_FILE, "utf-8"), {
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