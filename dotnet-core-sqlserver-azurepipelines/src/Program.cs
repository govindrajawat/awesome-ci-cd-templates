using Microsoft.EntityFrameworkCore;
using dotnet_core_sqlserver_azurepipelines.Models;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container
builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DefaultConnection")));

builder.Services.AddScoped<MessageService>();
builder.Services.AddControllers();

// Add Swagger/OpenAPI
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline
if (app.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseAuthorization();
app.MapControllers();

// Initialize database
using (var scope = app.Services.CreateScope())
{
    var dbContext = scope.ServiceProvider.GetRequiredService<AppDbContext>();
    
    try
    {
        // Ensure database is created and migrations are applied
        dbContext.Database.Migrate();
        
        // Seed data if no messages exist
        if (!dbContext.Messages.Any())
        {
            dbContext.Messages.Add(new Message { Content = "Hello, World!" });
            dbContext.SaveChanges();
        }
    }
    catch (Exception ex)
    {
        // Log the error but don't fail the application startup
        Console.WriteLine($"Database initialization error: {ex.Message}");
    }
}

app.Run();