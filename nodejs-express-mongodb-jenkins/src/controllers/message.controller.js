const Message = require('../models/message.model');

exports.getMessages = async (req, res) => {
  try {
    const messages = await Message.find();
    if (messages.length === 0) {
      const defaultMessage = await Message.create({ content: 'Hello, World!' });
      return res.json([defaultMessage]);
    }
    res.json(messages);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
};

exports.createMessage = async (req, res) => {
  try {
    const message = new Message({
      content: req.body.content
    });
    const newMessage = await message.save();
    res.status(201).json(newMessage);
  } catch (err) {
    res.status(400).json({ message: err.message });
  }
};