using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class TransactionType
{
    public long Id { get; set; }

    public Guid Guid { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<Transaction> Transactions { get; } = new List<Transaction>();
}
