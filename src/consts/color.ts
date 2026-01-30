import type { ColorValue, ColorStyleValue } from '@/core/color.ts';

const defaultColor: ColorValue = '000000';

export const DEFAULT_ICON_COLOR: ColorValue = defaultColor;
export const DEFAULT_ACCOUNT_COLOR: ColorValue = defaultColor;
export const DEFAULT_CATEGORY_COLOR: ColorValue = defaultColor;

export const DEFAULT_COLOR_STYLE_VARIABLE: ColorStyleValue = 'var(--default-icon-color)';

const allAvailableColors: ColorValue[] = [
    '000000', 
    '8e8e93', 
    'ff3b30', 
    'ff2d55', 
    'ff6b22', 
    'ff9500', 
    'ffcc00', 
    'cddc39', 
    '009688', 
    '4cd964', 
    '5ac8fa', 
    '2196f3', 
    '673ab7', 
    '9c27b0', 
];

export const ALL_ACCOUNT_COLORS: ColorValue[] = allAvailableColors;
export const ALL_CATEGORY_COLORS: ColorValue[] = allAvailableColors;

export const DEFAULT_CHART_COLORS: ColorValue[] = [
    'cc4a66',
    'e3564a',
    'fc892c',
    'ffc349',
    '4dd291',
    '24ceb3',
    '2ab4d0',
    '065786',
    '713670',
    '8e1d51'
];
