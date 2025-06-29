using Microsoft.AspNetCore.Mvc;
using Moq;
using dotnet_core_sqlserver_azurepipelines.Controllers;
using dotnet_core_sqlserver_azurepipelines.Services;
using dotnet_core_sqlserver_azurepipelines.Models;
using Xunit;

namespace dotnet_core_sqlserver_azurepipelines.Tests.Controllers;

public class MessagesControllerTests
{
    private readonly Mock<MessageService> _mockService;
    private readonly MessagesController _controller;
    
    public MessagesControllerTests()
    {
        _mockService = new Mock<MessageService>(null);
        _controller = new MessagesController(_mockService.Object);
    }
    
    [Fact]
    public void Get_ReturnsOkResultWithMessages()
    {
        // Arrange
        var messages = new List<Message> { new Message { Id = 1, Content = "Test" } };
        _mockService.Setup(s => s.GetAllMessages()).Returns(messages);
        
        // Act
        var result = _controller.Get();
        
        // Assert
        var okResult = Assert.IsType<OkObjectResult>(result);
        var returnMessages = Assert.IsType<List<Message>>(okResult.Value);
        Assert.Single(returnMessages);
    }
    
    [Fact]
    public void Post_ReturnsOkResultWithCreatedMessage()
    {
        // Arrange
        var message = new Message { Content = "New" };
        _mockService.Setup(s => s.CreateMessage(It.IsAny<Message>())).Returns(message);
        
        // Act
        var result = _controller.Post(message);
        
        // Assert
        var okResult = Assert.IsType<OkObjectResult>(result);
        var returnMessage = Assert.IsType<Message>(okResult.Value);
        Assert.Equal("New", returnMessage.Content);
    }
}