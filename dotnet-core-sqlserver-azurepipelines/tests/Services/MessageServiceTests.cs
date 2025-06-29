using Moq;
using dotnet_core_sqlserver_azurepipelines.Services;
using dotnet_core_sqlserver_azurepipelines.Models;
using Microsoft.EntityFrameworkCore;
using Xunit;

namespace dotnet_core_sqlserver_azurepipelines.Tests.Services;

public class MessageServiceTests
{
    [Fact]
    public void GetAllMessages_ReturnsAllMessages()
    {
        // Arrange
        var options = new DbContextOptionsBuilder<AppDbContext>()
            .UseInMemoryDatabase(databaseName: "TestDb")
            .Options;
        
        using var context = new AppDbContext(options);
        context.Messages.Add(new Message { Content = "Test" });
        context.SaveChanges();
        
        var service = new MessageService(context);
        
        // Act
        var messages = service.GetAllMessages();
        
        // Assert
        Assert.Single(messages);
        Assert.Equal("Test", messages.First().Content);
    }
    
    [Fact]
    public void CreateMessage_AddsMessageToDatabase()
    {
        // Arrange
        var options = new DbContextOptionsBuilder<AppDbContext>()
            .UseInMemoryDatabase(databaseName: "TestDb")
            .Options;
        
        using var context = new AppDbContext(options);
        var service = new MessageService(context);
        var message = new Message { Content = "New" };
        
        // Act
        var created = service.CreateMessage(message);
        
        // Assert
        Assert.Equal(1, context.Messages.Count());
        Assert.Equal("New", context.Messages.First().Content);
    }
}