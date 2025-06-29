package com.example.demo.service;

import com.example.demo.model.Message;
import com.example.demo.model.MessageRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import java.util.Arrays;
import java.util.List;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
public class MessageServiceTest {
    @Mock
    private MessageRepository messageRepository;
    
    @InjectMocks
    private MessageService messageService;
    
    @Test
    public void getAllMessages_ShouldReturnMessages() {
        Message message = new Message("Test message");
        when(messageRepository.findAll()).thenReturn(Arrays.asList(message));
        
        List<Message> messages = messageService.getAllMessages();
        
        assertEquals(1, messages.size());
        assertEquals("Test message", messages.get(0).getContent());
    }
    
    @Test
    public void createMessage_ShouldSaveMessage() {
        Message message = new Message("New message");
        when(messageRepository.save(message)).thenReturn(message);
        
        Message created = messageService.createMessage(message);
        
        assertEquals("New message", created.getContent());
        verify(messageRepository).save(message);
    }
}