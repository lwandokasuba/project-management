using Microsoft.AspNetCore.Identity;
using Projects.Models;

namespace Projects.Data;

// Add profile data for application users by adding properties to the ApplicationUser class
public class ApplicationUser : IdentityUser
{
    public Staff? Staff { get; set; } = null!;
    
}

