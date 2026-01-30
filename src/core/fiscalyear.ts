import type { TextualYearMonth, MonthDay, UnixTimeRange } from './datetime.ts';

export class FiscalYearStart {
    public static readonly JanuaryFirstDay = new FiscalYearStart(1, 1);
    public static readonly Default = FiscalYearStart.JanuaryFirstDay;

    private static readonly MONTH_MAX_DAYS: number[] = [
        31, 
        28, 
        31, 
        30, 
        31, 
        30, 
        31, 
        31, 
        30, 
        31, 
        30, 
        31 
    ];

    public readonly month: number; 
    public readonly day: number;
    public readonly value: number;

    private constructor(month: number, day: number) {
        this.month = month;
        this.day = day;
        this.value = (month << 8) | day;
    }

    public static of(month: number, day: number): FiscalYearStart | undefined {
        if (!FiscalYearStart.isValidFiscalYearMonthDay(month, day)) {
            return undefined;
        }

        return new FiscalYearStart(month, day);
    }

    
    public static valueOf(value: number): FiscalYearStart | undefined {
        if (value < 0x0101 || value > 0x0C1F) {
            return undefined;
        }

        const month = (value >> 8) & 0xFF;  
        const day = value & 0xFF;           

        return FiscalYearStart.of(month, day);
    }

    
    public static parse(monthDay: string): FiscalYearStart | undefined {
        if (!monthDay || !monthDay.includes('-')) {
            return undefined;
        }

        const parts = monthDay.split('-');

        if (parts.length !== 2) {
            return undefined;
        }

        const month = parseInt(parts[0] as string, 10);
        const day = parseInt(parts[1] as string, 10);

        return FiscalYearStart.of(month, day);
    }

    public toMonthDashDayString(): TextualYearMonth {
        return `${this.month.toString().padStart(2, '0')}-${this.day.toString().padStart(2, '0')}` as TextualYearMonth;
    }

    public toMonthDay(): MonthDay {
        return {
            month: this.month,
            day: this.day
        };
    }

    private static isValidFiscalYearMonthDay(month: number, day: number): boolean {
        return 1 <= month && month <= 12 && 1 <= day && day <= (FiscalYearStart.MONTH_MAX_DAYS[month - 1] as number);
    }
}

export class FiscalYearUnixTime implements UnixTimeRange {
    public readonly year: number;
    public readonly minUnixTime: number;
    public readonly maxUnixTime: number;

    private constructor(fiscalYear: number, minUnixTime: number, maxUnixTime: number) {
        this.year = fiscalYear;
        this.minUnixTime = minUnixTime;
        this.maxUnixTime = maxUnixTime;
    }

    public static of(fiscalYear: number, minUnixTime: number, maxUnixTime: number): FiscalYearUnixTime {
        return new FiscalYearUnixTime(fiscalYear, minUnixTime, maxUnixTime);
    }
}

export const LANGUAGE_DEFAULT_FISCAL_YEAR_FORMAT_VALUE: number = 0;

export class FiscalYearFormat {
    private static readonly allInstances: FiscalYearFormat[] = [];
    private static readonly allInstancesByType: Record<number, FiscalYearFormat> = {};
    private static readonly allInstancesByTypeName: Record<string, FiscalYearFormat> = {};

    public static readonly StartYYYY_EndYYYY = new FiscalYearFormat(1, 'StartYYYY_EndYYYY');
    public static readonly StartYYYY_EndYY = new FiscalYearFormat(2, 'StartYYYY_EndYY');
    public static readonly StartYY_EndYY = new FiscalYearFormat(3, 'StartYY_EndYY');
    public static readonly EndYYYY = new FiscalYearFormat(4, 'EndYYYY');
    public static readonly EndYY = new FiscalYearFormat(5, 'EndYY');

    public static readonly Default = FiscalYearFormat.EndYYYY;

    public readonly type: number;
    public readonly typeName: string;

    private constructor(type: number, typeName: string) {
        this.type = type;
        this.typeName = typeName;

        FiscalYearFormat.allInstances.push(this);
        FiscalYearFormat.allInstancesByType[type] = this;
        FiscalYearFormat.allInstancesByTypeName[typeName] = this;
    }

    public static values(): FiscalYearFormat[] {
        return FiscalYearFormat.allInstances;
    }

    public static valueOf(type: number): FiscalYearFormat | undefined {
        return FiscalYearFormat.allInstancesByType[type];
    }

    public static parse(typeName: string): FiscalYearFormat | undefined {
        return FiscalYearFormat.allInstancesByTypeName[typeName];
    }
}
