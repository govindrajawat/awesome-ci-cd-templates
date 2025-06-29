using dotnet_core_sqlserver_azurepipelines.Models;

namespace dotnet_core_sqlserver_azurepipelines.Services;

public class MessageService
{
    private readonly AppDbContext _context;
    
    public MessageService(AppDbContext context)
    {
        _context = context;
    }
    
    public IEnumerable<Message> GetAllMessages()
    {
        return _context.Messages.ToList();
    }
    
    public Message CreateMessage(Message message)
    {
        _context.Messages.Add(message);
        _context.SaveChanges();
        return message;
    }
}