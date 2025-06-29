package com.example.demo.controller;

import com.example.demo.model.Message;
import com.example.demo.service.MessageService;
import org.springframework.web.bind.annotation.*;
import java.util.List;

@RestController
@RequestMapping("/api/messages")
public class MessageController {
    private final MessageService messageService;
    
    public MessageController(MessageService messageService) {
        this.messageService = messageService;
    }
    
    @GetMapping
    public List<Message> getAllMessages() {
        return messageService.getAllMessages();
    }
    
    @PostMapping
    public Message createMessage(@RequestBody Message message) {
        return messageService.createMessage(message);
    }
}