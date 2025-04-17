namespace Projects.Models
{
  public class Project : BaseEntity
  {
    public int ProjectId { get; set; }
    public required string Name { get; set; }
    public string? Description { get; set; } = string.Empty;
    public required DateTime StartDate { get; set; }
    public DateTime? EndDate { get; set; }
    public required ProjectStatus Status { get; set; } = ProjectStatus.CREATED;
  }

  public enum ProjectStatus
  {
    CREATED,
    IN_PROGRESS,
    COMPLETED,
    SUSPENDED
  }
}