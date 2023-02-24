using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class Amenity
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public string Name { get; set; } = null!;

    public string? IconName { get; set; }

    public virtual ICollection<ItemAmenity> ItemAmenities { get; } = new List<ItemAmenity>();
}
