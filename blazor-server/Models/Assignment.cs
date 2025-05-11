namespace Projects.Models
{
  public class Assignment : BaseEntity
  {
    public int AssignmentId { get; set; }
    public required string Title { get; set; } = string.Empty;
    public string? Description { get; set; } = string.Empty;
    public required int ProjectId { get; set; }
    public Project? Project { get; set; }
    public required int StaffId { get; set; }
    public Staff? Staff { get; set; }
    public List<Expense>? Expenses { get; set; } = [];
    public List<Gain>? Gains { get; set; } = [];
    public required DateTime StartDate { get; set; } = DateTime.UtcNow;
    public DateTime? EndDate { get; set; }
    public required AssignmentStatus Status { get; set; } = AssignmentStatus.ACTIVE;
  }

  public enum AssignmentStatus
  {
    ACTIVE,
    INACTIVE,
    COMPLETED
  }
}