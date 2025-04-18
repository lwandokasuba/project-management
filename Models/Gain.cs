namespace Projects.Models
{
  public class Gain : BaseEntity
  {
    public int GainId { get; set; }
    public required int AssignmentId { get; set; }
    public Assignment? Assignment { get; set; }
    public required decimal Amount { get; set; }
    public required DateTime Date { get; set; } = DateTime.UtcNow;
    public string? Description { get; set; } = string.Empty;
  }
}