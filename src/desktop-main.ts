import { createApp, ref } from 'vue';
import { createPinia } from 'pinia';
import { type I18n, type Composer, createI18n } from 'vue-i18n';

import { createVuetify } from 'vuetify';
import { VAlert } from 'vuetify/components/VAlert';
import { VApp } from 'vuetify/components/VApp';
import { VAvatar } from 'vuetify/components/VAvatar';
import { VAutocomplete } from 'vuetify/components/VAutocomplete';
import { VBadge } from 'vuetify/components/VBadge';
import { VBtn } from 'vuetify/components/VBtn';
import { VBtnGroup } from 'vuetify/components/VBtnGroup';
import { VBtnToggle } from 'vuetify/components/VBtnToggle';
import { VCard, VCardActions, VCardItem, VCardSubtitle, VCardText, VCardTitle } from 'vuetify/components/VCard';
import { VCheckbox, VCheckboxBtn } from 'vuetify/components/VCheckbox';
import { VChip } from 'vuetify/components/VChip';
import { VDataTable } from 'vuetify/components/VDataTable';
import { VDialog } from 'vuetify/components/VDialog';
import { VDivider } from 'vuetify/components/VDivider';
import { VExpansionPanel, VExpansionPanelText, VExpansionPanelTitle, VExpansionPanels } from 'vuetify/components/VExpansionPanel';
import { VForm } from 'vuetify/components/VForm';
import { VContainer, VCol, VRow, VSpacer } from 'vuetify/components/VGrid';
import { VIcon } from 'vuetify/components/VIcon';
import { VImg } from 'vuetify/components/VImg';
import { VInput } from 'vuetify/components/VInput';
import { VLabel } from 'vuetify/components/VLabel';
import { VLayout } from 'vuetify/components/VLayout';
import { VList, VListGroup, VListImg, VListItem, VListItemAction, VListItemMedia, VListItemSubtitle, VListItemTitle, VListSubheader } from 'vuetify/components/VList';
import { VMain } from 'vuetify/components/VMain';
import { VMenu } from 'vuetify/components/VMenu';
import { VNavigationDrawer } from 'vuetify/components/VNavigationDrawer';
import { VOverlay } from 'vuetify/components/VOverlay';
import { VPagination } from 'vuetify/components/VPagination';
import { VProgressCircular } from 'vuetify/components/VProgressCircular';
import { VProgressLinear } from 'vuetify/components/VProgressLinear';
import { VSelect } from 'vuetify/components/VSelect';
import { VSkeletonLoader } from 'vuetify/components/VSkeletonLoader';
import { VSlideGroup, VSlideGroupItem } from 'vuetify/components/VSlideGroup';
import { VSnackbar } from 'vuetify/components/VSnackbar';
import { VSwitch } from 'vuetify/components/VSwitch';
import { VTabs, VTab } from 'vuetify/components/VTabs';
import { VTable } from 'vuetify/components/VTable';
import { VTextarea } from 'vuetify/components/VTextarea';
import { VTextField } from 'vuetify/components/VTextField';
import { VToolbar } from 'vuetify/components/VToolbar';
import { VTooltip } from 'vuetify/components/VTooltip';
import { VWindow, VWindowItem } from 'vuetify/components/VWindow';
import type { LocaleInstance } from 'vuetify/lib/framework.d.ts';

import { aliases, mdi } from 'vuetify/iconsets/mdi-svg';
import 'vuetify/styles';

import * as echarts from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { LineChart, BarChart, PieChart, ScatterChart, CandlestickChart, RadarChart, SankeyChart } from 'echarts/charts';
import {
    GridComponent,
    TooltipComponent,
    LegendComponent,
} from 'echarts/components';
import VChart from 'vue-echarts';

import 'line-awesome/dist/line-awesome/css/line-awesome.css';

import { PerfectScrollbar } from 'vue3-perfect-scrollbar';

import { VueDatePicker } from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';

import draggable from 'vuedraggable';

import router from '@/router/desktop.ts';

import { DecimalSeparator } from '@/core/numeral.ts';
import { getI18nOptions, getRtlLocales } from '@/locales/helpers.ts';

import PinCodeInput from '@/components/common/PinCodeInput.vue';
import MapView from '@/components/common/MapView.vue';
import DateTimePicker from '@/components/common/DateTimePicker.vue';
import MonthPicker from '@/components/common/MonthPicker.vue';
import TransactionCalendar from '@/components/common/TransactionCalendar.vue';

import ItemIcon from '@/components/desktop/ItemIcon.vue';
import BtnVerticalGroup from '@/components/desktop/BtnVerticalGroup.vue';
import NumberInput from '@/components/desktop/NumberInput.vue';
import AmountInput from '@/components/desktop/AmountInput.vue';
import AmountInputDialog from '@/components/desktop/AmountInputDialog.vue';
import LanguageSelect from '@/components/desktop/LanguageSelect.vue';
import LanguageSelectButton from '@/components/desktop/LanguageSelectButton.vue';
import CurrencySelect from '@/components/desktop/CurrencySelect.vue';
import DateTimeSelect from '@/components/desktop/DateTimeSelect.vue';
import DateSelect from '@/components/desktop/DateSelect.vue';
import FiscalYearStartSelect from '@/components/desktop/FiscalYearStartSelect.vue';
import ColorSelect from '@/components/desktop/ColorSelect.vue';
import IconSelect from '@/components/desktop/IconSelect.vue';
import TwoColumnSelect from '@/components/desktop/TwoColumnSelect.vue';
import TransactionTagAutoComplete from '@/components/desktop/TransactionTagAutoComplete.vue';
import ScheduleFrequencySelect from '@/components/desktop/ScheduleFrequencySelect.vue';
import StepsBar from '@/components/desktop/StepsBar.vue';
import ConfirmDialog from '@/components/desktop/ConfirmDialog.vue';
import SnackBar from '@/components/desktop/SnackBar.vue';
import PieChartComponent from '@/components/desktop/PieChart.vue';
import RadarChartComponent from '@/components/desktop/RadarChart.vue';
import AxisChart from '@/components/desktop/AxisChart.vue';
import TrendsChart from '@/components/desktop/TrendsChart.vue';
import RenameDialog from '@/components/desktop/RenameDialog.vue';
import DateRangeSelectionDialog from '@/components/desktop/DateRangeSelectionDialog.vue';
import MonthSelectionDialog from '@/components/desktop/MonthSelectionDialog.vue';
import MonthRangeSelectionDialog from '@/components/desktop/MonthRangeSelectionDialog.vue';
import AccountBalanceTrendsChart from '@/components/desktop/AccountBalanceTrendsChart.vue';
import AccountAndCategorySankeyChart from '@/components/desktop/AccountAndCategorySankeyChart.vue';
import SwitchToMobileDialog from '@/components/desktop/SwitchToMobileDialog.vue';

import TextFieldAutoWidth from '@/directives/desktop/textfieldAutoWidth.ts';

import '@/styles/desktop/template/vuetify/index.scss';
import '@/styles/desktop/template/template/index.scss';
import '@/styles/desktop/template/layout/index.scss';
import '@/styles/desktop/template/layout/component/index.scss';
import '@/styles/desktop/template/layout/_default-layout.scss';
import '@/styles/desktop/global.scss';
import '@/styles/desktop/font-size.scss';
import '@/styles/desktop/amount-color.scss';

import App from './DesktopApp.vue';

const app = createApp(App);
const pinia = createPinia();
const i18n = createI18n(getI18nOptions()) as I18n<Record<string, unknown>, Record<string, unknown>, Record<string, unknown>, string, false>;
const vuetify = createVuetify({
    components: {
        VAlert,
        VApp,
        VAvatar,
        VAutocomplete,
        VBadge,
        VBtn,
        VBtnGroup,
        VBtnToggle,
        VCard,
        VCardActions,
        VCardItem,
        VCardSubtitle,
        VCardText,
        VCardTitle,
        VCheckbox,
        VCheckboxBtn,
        VChip,
        VDataTable,
        VDialog,
        VDivider,
        VExpansionPanel,
        VExpansionPanelText,
        VExpansionPanelTitle,
        VExpansionPanels,
        VForm,
        VContainer,
        VCol,
        VRow,
        VSpacer,
        VIcon,
        VImg,
        VInput,
        VLabel,
        VLayout,
        VList,
        VListGroup,
        VListImg,
        VListItem,
        VListItemAction,
        VListItemMedia,
        VListItemSubtitle,
        VListItemTitle,
        VListSubheader,
        VMain,
        VMenu,
        VNavigationDrawer,
        VOverlay,
        VPagination,
        VProgressCircular,
        VProgressLinear,
        VSelect,
        VSkeletonLoader,
        VSlideGroup,
        VSlideGroupItem,
        VSnackbar,
        VSwitch,
        VTabs,
        VTab,
        VTable,
        VTextarea,
        VTextField,
        VToolbar,
        VTooltip,
        VWindow,
        VWindowItem
    },
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: {
            mdi
        }
    },
    defaults: {
        VAlert: {
            VBtn: {
                color: undefined
            }
        },
        VAutocomplete: {
            variant: 'outlined',
            density: 'comfortable',
            color: 'primary',
            hideDetails: 'auto'
        },
        VAvatar: {
            variant: 'flat',
            VIcon: {
                size: 24,
            },
        },
        VBadge: {
            color: 'primary'
        },
        VBtn: {
            color: 'primary'
        },
        VCheckbox: {
            color: 'primary',
            hideDetails: 'auto'
        },
        VChip: {
            elevation: 0
        },
        VList: {
            color: 'primary'
        },
        VPagination: {
            density: 'comfortable',
            activeColor: 'primary'
        },
        VRadio: {
            density: 'comfortable',
            color: 'primary',
            hideDetails: 'auto'
        },
        VSelect: {
            variant: 'outlined',
            density: 'comfortable',
            color: 'primary',
            hideDetails: 'auto'
        },
        VSlider: {
            color: 'primary',
            hideDetails: 'auto'
        },
        VSwitch: {
            inset: true,
            color: 'primary',
            hideDetails: 'auto'
        },
        VProgressCircular: {
            size: 40
        },
        VSnackbar: {
            timeout: 3000
        },
        VTable: {
            hover: true
        },
        VTabs: {
            color: 'primary',
            VSlideGroup: {
                showArrows: true
            }
        },
        VTextarea: {
            variant: 'outlined',
            density: 'comfortable',
            color: 'primary',
            hideDetails: 'auto'
        },
        VTextField: {
            variant: 'outlined',
            density: 'comfortable',
            color: 'primary',
            hideDetails: 'auto'
        },
        VToolbar: {
            color: 'primary'
        },
        VTooltip: {
            location: 'top'
        },
        VWindow: {
            touch: false
        }
    },
    theme: {
        defaultTheme: 'light',
        themes: {
            light: {
                dark: false,
                colors: {
                    'primary': '#10B981',
                    'primary-darken-1': '#064E3B',
                    'on-primary': '#ffffff',
                    'secondary': '#064E3B',
                    'secondary-darken-1': '#022C22',
                    'on-secondary': '#ffffff',
                    'success': '#10B981',
                    'success-darken-1': '#059669',
                    'on-success': '#ffffff',
                    'info': '#0ea5e9',
                    'info-darken-1': '#0284c7',
                    'on-info': '#ffffff',
                    'warning': '#f59e0b',
                    'warning-darken-1': '#d97706',
                    'on-warning': '#ffffff',
                    'error': '#f43f5e',
                    'error-darken-1': '#e11d48',
                    'on-error': '#ffffff',
                    'teal': '#10B981',
                    'background': '#F8FAFC',
                    'on-background': '#0F172A',
                    'surface': '#FFFFFF',
                    'on-surface': '#0F172A',
                    'notification-background': '#FFFFFF',
                    'on-notification-background': '#0F172A',
                    'grey': '#94A3B8',
                    'grey-50': '#F8FAFC',
                    'grey-100': '#F1F5F9',
                    'grey-200': '#E2E8F0',
                    'grey-300': '#CBD5E1',
                    'grey-400': '#94A3B8',
                    'grey-500': '#64748B',
                    'grey-600': '#475569',
                    'grey-700': '#334155',
                    'grey-800': '#1E293B',
                    'grey-900': '#0F172A',
                    'perfect-scrollbar-thumb': '#CBD5E1',
                    'skin-bordered-background': '#FFFFFF',
                    'skin-bordered-surface': '#FFFFFF',
                    'expansion-panel-text-custom-bg': '#F8FAFC'
                },
                variables: {
                    'code-color': '#10B981',
                    'overlay-scrim-background': '#0F172A',
                    'tooltip-background': '#0F172A',
                    'tooltip-color': '#ffffff',
                    'overlay-scrim-opacity': 0.4,
                    'hover-opacity': 0.04,
                    'focus-opacity': 0.1,
                    'selected-opacity': 0.08,
                    'activated-opacity': 0.12,
                    'pressed-opacity': 0.14,
                    'dragged-opacity': 0.1,
                    'disabled-opacity': 0.4,
                    'border-color': '#0F172A',
                    'border-opacity': 0.08,
                    'table-header-color': '#F8FAFC',
                    'high-emphasis-opacity': 0.9,
                    'medium-emphasis-opacity': 0.7,
                    'shadow-key-umbra-color': 'transparent',
                    'shadow-xs-opacity': '0',
                    'shadow-sm-opacity': '0',
                    'shadow-md-opacity': '0',
                    'shadow-lg-opacity': '0',
                    'shadow-xl-opacity': '0',
                }
            },
            dark: {
                dark: true,
                colors: {
                    'primary': '#10B981',
                    'primary-darken-1': '#064E3B',
                    'on-primary': '#ffffff',
                    'secondary': '#34D399',
                    'secondary-darken-1': '#059669',
                    'on-secondary': '#064E3B',
                    'success': '#10B981',
                    'success-darken-1': '#059669',
                    'on-success': '#ffffff',
                    'info': '#38BDF8',
                    'info-darken-1': '#0284C7',
                    'on-info': '#ffffff',
                    'warning': '#FBBF24',
                    'warning-darken-1': '#D97706',
                    'on-warning': '#0F172A',
                    'error': '#FB7185',
                    'error-darken-1': '#E11D48',
                    'on-error': '#ffffff',
                    'teal': '#34D399',
                    'background': '#0B0F19',
                    'on-background': '#F8FAFC',
                    'surface': '#111827',
                    'on-surface': '#F8FAFC',
                    'notification-background': '#1F2937',
                    'on-notification-background': '#F8FAFC',
                    'grey': '#64748B',
                    'grey-50': '#0F172A',
                    'grey-100': '#1E293B',
                    'grey-200': '#334155',
                    'grey-300': '#475569',
                    'grey-400': '#64748B',
                    'grey-500': '#94A3B8',
                    'grey-600': '#CBD5E1',
                    'grey-700': '#E2E8F0',
                    'grey-800': '#F1F5F9',
                    'grey-900': '#F8FAFC',
                    'perfect-scrollbar-thumb': '#334155',
                    'skin-bordered-background': '#111827',
                    'skin-bordered-surface': '#111827',
                    'expansion-panel-text-custom-bg': '#1E293B'
                },
                variables: {
                    'code-color': '#34D399',
                    'overlay-scrim-background': '#000000',
                    'tooltip-background': '#1F2937',
                    'tooltip-color': '#F8FAFC',
                    'overlay-scrim-opacity': 0.6,
                    'hover-opacity': 0.04,
                    'focus-opacity': 0.1,
                    'selected-opacity': 0.08,
                    'activated-opacity': 0.12,
                    'pressed-opacity': 0.14,
                    'disabled-opacity': 0.4,
                    'dragged-opacity': 0.1,
                    'border-color': '#FFFFFF',
                    'border-opacity': 0.08,
                    'table-header-color': '#1E293B',
                    'high-emphasis-opacity': 0.9,
                    'medium-emphasis-opacity': 0.7,
                    'shadow-key-umbra-color': 'transparent',
                    'shadow-xs-opacity': '0',
                    'shadow-sm-opacity': '0',
                    'shadow-md-opacity': '0',
                    'shadow-lg-opacity': '0',
                    'shadow-xl-opacity': '0',
                }
            }
        }
    },
    locale: {
        rtl: getRtlLocales(),
        adapter: ((i18nGlobal: Composer) => {
            const instance: LocaleInstance = {
                name: 'Split Chill AI i18n',
                decimalSeparator: ref<string>(DecimalSeparator.Default.symbol), 
                messages: i18nGlobal.messages,
                current: i18nGlobal.locale,
                fallback: i18nGlobal.locale, 
                t: (key: string, ...params: unknown[]): string => {
                    if (!key) {
                        return '';
                    }

                    if (!key.startsWith('$vuetify.')) {
                        return key;
                    }

                    key = key.substring(9); 
                    const mappedTextKey = vuetifyI18nTextKeyMap[key];

                    if (!mappedTextKey) {
                        return key;
                    }

                    if (params && params.length > 0) {
                        return i18nGlobal.t(mappedTextKey, params) as string;
                    } else {
                        return i18nGlobal.t(mappedTextKey) as string;
                    }
                },
                n: (value: number): string => {
                    return i18nGlobal.n(value);
                },
                provide: (): LocaleInstance => {
                    return instance;
                }
            };

            return instance;
        })(i18n.global) as LocaleInstance,
    }
});


const vuetifyI18nTextKeyMap: Record<string, string> = {
    'open': 'Open',
    'close': 'Close'
}

echarts.use([
    CanvasRenderer,
    LineChart,
    BarChart,
    PieChart,
    ScatterChart,
    CandlestickChart,
    RadarChart,
    SankeyChart,
    GridComponent,
    TooltipComponent,
    LegendComponent
]);

app.use(pinia);
app.use(i18n);
app.use(vuetify);
app.use(router);

app.component('VChart', VChart);
app.component('PerfectScrollbar', PerfectScrollbar);
app.component('VueDatePicker', VueDatePicker);
app.component('DraggableList', draggable);

app.component('PinCodeInput', PinCodeInput);
app.component('MapView', MapView);
app.component('DateTimePicker', DateTimePicker);
app.component('MonthPicker', MonthPicker);
app.component('TransactionCalendar', TransactionCalendar);

app.component('ItemIcon', ItemIcon);
app.component('BtnVerticalGroup', BtnVerticalGroup);
app.component('NumberInput', NumberInput);
app.component('AmountInput', AmountInput);
app.component('AmountInputDialog', AmountInputDialog);
app.component('LanguageSelect', LanguageSelect);
app.component('LanguageSelectButton', LanguageSelectButton);
app.component('CurrencySelect', CurrencySelect);
app.component('DateTimeSelect', DateTimeSelect);
app.component('DateSelect', DateSelect);
app.component('FiscalYearStartSelect', FiscalYearStartSelect);
app.component('ColorSelect', ColorSelect);
app.component('IconSelect', IconSelect);
app.component('TwoColumnSelect', TwoColumnSelect);
app.component('TransactionTagAutoComplete', TransactionTagAutoComplete);
app.component('ScheduleFrequencySelect', ScheduleFrequencySelect);
app.component('StepsBar', StepsBar);
app.component('ConfirmDialog', ConfirmDialog);
app.component('SnackBar', SnackBar);
app.component('PieChart', PieChartComponent);
app.component('RadarChart', RadarChartComponent);
app.component('AxisChart', AxisChart);
app.component('TrendsChart', TrendsChart);
app.component('RenameDialog', RenameDialog);
app.component('DateRangeSelectionDialog', DateRangeSelectionDialog);
app.component('MonthSelectionDialog', MonthSelectionDialog);
app.component('MonthRangeSelectionDialog', MonthRangeSelectionDialog);
app.component('AccountBalanceTrendsChart', AccountBalanceTrendsChart);
app.component('AccountAndCategorySankeyChart', AccountAndCategorySankeyChart);
app.component('SwitchToMobileDialog', SwitchToMobileDialog);

app.directive('TextFieldAutoWidth', TextFieldAutoWidth);

app.mount('#app');
