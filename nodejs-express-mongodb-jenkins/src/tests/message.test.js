const request = require('supertest');
const app = require('../app');
const Message = require('../models/message.model');

describe('Message API', () => {
  beforeAll(async () => {
    await Message.deleteMany({});
  });

  afterEach(async () => {
    await Message.deleteMany({});
  });

  describe('GET /api/messages', () => {
    it('should return the default message if no messages exist', async () => {
      const res = await request(app).get('/api/messages');
      expect(res.statusCode).toEqual(200);
      expect(res.body[0].content).toBe('Hello, World!');
    });

    it('should return all messages', async () => {
      await Message.create({ content: 'Test message 1' });
      await Message.create({ content: 'Test message 2' });
      
      const res = await request(app).get('/api/messages');
      expect(res.statusCode).toEqual(200);
      expect(res.body.length).toBe(2);
    });
  });

  describe('POST /api/messages', () => {
    it('should create a new message', async () => {
      const res = await request(app)
        .post('/api/messages')
        .send({ content: 'New message' });
      
      expect(res.statusCode).toEqual(201);
      expect(res.body.content).toBe('New message');
    });
  });
});