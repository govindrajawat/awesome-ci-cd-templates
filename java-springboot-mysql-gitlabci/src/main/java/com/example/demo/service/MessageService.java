package com.example.demo.service;

import com.example.demo.model.Message;
import com.example.demo.model.MessageRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import jakarta.annotation.PostConstruct;
import java.util.List;

@Service
public class MessageService {
    @Autowired
    private MessageRepository messageRepository;
    
    @PostConstruct
    public void init() {
        if (messageRepository.count() == 0) {
            messageRepository.save(new Message("Hello, World!"));
        }
    }
    
    public List<Message> getAllMessages() {
        return messageRepository.findAll();
    }
    
    public Message createMessage(Message message) {
        return messageRepository.save(message);
    }
}