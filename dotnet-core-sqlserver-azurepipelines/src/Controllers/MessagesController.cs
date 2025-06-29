using Microsoft.AspNetCore.Mvc;
using dotnet_core_sqlserver_azurepipelines.Services;
using dotnet_core_sqlserver_azurepipelines.Models;

namespace dotnet_core_sqlserver_azurepipelines.Controllers;

[ApiController]
[Route("api/[controller]")]
public class MessagesController : ControllerBase
{
    private readonly MessageService _messageService;
    
    public MessagesController(MessageService messageService)
    {
        _messageService = messageService;
    }
    
    [HttpGet]
    public IActionResult Get()
    {
        return Ok(_messageService.GetAllMessages());
    }
    
    [HttpPost]
    public IActionResult Post(Message message)
    {
        return Ok(_messageService.CreateMessage(message));
    }
}