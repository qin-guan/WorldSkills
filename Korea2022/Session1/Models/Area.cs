using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class Area
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<Attraction> Attractions { get; } = new List<Attraction>();

    public virtual ICollection<Item> Items { get; } = new List<Item>();
}
