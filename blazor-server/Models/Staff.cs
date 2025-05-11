using Projects.Data;

namespace Projects.Models
{
  public class Staff : BaseEntity
  {
    public int StaffId { get; set; }
    public required string Name { get; set; }
    public string? Position { get; set; } = string.Empty;
    public required DateTime HireDate { get; set; } = DateTime.UtcNow;
    public DateTime? TerminationDate { get; set; }
    public required StaffStatus Status { get; set; } = StaffStatus.ACTIVE;
    public List<Assignment>? Assignments { get; set; } = [];
    public string? UserId { get; set; }
    public ApplicationUser? User { get; set; }
  }

  public enum StaffStatus
  {
    ACTIVE,
    INACTIVE,
    TERMINATED
  }
}