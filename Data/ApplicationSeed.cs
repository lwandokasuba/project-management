using Microsoft.EntityFrameworkCore;
using Projects.Models;


namespace Projects.Data;

public class SeedData
{
    public static void Initialize(IServiceProvider serviceProvider)
    {
        using var context = new ApplicationDbContext(
            serviceProvider.GetRequiredService<
                DbContextOptions<ApplicationDbContext>>());

        if (context == null || context.Projects == null)
        {
            throw new NullReferenceException(
                "Null ApplicationDbContext or Projects DbSet");
        }

        if (context.Projects.Any())
        {
            return;
        }

        context.Projects.AddRange(
            new Project
            {
                Name = "Project 1",
                Description = "Description for Project 1",
                StartDate = DateTime.UtcNow,
                Status = ProjectStatus.CREATED
            },
            new Project
            {
                Name = "Project 2",
                Description = "Description for Project 2",
                StartDate = DateTime.UtcNow,
                Status = ProjectStatus.IN_PROGRESS
            },
            new Project
            {
                Name = "Project 3",
                Description = "Description for Project 3",
                StartDate = DateTime.UtcNow,
                Status = ProjectStatus.COMPLETED
            }
        );

        context.SaveChanges();
    }
}