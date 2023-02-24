using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class User
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public long UserTypeId { get; set; }

    public string Username { get; set; } = null!;

    public string Password { get; set; } = null!;

    public string FullName { get; set; } = null!;

    public bool Gender { get; set; }

    public DateTime BirthDate { get; set; }

    public int FamilyCount { get; set; }

    public virtual ICollection<Item> Items { get; } = new List<Item>();

    public virtual UserType UserType { get; set; } = null!;
}
