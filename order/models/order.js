const { DataTypes, Model } = require('sequelize');
const { sequelize } = require('../configs/db')
class User extends Model {}

User.init(
  {
    // Model attributes are defined here
    orderId: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    
  },
  {
    // Other model options go here
    sequelize, // We need to pass the connection instance
    modelName: 'Order', // We need to choose the model name
  },
);