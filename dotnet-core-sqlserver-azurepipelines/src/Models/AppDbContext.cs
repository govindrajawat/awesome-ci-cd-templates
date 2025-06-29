using Microsoft.EntityFrameworkCore;

namespace dotnet_core_sqlserver_azurepipelines.Models;

public class AppDbContext : DbContext
{
    public AppDbContext(DbContextOptions<AppDbContext> options) : base(options) { }
    
    public DbSet<Message> Messages => Set<Message>();
}