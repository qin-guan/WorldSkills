using System;
using System.Collections.Generic;

namespace Session1.Models;

public partial class DimDate
{
    public long Id { get; set; }

    public DateTime Date { get; set; }

    public int Year { get; set; }

    public int Quarter { get; set; }

    public int Month { get; set; }

    public string MonthName { get; set; } = null!;

    public int DayOfMonth { get; set; }

    public int DayOfWeek { get; set; }

    public string DayName { get; set; } = null!;

    public bool IsHoliday { get; set; }
}
