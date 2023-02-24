using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class CancellationPolicy
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public string Name { get; set; } = null!;

    public decimal Commission { get; set; }

    public virtual ICollection<CancellationRefundFee> CancellationRefundFees { get; } = new List<CancellationRefundFee>();

    public virtual ICollection<ItemPrice> ItemPrices { get; } = new List<ItemPrice>();
}
