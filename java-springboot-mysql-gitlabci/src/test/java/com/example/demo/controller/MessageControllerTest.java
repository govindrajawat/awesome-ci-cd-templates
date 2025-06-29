package com.example.demo.controller;

import com.example.demo.model.Message;
import com.example.demo.service.MessageService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import java.util.Arrays;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest
@AutoConfigureMockMvc
public class MessageControllerTest {
    @Autowired
    private MockMvc mockMvc;
    
    @MockBean
    private MessageService messageService;
    
    @Test
    public void getAllMessages_ShouldReturnMessages() throws Exception {
        Message message = new Message("Test message");
        when(messageService.getAllMessages()).thenReturn(Arrays.asList(message));
        
        mockMvc.perform(get("/api/messages"))
               .andExpect(status().isOk())
               .andExpect(jsonPath("$[0].content").value("Test message"));
    }
    
    @Test
    public void createMessage_ShouldReturnCreatedMessage() throws Exception {
        Message message = new Message("New message");
        when(messageService.createMessage(message)).thenReturn(message);
        
        mockMvc.perform(post("/api/messages")
               .contentType("application/json")
               .content("{\"content\":\"New message\"}"))
               .andExpect(status().isOk())
               .andExpect(jsonPath("$.content").value("New message"));
    }
}