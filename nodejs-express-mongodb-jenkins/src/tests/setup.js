const mongoose = require('mongoose');

// Configure test database connection
beforeAll(async () => {
  const mongoUri = process.env.MONGODB_URI || 'mongodb://root:example@localhost:27018/expressdb_test';
  await mongoose.connect(mongoUri, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  });
});

afterAll(async () => {
  await mongoose.connection.close();
});

// Global test timeout
jest.setTimeout(30000); 