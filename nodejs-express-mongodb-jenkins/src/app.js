const express = require('express');
const mongoose = require('mongoose');
const messageRoutes = require('./routes/message.routes');
const healthRoutes = require('./routes/health.routes');
require('dotenv').config();

const app = express();

// Middleware
app.use(express.json());

// Database connection
mongoose.connect(process.env.MONGODB_URI, {
  useNewUrlParser: true,
  useUnifiedTopology: true
})
.then(() => console.log('Connected to MongoDB'))
.catch(err => console.error('MongoDB connection error:', err));

// Routes
app.use('/api/messages', messageRoutes);
app.use('/', healthRoutes);

module.exports = app;