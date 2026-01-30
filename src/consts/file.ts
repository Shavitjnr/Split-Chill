import type { ImportFileCategoryAndTypes } from '@/core/file.ts';

export const SUPPORTED_IMAGE_EXTENSIONS: string = '.jpg,.jpeg,.png,.gif,.webp';

export const DEFAULT_DOCUMENT_LANGUAGE_FOR_IMPORT_FILE: string = 'en';
export const SUPPORTED_DOCUMENT_LANGUAGES_FOR_IMPORT_FILE: Record<string, string> = {
    DEFAULT_DOCUMENT_LANGUAGE_FOR_IMPORT_FILE: DEFAULT_DOCUMENT_LANGUAGE_FOR_IMPORT_FILE,
    'zh-Hans': 'zh-Hans',
    'zh-Hant': 'zh-Hans',
};

export const UTF_8 = 'utf-8';

export const SUPPORTED_FILE_ENCODINGS: string[] = [
    UTF_8, 
    'utf-8-bom', 
    'utf-16le', 
    'utf-16be', 
    'utf-16le-bom', 
    'utf-16be-bom', 
    'cp437', 
    'cp863', 
    'cp037', 
    'cp1047', 
    'cp1140', 
    "iso-8859-1", 
    'cp850', 
    'cp858', 
    'windows-1252', 
    'iso-8859-15', 
    'iso-8859-4', 
    'iso-8859-10', 
    'cp865', 
    'iso-8859-2', 
    'cp852', 
    'windows-1250', 
    'iso-8859-14', 
    'iso-8859-3', 
    'cp860', 
    'iso-8859-7', 
    'windows-1253', 
    'iso-8859-9', 
    'windows-1254', 
    'iso-8859-13', 
    'windows-1257', 
    'iso-8859-16', 
    'iso-8859-5', 
    'cp855', 
    'cp866', 
    'windows-1251', 
    'koi8r', 
    'koi8u', 
    'iso-8859-6', 
    'windows-1256', 
    'iso-8859-8', 
    'cp862', 
    'windows-1255', 
    'windows-874', 
    'windows-1258', 
    'gb18030', 
    'gbk', 
    'big5', 
    'euc-kr', 
    'euc-jp', 
    'iso-2022-jp', 
    'shift_jis', 
];

export const CHARDET_ENCODING_NAME_MAPPING: Record<string, string> = {
    'UTF-8': UTF_8,
    'UTF-16LE': 'utf-16le',
    'UTF-16BE': 'utf-16be',
    
    
    'ISO-2022-JP': 'iso-2022-jp',
    
    
    'Shift_JIS': 'shift_jis',
    'Big5': 'big5',
    'EUC-JP': 'euc-jp',
    'EUC-KR': 'euc-kr',
    'GB18030': 'gb18030',
    'ISO-8859-1': 'iso-8859-1',
    'ISO-8859-2': 'iso-8859-2',
    'ISO-8859-5': 'iso-8859-5',
    'ISO-8859-6': 'iso-8859-6',
    'ISO-8859-7': 'iso-8859-7',
    'ISO-8859-8': 'iso-8859-8',
    'ISO-8859-9': 'iso-8859-9',
    'windows-1250': 'windows-1250',
    'windows-1251': 'windows-1251',
    'windows-1252': 'windows-1252',
    'windows-1253': 'windows-1253',
    'windows-1254': 'windows-1254',
    'windows-1255': 'windows-1255',
    'windows-1256': 'windows-1256',
    'KOI8-R':'koi8r'
};

export const SUPPORTED_IMPORT_FILE_CATEGORY_AND_TYPES: ImportFileCategoryAndTypes[] = [
    {
        categoryName: 'Split Chill AI File Format',
        fileTypes: [
            {
                type: 'splitchill_ai',
                name: 'Split Chill AI Data Export File',
                extensions: '.csv,.tsv',
                subTypes: [
                    {
                        type: 'splitchill_ai_csv',
                        name: 'CSV (Comma-separated values) File',
                        extensions: '.csv',
                    },
                    {
                        type: 'splitchill_ai_tsv',
                        name: 'TSV (Tab-separated values) File',
                        extensions: '.tsv',
                    }
                ],
                document: {
                    supportMultiLanguages: true,
                    anchor: 'export-transactions'
                }
            }
        ]
    },
    {
        categoryName: 'Custom File Format',
        fileTypes: [
            {
                type: 'dsv',
                name: 'Delimiter-separated Values (DSV) File',
                extensions: '.csv,.tsv',
                subTypes: [
                    {
                        type: 'custom_csv',
                        name: 'CSV (Comma-separated values) File',
                        extensions: '.csv',
                    },
                    {
                        type: 'custom_tsv',
                        name: 'TSV (Tab-separated values) File',
                        extensions: '.tsv,.txt',
                    },
                    {
                        type: 'custom_ssv',
                        name: 'SSV (Semicolon-separated values) File',
                        extensions: '.txt',
                    }
                ],
                supportedEncodings: SUPPORTED_FILE_ENCODINGS,
                document: {
                    supportMultiLanguages: true,
                    anchor: 'how-to-import-delimiter-separated-values-dsv-file-or-data'
                }
            },
            {
                type: 'dsv_data',
                name: 'Delimiter-separated Values (DSV) Data',
                extensions: '.csv,.tsv',
                subTypes: [
                    {
                        type: 'custom_csv',
                        name: 'CSV (Comma-separated values) File',
                        extensions: '.csv',
                    },
                    {
                        type: 'custom_tsv',
                        name: 'TSV (Tab-separated values) File',
                        extensions: '.tsv,.txt',
                    },
                    {
                        type: 'custom_ssv',
                        name: 'SSV (Semicolon-separated values) File',
                        extensions: '.txt',
                    }
                ],
                dataFromTextbox: true,
                document: {
                    supportMultiLanguages: true,
                    anchor: 'how-to-import-delimiter-separated-values-dsv-file-or-data'
                }
            }
        ]
    },
    {
        categoryName: 'General Data Exchange Format',
        fileTypes: [
            {
                type: 'ofx',
                name: 'Open Financial Exchange (OFX) File',
                extensions: '.ofx'
            },
            {
                type: 'qfx',
                name: 'Quicken Financial Exchange (QFX) File',
                extensions: '.qfx'
            },
            {
                type: 'qif',
                name: 'Quicken Interchange Format (QIF) File',
                extensions: '.qif',
                subTypes: [
                    {
                        type: 'qif_ymd',
                        name: 'Year-month-day format',
                    },
                    {
                        type: 'qif_mdy',
                        name: 'Month-day-year format',
                    },
                    {
                        type: 'qif_dmy',
                        name: 'Day-month-year format',
                    }
                ],
                supportedAdditionalOptions: {
                    payeeAsTag: false,
                    payeeAsDescription: true
                }
            },
            {
                type: 'iif',
                name: 'Intuit Interchange Format (IIF) File',
                extensions: '.iif'
            }
        ]
    },
    {
        categoryName: 'General Bank Statement Format',
        fileTypes: [
            {
                type: 'camt052',
                name: 'Camt.052 Bank to Customer Statement File',
                extensions: '.xml'
            },
            {
                type: 'camt053',
                name: 'Camt.053 Bank to Customer Statement File',
                extensions: '.xml'
            },
            {
                type: 'mt940',
                name: 'MT940 Consumer Statement Message File',
                extensions: '.txt'
            }
        ]
    },
    {
        categoryName: 'Other Bank/Payment App Statement File',
        fileTypes: [
            {
                type: 'alipay_app_csv',
                name: 'Alipay (App) Statement File',
                extensions: '.csv',
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取支付宝app交易流水文件'
                }
            },
            {
                type: 'alipay_web_csv',
                name: 'Alipay (Web) Statement File',
                extensions: '.csv',
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取支付宝网页版交易流水文件'
                }
            },
            {
                type: 'wechat_pay_app',
                name: 'WeChat Pay Statement File',
                extensions: '.xlsx,.csv',
                subTypes: [
                    {
                        type: 'wechat_pay_app_xlsx',
                        name: 'Excel Workbook File',
                        extensions: '.xlsx',
                    },
                    {
                        type: 'wechat_pay_app_csv',
                        name: 'CSV (Comma-separated values) File',
                        extensions: '.csv',
                    }
                ],
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取微信支付账单文件'
                }
            },
            {
                type: 'jdcom_finance_app_csv',
                name: 'JD.com Finance Statement File',
                extensions: '.csv',
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取京东金融账单文件'
                }
            }
        ]
    },
    {
        categoryName: 'Other Finance App File Format',
        fileTypes: [
            {
                type: 'gnucash',
                name: 'GnuCash XML Database File',
                extensions: '.gnucash',
                document: {
                    supportMultiLanguages: true,
                    anchor: 'how-to-get-gnucash-xml-database-file'
                }
            },
            {
                type: 'firefly_iii_csv',
                name: 'Firefly III Data Export File',
                extensions: '.csv',
                document: {
                    supportMultiLanguages: true,
                    anchor: 'how-to-get-firefly-iii-data-export-file'
                }
            },
            {
                type: 'beancount',
                name: 'Beancount Data File',
                extensions: '.beancount'
            },
            {
                type: 'feidee_mymoney_csv',
                name: 'Feidee MyMoney (App) Data Export File',
                extensions: '.csv',
                supportedAdditionalOptions: {
                    memberAsTag: false,
                    projectAsTag: false,
                    merchantAsTag: false,
                },
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取随手记app数据导出文件'
                }
            },
            {
                type: 'feidee_mymoney_xls',
                name: 'Feidee MyMoney (Web) Data Export File',
                extensions: '.xls',
                supportedAdditionalOptions: {
                    memberAsTag: false,
                    projectAsTag: false,
                    merchantAsTag: false,
                },
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取随手记web版数据导出文件'
                }
            },
            {
                type: 'feidee_mymoney_elecloud_xlsx',
                name: 'Feidee MyMoney (Elecloud) Data Export File',
                extensions: '.xlsx',
                supportedAdditionalOptions: {
                    memberAsTag: false,
                    projectAsTag: false,
                    merchantAsTag: false,
                },
                document: {
                    supportMultiLanguages: 'zh-Hans',
                    anchor: '如何获取随手记神象云账本数据导出文件'
                }
            }
        ]
    }
];
