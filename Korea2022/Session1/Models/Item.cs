using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class Item
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public long UserId { get; set; }

    public long ItemTypeId { get; set; }

    public long AreaId { get; set; }

    public string Title { get; set; } = null!;

    public int Capacity { get; set; }

    public int NumberOfBeds { get; set; }

    public int NumberOfBedrooms { get; set; }

    public int NumberOfBathrooms { get; set; }

    public string ExactAddress { get; set; } = null!;

    public string ApproximateAddress { get; set; } = null!;

    public string Description { get; set; } = null!;

    public string HostRules { get; set; } = null!;

    public int MinimumNights { get; set; }

    public int MaximumNights { get; set; }

    public virtual Area Area { get; set; } = null!;

    public virtual ICollection<ItemAmenity> ItemAmenities { get; } = new List<ItemAmenity>();

    public virtual ICollection<ItemAttraction> ItemAttractions { get; } = new List<ItemAttraction>();

    public virtual ICollection<ItemPicture> ItemPictures { get; } = new List<ItemPicture>();

    public virtual ICollection<ItemPrice> ItemPrices { get; } = new List<ItemPrice>();

    public virtual ItemType ItemType { get; set; } = null!;

    public virtual User User { get; set; } = null!;
}
