namespace Projects.Models
{
  public class Expense : BaseEntity
  {
    public int ExpenseId { get; set; }
    public required int AssignmentId { get; set; }
    public Assignment? Assignment { get; set; }
    public required decimal Amount { get; set; }
    public required DateTime Date { get; set; } = DateTime.UtcNow;
    public string? Description { get; set; } = string.Empty;
  }
}