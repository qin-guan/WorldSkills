using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class Attraction
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public long AreaId { get; set; }

    public string Name { get; set; } = null!;

    public string Address { get; set; } = null!;

    public virtual Area Area { get; set; } = null!;

    public virtual ICollection<ItemAttraction> ItemAttractions { get; } = new List<ItemAttraction>();
}
