using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using Projects.Models;

namespace Projects.Data;

public class ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : IdentityDbContext<ApplicationUser>(options)
{
  public DbSet<Project> Projects { get; set; } = default!;
  public DbSet<Staff> Staffs { get; set; } = default!;
  public DbSet<Assignment> Assignments { get; set; } = default!;
  public DbSet<Expense> Expenses { get; set; } = default!;
  public DbSet<Gain> Gains { get; set; } = default!;

  protected override void OnModelCreating(ModelBuilder builder)
  {
    base.OnModelCreating(builder);

    // Projects
    builder.Entity<Project>().ToTable("Projects")
      .HasKey(p => p.ProjectId);

    // Staff
    builder.Entity<Staff>().ToTable("Staffs")
      .HasKey(s => s.StaffId);

    builder.Entity<Staff>()
      .HasOne(s => s.User)
      .WithOne(u => u.Staff)
      .HasForeignKey<Staff>(s => s.UserId);

    // Assignments
    builder.Entity<Assignment>().ToTable("Assignments")
      .HasKey(a => a.AssignmentId);

    builder.Entity<Assignment>()
      .HasOne(a => a.Project)
      .WithMany(p => p.Assignments)
      .HasForeignKey(a => a.ProjectId);

    builder.Entity<Assignment>()
      .HasOne(a => a.Staff)
      .WithMany(s => s.Assignments)
      .HasForeignKey(a => a.StaffId);

    builder.Entity<Assignment>()
      .Property(a => a.Status)
      .HasDefaultValue(AssignmentStatus.ACTIVE);

    // Expenses
    builder.Entity<Expense>().ToTable("Expenses")
      .HasKey(e => e.ExpenseId);

    builder.Entity<Expense>()
      .HasOne(e => e.Assignment)
      .WithMany(a => a.Expenses)
      .HasForeignKey(e => e.AssignmentId);

    builder.Entity<Expense>()
      .Property(e => e.Date)
      .HasDefaultValueSql("CURRENT_TIMESTAMP");

    builder.Entity<Expense>()
      .Property(e => e.Amount)
      .HasColumnType("decimal(18,2)");

    // Gains
    builder.Entity<Gain>().ToTable("Gains")
      .HasKey(g => g.GainId);

    builder.Entity<Gain>()
      .Property(g => g.Date)
      .HasDefaultValueSql("CURRENT_TIMESTAMP");
    
    builder.Entity<Gain>()
      .Property(g => g.Amount)
      .HasColumnType("decimal(18,2)");

    builder.Entity<Gain>()
      .HasOne(g => g.Assignment)
      .WithMany(p => p.Gains)
      .HasForeignKey(g => g.AssignmentId);
  }
}
