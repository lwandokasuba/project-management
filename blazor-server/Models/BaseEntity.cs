// using System.ComponentModel;
// using System.ComponentModel.DataAnnotations.Schema;

namespace Projects.Models
{
    public abstract class BaseEntity
    {
        // [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
        // [DatabaseGenerated(DatabaseGeneratedOption.Computed)]
        public DateTime UpdatedAt { get; set; } = DateTime.UtcNow;
    }
}