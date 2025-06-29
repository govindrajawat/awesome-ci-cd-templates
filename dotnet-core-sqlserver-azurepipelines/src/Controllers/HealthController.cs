using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using dotnet_core_sqlserver_azurepipelines.Models;

namespace dotnet_core_sqlserver_azurepipelines.Controllers;

[ApiController]
[Route("[controller]")]
public class HealthController : ControllerBase
{
    private readonly AppDbContext _context;

    public HealthController(AppDbContext context)
    {
        _context = context;
    }

    [HttpGet]
    public async Task<IActionResult> Get()
    {
        try
        {
            // Check database connection
            var canConnect = await _context.Database.CanConnectAsync();
            
            var health = new
            {
                Status = canConnect ? "Healthy" : "Unhealthy",
                Timestamp = DateTime.UtcNow,
                Database = canConnect ? "Connected" : "Disconnected",
                Uptime = Environment.TickCount64
            };

            return canConnect ? Ok(health) : StatusCode(503, health);
        }
        catch (Exception ex)
        {
            var health = new
            {
                Status = "Unhealthy",
                Timestamp = DateTime.UtcNow,
                Error = ex.Message,
                Uptime = Environment.TickCount64
            };

            return StatusCode(503, health);
        }
    }
} 