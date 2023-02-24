using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class ItemPicture
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public long ItemId { get; set; }

    public string PictureFileName { get; set; } = null!;

    public virtual Item Item { get; set; } = null!;
}
