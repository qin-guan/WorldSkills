using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class UserType
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<User> Users { get; } = new List<User>();
}
