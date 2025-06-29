using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using dotnet_core_sqlserver_azurepipelines.Controllers;
using dotnet_core_sqlserver_azurepipelines.Models;
using Xunit;

namespace dotnet_core_sqlserver_azurepipelines.Tests.Controllers;

public class HealthControllerTests
{
    [Fact]
    public async Task Get_WhenDatabaseIsAvailable_ReturnsHealthyStatus()
    {
        // Arrange
        var options = new DbContextOptionsBuilder<AppDbContext>()
            .UseInMemoryDatabase(databaseName: "HealthTestDb")
            .Options;
        
        using var context = new AppDbContext(options);
        var controller = new HealthController(context);
        
        // Act
        var result = await controller.Get();
        
        // Assert
        var okResult = Assert.IsType<OkObjectResult>(result);
        var health = okResult.Value as dynamic;
        Assert.Equal("Healthy", health.Status.ToString());
        Assert.Equal("Connected", health.Database.ToString());
    }
    
    [Fact]
    public async Task Get_WhenDatabaseIsUnavailable_ReturnsUnhealthyStatus()
    {
        // Arrange
        var options = new DbContextOptionsBuilder<AppDbContext>()
            .UseInMemoryDatabase(databaseName: "HealthTestDb2")
            .Options;
        
        using var context = new AppDbContext(options);
        await context.Database.EnsureDeletedAsync(); // This will make the database unavailable
        
        var controller = new HealthController(context);
        
        // Act
        var result = await controller.Get();
        
        // Assert
        var statusResult = Assert.IsType<ObjectResult>(result);
        Assert.Equal(503, statusResult.StatusCode);
    }
} 