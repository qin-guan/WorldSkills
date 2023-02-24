using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class ItemAmenity
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public long ItemId { get; set; }

    public long AmenityId { get; set; }

    public virtual Amenity Amenity { get; set; } = null!;

    public virtual Item Item { get; set; } = null!;
}
