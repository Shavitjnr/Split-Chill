import type { CurrencyInfo } from '@/core/currency.ts';


export const ALL_CURRENCIES: Record<string, CurrencyInfo> = {
    'AED': { 
        code: 'AED',
        fraction: 2,
        symbol: {
            normal: 'Dh',
            plural: 'Dhs'
        },
        unit: 'Dirham'
    },
    'AFN': { 
        code: 'AFN',
        fraction: 2,
        symbol: {
            normal: 'Af.',
            plural: 'Afs.'
        },
        unit: 'Afghani'
    },
    'ALL': { 
        code: 'ALL',
        fraction: 2,
        symbol: {
            normal: 'L'
        },
        unit: 'Lek'
    },
    'AMD': { 
        code: 'AMD',
        fraction: 2,
        symbol: {
            normal: '֏'
        },
        unit: 'Dram'
    },
    'ANG': { 
        code: 'ANG',
        fraction: 2,
        symbol: {
            normal: 'ƒ'
        },
        unit: 'Guilder'
    },
    'AOA': { 
        code: 'AOA',
        fraction: 2,
        symbol: {
            normal: 'Kz'
        },
        unit: 'Kwanza'
    },
    'ARS': { 
        code: 'ARS',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'AUD': { 
        code: 'AUD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'AWG': { 
        code: 'AWG',
        fraction: 2,
        symbol: {
            normal: 'Afl.'
        },
        unit: 'Florin'
    },
    'AZN': { 
        code: 'AZN',
        fraction: 2,
        symbol: {
            normal: '₼'
        },
        unit: 'Manat'
    },
    'BAM': { 
        code: 'BAM',
        fraction: 2,
        symbol: {
            normal: 'KM'
        },
        unit: 'Mark'
    },
    'BBD': { 
        code: 'BBD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'BDT': { 
        code: 'BDT',
        fraction: 2,
        symbol: {
            normal: '৳'
        },
        unit: 'Taka'
    },
    'BGN': { 
        code: 'BGN',
        fraction: 2,
        symbol: {
            normal: 'лв'
        },
        unit: 'Lev'
    },
    'BHD': { 
        code: 'BHD',
        fraction: 3,
        symbol: {
            normal: 'BD'
        },
        unit: 'Dinar'
    },
    'BIF': { 
        code: 'BIF',
        fraction: 0,
        symbol: {
            normal: 'FBu'
        },
        unit: 'Franc'
    },
    'BMD': { 
        code: 'BMD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'BND': { 
        code: 'BND',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'BOB': { 
        code: 'BOB',
        fraction: 2,
        symbol: {
            normal: 'Bs'
        },
        unit: 'Boliviano'
    },
    'BRL': { 
        code: 'BRL',
        fraction: 2,
        symbol: {
            normal: 'R$'
        },
        unit: 'Real'
    },
    'BSD': { 
        code: 'BSD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'BTN': { 
        code: 'BTN',
        fraction: 2,
        symbol: {
            normal: 'Nu.'
        },
        unit: 'Ngultrum'
    },
    'BWP': { 
        code: 'BWP',
        fraction: 2,
        symbol: {
            normal: 'P'
        },
        unit: 'Pula'
    },
    'BYN': { 
        code: 'BYN',
        fraction: 2,
        symbol: {
            normal: 'Rbl',
            plural: 'Rbls'
        },
        unit: 'Ruble'
    },
    'BZD': { 
        code: 'BZD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'CAD': { 
        code: 'CAD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'CDF': { 
        code: 'CDF',
        fraction: 2,
        symbol: {
            normal: 'FC'
        },
        unit: 'Franc'
    },
    'CHF': { 
        code: 'CHF',
        fraction: 2,
        symbol: {
            normal: 'CHF'
        },
        unit: 'Franc'
    },
    'CLP': { 
        code: 'CLP',
        fraction: 0,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'CNY': { 
        code: 'CNY',
        fraction: 2,
        symbol: {
            normal: '¥'
        },
        unit: 'Yuan'
    },
    'COP': { 
        code: 'COP',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'CRC': { 
        code: 'CRC',
        fraction: 2,
        symbol: {
            normal: '₡'
        },
        unit: 'Colon'
    },
    'CUC': { 
        code: 'CUC',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'CUP': { 
        code: 'CUP',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'CVE': { 
        code: 'CVE',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Escudo'
    },
    'CZK': { 
        code: 'CZK',
        fraction: 2,
        symbol: {
            normal: 'Kč'
        },
        unit: 'Koruna'
    },
    'DJF': { 
        code: 'DJF',
        fraction: 0,
        symbol: {
            normal: 'Fdj'
        },
        unit: 'Franc'
    },
    'DKK': { 
        code: 'DKK',
        fraction: 2,
        symbol: {
            normal: 'kr.'
        },
        unit: 'Krone'
    },
    'DOP': { 
        code: 'DOP',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'DZD': { 
        code: 'DZD',
        fraction: 2,
        symbol: {
            normal: 'DA'
        },
        unit: 'Dinar'
    },
    'EGP': { 
        code: 'EGP',
        fraction: 2,
        symbol: {
            normal: '£'
        },
        unit: 'Pound'
    },
    'ERN': { 
        code: 'ERN',
        fraction: 2,
        symbol: {
            normal: 'Nfk'
        },
        unit: 'Nakfa'
    },
    'ETB': { 
        code: 'ETB',
        fraction: 2,
        symbol: {
            normal: 'Br'
        },
        unit: 'Birr'
    },
    'EUR': { 
        code: 'EUR',
        fraction: 2,
        symbol: {
            normal: '€'
        },
        unit: 'Euro'
    },
    'FJD': { 
        code: 'FJD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'FKP': { 
        code: 'FKP',
        fraction: 2,
        symbol: {
            normal: '£'
        },
        unit: 'Pound'
    },
    'GBP': { 
        code: 'GBP',
        fraction: 2,
        symbol: {
            normal: '£'
        },
        unit: 'Pound'
    },
    'GEL': { 
        code: 'GEL',
        fraction: 2,
        symbol: {
            normal: 'ლ'
        },
        unit: 'Lari'
    },
    'GHS': { 
        code: 'GHS',
        fraction: 2,
        symbol: {
            normal: 'GH₵'
        },
        unit: 'Cedi'
    },
    'GIP': { 
        code: 'GIP',
        fraction: 2,
        symbol: {
            normal: '£'
        },
        unit: 'Pound'
    },
    'GMD': { 
        code: 'GMD',
        fraction: 2,
        symbol: {
            normal: 'D'
        },
        unit: 'Dalasi'
    },
    'GNF': { 
        code: 'GNF',
        fraction: 0,
        symbol: {
            normal: 'FG'
        },
        unit: 'Franc'
    },
    'GTQ': { 
        code: 'GTQ',
        fraction: 2,
        symbol: {
            normal: 'Q'
        },
        unit: 'Quetzal'
    },
    'GYD': { 
        code: 'GYD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'HKD': { 
        code: 'HKD',
        fraction: 2,
        symbol: {
            normal: 'HK$'
        },
        unit: 'Dollar'
    },
    'HNL': { 
        code: 'HNL',
        fraction: 2,
        symbol: {
            normal: 'L'
        },
        unit: 'Lempira'
    },
    'HTG': { 
        code: 'HTG',
        fraction: 2,
        symbol: {
            normal: 'G'
        },
        unit: 'Gourde'
    },
    'HUF': { 
        code: 'HUF',
        fraction: 2,
        symbol: {
            normal: 'Ft'
        },
        unit: 'Forint'
    },
    'IDR': { 
        code: 'IDR',
        fraction: 2,
        symbol: {
            normal: 'Rp'
        },
        unit: 'Rupiah'
    },
    'ILS': { 
        code: 'ILS',
        fraction: 2,
        symbol: {
            normal: '₪'
        },
        unit: 'Shekel'
    },
    'INR': { 
        code: 'INR',
        fraction: 2,
        symbol: {
            normal: '₹'
        },
        unit: 'Rupee'
    },
    'IQD': { 
        code: 'IQD',
        fraction: 3,
        symbol: {
            normal: 'ID'
        },
        unit: 'Dinar'
    },
    'IRR': { 
        code: 'IRR',
        fraction: 2,
        symbol: {
            normal: 'Rl',
            plural: 'Rls'
        },
        unit: 'Rial'
    },
    'ISK': { 
        code: 'ISK',
        fraction: 0,
        symbol: {
            normal: 'kr'
        },
        unit: 'Krona'
    },
    'JMD': { 
        code: 'JMD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'JOD': { 
        code: 'JOD',
        fraction: 3,
        symbol: {
            normal: 'د.أ'
        },
        unit: 'Dinar'
    },
    'JPY': { 
        code: 'JPY',
        fraction: 0,
        symbol: {
            normal: '¥'
        },
        unit: 'Yen'
    },
    'KES': { 
        code: 'KES',
        fraction: 2,
        symbol: {
            normal: '/='
        },
        unit: 'Shilling'
    },
    'KGS': { 
        code: 'KGS',
        fraction: 2,
        symbol: {
            normal: '⃀'
        },
        unit: 'Som'
    },
    'KHR': { 
        code: 'KHR',
        fraction: 2,
        symbol: {
            normal: '៛'
        },
        unit: 'Riel'
    },
    'KMF': { 
        code: 'KMF',
        fraction: 0,
        symbol: {
            normal: 'CF'
        },
        unit: 'Franc'
    },
    'KPW': { 
        code: 'KPW',
        fraction: 2,
        symbol: {
            normal: '₩'
        },
        unit: 'Won'
    },
    'KRW': { 
        code: 'KRW',
        fraction: 0,
        symbol: {
            normal: '₩'
        },
        unit: 'Won'
    },
    'KWD': { 
        code: 'KWD',
        fraction: 3,
        symbol: {
            normal: 'KD'
        },
        unit: 'Dinar'
    },
    'KYD': { 
        code: 'KYD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'KZT': { 
        code: 'KZT',
        fraction: 2,
        symbol: {
            normal: '₸'
        },
        unit: 'Tenge'
    },
    'LAK': { 
        code: 'LAK',
        fraction: 2,
        symbol: {
            normal: '₭'
        },
        unit: 'Kip'
    },
    'LBP': { 
        code: 'LBP',
        fraction: 2,
        symbol: {
            normal: 'LL'
        },
        unit: 'Pound'
    },
    'LKR': { 
        code: 'LKR',
        fraction: 2,
        symbol: {
            normal: 'රු'
        },
        unit: 'Rupee'
    },
    'LRD': { 
        code: 'LRD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'LSL': { 
        code: 'LSL',
        fraction: 2,
        symbol: {
            normal: 'L',
            plural: 'M'
        },
        unit: 'Loti'
    },
    'LYD': { 
        code: 'LYD',
        fraction: 3,
        symbol: {
            normal: 'LD'
        },
        unit: 'Dinar'
    },
    'MAD': { 
        code: 'MAD',
        fraction: 2,
        symbol: {
            normal: 'DH'
        },
        unit: 'Dirham'
    },
    'MDL': { 
        code: 'MDL',
        fraction: 2,
        symbol: {
            normal: 'L'
        },
        unit: 'Leu'
    },
    'MGA': { 
        code: 'MGA',
        fraction: 2,
        symbol: {
            normal: 'Ar'
        },
        unit: 'Ariary'
    },
    'MKD': { 
        code: 'MKD',
        fraction: 2,
        symbol: {
            normal: 'DEN'
        },
        unit: 'Denar'
    },
    'MMK': { 
        code: 'MMK',
        fraction: 2,
        symbol: {
            normal: 'K',
            plural: 'Ks.'
        },
        unit: 'Kyat'
    },
    'MNT': { 
        code: 'MNT',
        fraction: 2,
        symbol: {
            normal: '₮'
        },
        unit: 'Tugrik'
    },
    'MOP': { 
        code: 'MOP',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Pataca'
    },
    'MRU': { 
        code: 'MRU',
        fraction: 2,
        symbol: {
            normal: 'UM'
        },
        unit: 'Ouguiya'
    },
    'MUR': { 
        code: 'MUR',
        fraction: 2,
        symbol: {
            normal: 'Re.',
            plural: 'Rs.'
        },
        unit: 'Rupee'
    },
    'MVR': { 
        code: 'MVR',
        fraction: 2,
        symbol: {
            normal: 'Rf.'
        },
        unit: 'Rufiyaa'
    },
    'MWK': { 
        code: 'MWK',
        fraction: 2,
        symbol: {
            normal: 'K'
        },
        unit: 'Kwacha'
    },
    'MXN': { 
        code: 'MXN',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'MYR': { 
        code: 'MYR',
        fraction: 2,
        symbol: {
            normal: 'RM'
        },
        unit: 'Ringgit'
    },
    'MZN': { 
        code: 'MZN',
        fraction: 2,
        symbol: {
            normal: 'MT'
        },
        unit: 'Metical'
    },
    'NAD': { 
        code: 'NAD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'NGN': { 
        code: 'NGN',
        fraction: 2,
        symbol: {
            normal: '₦'
        },
        unit: 'Naira'
    },
    'NIO': { 
        code: 'NIO',
        fraction: 2,
        symbol: {
            normal: 'C$'
        },
        unit: 'Cordoba'
    },
    'NOK': { 
        code: 'NOK',
        fraction: 2,
        symbol: {
            normal: 'kr'
        },
        unit: 'Krone'
    },
    'NPR': { 
        code: 'NPR',
        fraction: 2,
        symbol: {
            normal: 'रु'
        },
        unit: 'Rupee'
    },
    'NZD': { 
        code: 'NZD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'OMR': { 
        code: 'OMR',
        fraction: 3,
        symbol: {
            normal: 'R.O'
        },
        unit: 'Rial'
    },
    'PAB': { 
        code: 'PAB',
        fraction: 2,
        symbol: {
            normal: 'B/.'
        },
        unit: 'Balboa'
    },
    'PEN': { 
        code: 'PEN',
        fraction: 2,
        symbol: {
            normal: 'S/'
        },
        unit: 'Sol'
    },
    'PGK': { 
        code: 'PGK',
        fraction: 2,
        symbol: {
            normal: 'K'
        },
        unit: 'Kina'
    },
    'PHP': { 
        code: 'PHP',
        fraction: 2,
        symbol: {
            normal: '₱'
        },
        unit: 'Peso'
    },
    'PKR': { 
        code: 'PKR',
        fraction: 2,
        symbol: {
            normal: 'Re.',
            plural: 'Rs.'
        },
        unit: 'Rupee'
    },
    'PLN': { 
        code: 'PLN',
        fraction: 2,
        symbol: {
            normal: 'zł'
        },
        unit: 'Zloty'
    },
    'PYG': { 
        code: 'PYG',
        fraction: 0,
        symbol: {
            normal: '₲'
        },
        unit: 'Guarani'
    },
    'QAR': { 
        code: 'QAR',
        fraction: 2,
        symbol: {
            normal: 'QR'
        },
        unit: 'Rial'
    },
    'RON': { 
        code: 'RON',
        fraction: 2,
        symbol: {
            normal: 'L'
        },
        unit: 'Leu'
    },
    'RSD': { 
        code: 'RSD',
        fraction: 2,
        symbol: {
            normal: 'din.'
        },
        unit: 'Dinar'
    },
    'RUB': { 
        code: 'RUB',
        fraction: 2,
        symbol: {
            normal: '₽'
        },
        unit: 'Ruble'
    },
    'RWF': { 
        code: 'RWF',
        fraction: 0,
        symbol: {
            normal: 'FRw'
        },
        unit: 'Franc'
    },
    'SAR': { 
        code: 'SAR',
        fraction: 2,
        symbol: {
            normal: 'SAR'
        },
        unit: 'Riyal'
    },
    'SBD': { 
        code: 'SBD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'SCR': { 
        code: 'SCR',
        fraction: 2,
        symbol: {
            normal: 'Re.',
            plural: 'Rs.'
        },
        unit: 'Rupee'
    },
    'SDG': { 
        code: 'SDG',
        fraction: 2,
        symbol: {
            normal: 'LS'
        },
        unit: 'Pound'
    },
    'SEK': { 
        code: 'SEK',
        fraction: 2,
        symbol: {
            normal: 'kr'
        },
        unit: 'Krona'
    },
    'SGD': { 
        code: 'SGD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'SHP': { 
        code: 'SHP',
        fraction: 2,
        symbol: {
            normal: '£'
        },
        unit: 'Pound'
    },
    'SLE': { 
        code: 'SLE',
        fraction: 2,
        symbol: {
            normal: 'Le'
        },
        unit: 'Leone'
    },
    'SOS': { 
        code: 'SOS',
        fraction: 2,
        symbol: {
            normal: 'Sh.So.'
        },
        unit: 'Shilling'
    },
    'SRD': { 
        code: 'SRD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'SSP': { 
        code: 'SSP',
        fraction: 2,
        symbol: {
            normal: 'SS£'
        },
        unit: 'Pound'
    },
    'STN': { 
        code: 'STN',
        fraction: 2,
        symbol: {
            normal: 'Db'
        },
        unit: 'Dobra'
    },
    'SVC': { 
        code: 'SVC',
        fraction: 2,
        symbol: {
            normal: '₡'
        },
        unit: 'Colon'
    },
    'SYP': { 
        code: 'SYP',
        fraction: 2,
        symbol: {
            normal: 'LS'
        },
        unit: 'Pound'
    },
    'SZL': { 
        code: 'SZL',
        fraction: 2,
        symbol: {
            normal: 'E'
        },
        unit: 'Lilangeni'
    },
    'THB': { 
        code: 'THB',
        fraction: 2,
        symbol: {
            normal: '฿'
        },
        unit: 'Baht'
    },
    'TJS': { 
        code: 'TJS',
        fraction: 2,
        symbol: {
            normal: 'SM'
        },
        unit: 'Somoni'
    },
    'TMT': { 
        code: 'TMT',
        fraction: 2,
        symbol: {
            normal: 'm'
        },
        unit: 'Manat'
    },
    'TND': { 
        code: 'TND',
        fraction: 3,
        symbol: {
            normal: 'DT'
        },
        unit: 'Dinar'
    },
    'TOP': { 
        code: 'TOP',
        fraction: 2,
        symbol: {
            normal: 'T$'
        },
        unit: 'Paanga'
    },
    'TRY': { 
        code: 'TRY',
        fraction: 2,
        symbol: {
            normal: '₺'
        },
        unit: 'Lira'
    },
    'TTD': { 
        code: 'TTD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'TWD': { 
        code: 'TWD',
        fraction: 2,
        symbol: {
            normal: 'NT$'
        },
        unit: 'Dollar'
    },
    'TZS': { 
        code: 'TZS',
        fraction: 2,
        symbol: {
            normal: '/='
        },
        unit: 'Shilling'
    },
    'UAH': { 
        code: 'UAH',
        fraction: 2,
        symbol: {
            normal: '₴'
        },
        unit: 'Hryvnia'
    },
    'UGX': { 
        code: 'UGX',
        fraction: 0,
        symbol: {
            normal: '/='
        },
        unit: 'Shilling'
    },
    'USD': { 
        code: 'USD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'UYU': { 
        code: 'UYU',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Peso'
    },
    'UZS': { 
        code: 'UZS',
        fraction: 2,
        unit: 'Sum'
    },
    'VED': { 
        code: 'VED',
        fraction: 2,
        symbol: {
            normal: 'Bs.D'
        },
        unit: 'Bolivar'
    },
    'VES': { 
        code: 'VES',
        fraction: 2,
        symbol: {
            normal: 'Bs.S'
        },
        unit: 'Bolivar'
    },
    'VND': { 
        code: 'VND',
        fraction: 0,
        symbol: {
            normal: '₫'
        },
        unit: 'Dong'
    },
    'VUV': { 
        code: 'VUV',
        fraction: 0,
        symbol: {
            normal: 'VT'
        },
        unit: 'Vatu'
    },
    'WST': { 
        code: 'WST',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Tala'
    },
    'XAF': { 
        code: 'XAF',
        fraction: 0,
        symbol: {
            normal: 'F.CFA'
        },
        unit: 'Franc'
    },
    'XCD': { 
        code: 'XCD',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    },
    'XOF': { 
        code: 'XOF',
        fraction: 0,
        symbol: {
            normal: 'F.CFA'
        },
        unit: 'Franc'
    },
    'XPF': { 
        code: 'XPF',
        fraction: 0,
        symbol: {
            normal: 'F'
        },
        unit: 'Franc'
    },
    'XSU': { 
        code: 'XSU',
        symbol: {
            normal: 'S/.'
        },
        unit: 'Sucre'
    },
    'YER': { 
        code: 'YER',
        fraction: 2,
        symbol: {
            normal: 'YRl',
            plural: 'YRls'
        },
        unit: 'Rial'
    },
    'ZAR': { 
        code: 'ZAR',
        fraction: 2,
        symbol: {
            normal: 'R'
        },
        unit: 'Rand'
    },
    'ZMW': { 
        code: 'ZMW',
        fraction: 2,
        symbol: {
            normal: 'K'
        },
        unit: 'Kwacha'
    },
    'ZWG': { 
        code: 'ZWG',
        fraction: 2,
        symbol: {
            normal: 'ZiG'
        },
        unit: 'ZiG'
    },
    'ZWL': { 
        code: 'ZWL',
        fraction: 2,
        symbol: {
            normal: '$'
        },
        unit: 'Dollar'
    }
};

export const DEFAULT_CURRENCY_SYMBOL: string = '¤';
export const DEFAULT_CURRENCY_CODE: string = (ALL_CURRENCIES['USD'] as CurrencyInfo).code;
export const PARENT_ACCOUNT_CURRENCY_PLACEHOLDER: string = '---';
