<template>
    <v-row class="match-height">
        <v-col cols="12" lg="4" md="12">
            <v-card class="executive-hero-card" :class="{ 'disabled': loadingOverview }">
                <v-card-text class="d-flex flex-column justify-space-between h-100 pa-6">
                    <div>
                        <!-- Header pill with month and refresh -->
                        <div class="d-flex align-center justify-space-between mb-4">
                            <div class="overview-badge-pill">
                                <v-icon :icon="mdiCalendarOutline" size="14" class="me-1 text-emerald" />
                                <span>{{ displayDateRange?.thisMonth?.displayTime }} • {{ tt('Expense') }}</span>
                            </div>
                            <v-btn density="compact" color="default" variant="text" size="26"
                                   :icon="true" :loading="loadingOverview" @click="reload(true)">
                                <template #loader>
                                    <v-progress-circular indeterminate size="18"/>
                                </template>
                                <v-icon :icon="mdiRefresh" size="20" />
                                <v-tooltip activator="parent">{{ tt('Refresh') }}</v-tooltip>
                            </v-btn>
                        </div>

                        <!-- Main Expense Amount -->
                        <div class="overview-amount-container d-flex align-center mt-2 mb-1">
                            <span class="overview-main-amount" v-if="!loadingOverview || (transactionOverview && transactionOverview.thisMonth && transactionOverview.thisMonth.valid)">
                                {{ transactionOverview && transactionOverview.thisMonth ? getDisplayExpenseAmount(transactionOverview.thisMonth) : '-' }}
                            </span>
                            <v-skeleton-loader class="d-inline-block skeleton-no-margin" width="140px" type="text" :loading="true" v-else-if="loadingOverview && (!transactionOverview || !transactionOverview.thisMonth || !transactionOverview.thisMonth.valid)"></v-skeleton-loader>
                            <v-btn class="ms-2" density="compact" color="default" variant="text"
                                   :icon="true" @click="showAmountInHomePage = !showAmountInHomePage">
                                <v-icon :icon="showAmountInHomePage ? mdiEyeOffOutline : mdiEyeOutline" size="20" />
                            </v-btn>
                        </div>

                        <!-- Monthly Income Sub-metric -->
                        <div class="overview-income-box d-flex align-center mt-3 mb-6">
                            <span class="overview-income-indicator me-2">
                                <v-icon :icon="mdiArrowUp" size="14" class="text-emerald" />
                            </span>
                            <span class="text-caption font-weight-medium text-slate-muted me-2">{{ tt('Monthly income') }}:</span>
                            <span class="text-body-2 font-weight-bold text-emerald" v-if="!loadingOverview || (transactionOverview && transactionOverview.thisMonth && transactionOverview.thisMonth.valid)">
                                {{ transactionOverview && transactionOverview.thisMonth ? getDisplayIncomeAmount(transactionOverview.thisMonth) : '-' }}
                            </span>
                            <v-skeleton-loader class="d-inline-block skeleton-no-margin" width="80px" type="text" :loading="true" v-else-if="loadingOverview && (!transactionOverview || !transactionOverview.thisMonth || !transactionOverview.thisMonth.valid)"></v-skeleton-loader>
                        </div>
                    </div>

                    <!-- Bottom Action -->
                    <div>
                        <v-btn size="small" variant="flat" color="primary" class="px-5 font-weight-semibold" :to="`/transaction/list?${overviewStore.getTransactionListPageParams({ dateType: DateRange.ThisMonth.type })}`">
                            {{ tt('View Details') }}
                        </v-btn>
                    </div>
                </v-card-text>
            </v-card>
        </v-col>

        <v-col cols="12" lg="8" md="12">
            <v-card class="asset-summary-card" :class="{ 'disabled': loadingOverview }">
                <template #title>
                    <div class="d-flex align-center justify-space-between pt-1">
                        <span class="font-weight-bold text-h6">{{ tt('Asset Summary') }}</span>
                        <span class="text-caption text-slate-muted font-weight-medium" v-if="!loadingOverview || (allAccounts && allAccounts.length)">
                            {{ tt('format.misc.youHaveAccounts', { count: displayAccountCount }) }}
                        </span>
                        <v-skeleton-loader class="skeleton-no-margin" width="140px" type="text" :loading="true" v-else-if="loadingOverview && (!allAccounts || !allAccounts.length)"></v-skeleton-loader>
                    </div>
                </template>

                <v-card-text class="pt-4 pb-6">
                    <v-row class="asset-metrics-grid">
                        <v-col cols="12" sm="4">
                            <div class="metric-tile metric-tile-assets d-flex align-center pa-4">
                                <div class="metric-icon-box bg-emerald-subtle me-3">
                                    <v-icon size="22" color="primary" :icon="mdiBankOutline"/>
                                </div>
                                <div class="d-flex flex-column text-truncate">
                                    <span class="text-caption text-slate-muted font-weight-medium mb-1">{{ tt('Total assets') }}</span>
                                    <span class="text-h6 font-weight-bold text-truncate" v-if="!loadingOverview || (allAccounts && allAccounts.length)">{{ totalAssets }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin" width="90px" type="text" :loading="true" v-else-if="loadingOverview && (!allAccounts || !allAccounts.length)"></v-skeleton-loader>
                                </div>
                            </div>
                        </v-col>

                        <v-col cols="12" sm="4">
                            <div class="metric-tile metric-tile-liabilities d-flex align-center pa-4">
                                <div class="metric-icon-box bg-rose-subtle me-3">
                                    <v-icon size="22" color="expense" :icon="mdiCreditCardOutline"/>
                                </div>
                                <div class="d-flex flex-column text-truncate">
                                    <span class="text-caption text-slate-muted font-weight-medium mb-1">{{ tt('Total liabilities') }}</span>
                                    <span class="text-h6 font-weight-bold text-truncate" v-if="!loadingOverview || (allAccounts && allAccounts.length)">{{ totalLiabilities }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin" width="90px" type="text" :loading="true" v-else-if="loadingOverview && (!allAccounts || !allAccounts.length)"></v-skeleton-loader>
                                </div>
                            </div>
                        </v-col>

                        <v-col cols="12" sm="4">
                            <div class="metric-tile metric-tile-net d-flex align-center pa-4">
                                <div class="metric-icon-box bg-mint-subtle me-3">
                                    <v-icon size="22" color="primary" :icon="mdiPiggyBankOutline"/>
                                </div>
                                <div class="d-flex flex-column text-truncate">
                                    <span class="text-caption text-slate-muted font-weight-medium mb-1">{{ tt('Net assets') }}</span>
                                    <span class="text-h6 font-weight-bold text-truncate" v-if="!loadingOverview || (allAccounts && allAccounts.length)">{{ netAssets }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin" width="90px" type="text" :loading="true" v-else-if="loadingOverview && (!allAccounts || !allAccounts.length)"></v-skeleton-loader>
                                </div>
                            </div>
                        </v-col>
                    </v-row>
                </v-card-text>
            </v-card>
        </v-col>

        <v-col cols="12" md="6">
            <v-row>
                <v-col cols="12" md="6">
                    <income-expense-overview-card
                        :loading="loadingOverview" :disabled="loadingOverview" :icon="mdiCalendarTodayOutline"
                        :title="tt('Today')"
                        :expense-amount="transactionOverview.today && transactionOverview.today.valid ? getDisplayExpenseAmount(transactionOverview.today) : ''"
                        :income-amount="transactionOverview.today && transactionOverview.today.valid ? getDisplayIncomeAmount(transactionOverview.today) : ''"
                        :datetime="displayDateRange?.today?.displayTime || ''"
                    >
                        <template #menus>
                            <v-list-item :prepend-icon="mdiListBoxOutline" :to="`/transaction/list?${overviewStore.getTransactionListPageParams({ dateType: DateRange.Today.type })}`">
                                <v-list-item-title>{{ tt('View Details') }}</v-list-item-title>
                            </v-list-item>
                        </template>
                    </income-expense-overview-card>
                </v-col>

                <v-col cols="12" md="6">
                    <income-expense-overview-card
                        :loading="loadingOverview" :disabled="loadingOverview" :icon="mdiCalendarWeekOutline"
                        :title="tt('This Week')"
                        :expense-amount="transactionOverview.thisWeek && transactionOverview.thisWeek.valid ? getDisplayExpenseAmount(transactionOverview.thisWeek) : ''"
                        :income-amount="transactionOverview.thisWeek && transactionOverview.thisWeek.valid ? getDisplayIncomeAmount(transactionOverview.thisWeek) : ''"
                        :datetime="displayDateRange?.thisWeek?.startTime + '-' + displayDateRange?.thisWeek?.endTime"
                    >
                        <template #menus>
                            <v-list-item :prepend-icon="mdiListBoxOutline" :to="`/transaction/list?${overviewStore.getTransactionListPageParams({ dateType: DateRange.ThisWeek.type })}`">
                                <v-list-item-title>{{ tt('View Details') }}</v-list-item-title>
                            </v-list-item>
                        </template>
                    </income-expense-overview-card>
                </v-col>

                <v-col cols="12" md="6">
                    <income-expense-overview-card
                        :loading="loadingOverview" :disabled="loadingOverview" :icon="mdiCalendarMonthOutline"
                        :title="tt('This Month')"
                        :expense-amount="transactionOverview.thisMonth && transactionOverview.thisMonth.valid ? getDisplayExpenseAmount(transactionOverview.thisMonth) : ''"
                        :income-amount="transactionOverview.thisMonth && transactionOverview.thisMonth.valid ? getDisplayIncomeAmount(transactionOverview.thisMonth) : ''"
                        :datetime="displayDateRange?.thisMonth?.startTime + '-' + displayDateRange?.thisMonth?.endTime"
                    >
                        <template #menus>
                            <v-list-item :prepend-icon="mdiListBoxOutline" :to="`/transaction/list?${overviewStore.getTransactionListPageParams({ dateType: DateRange.ThisMonth.type })}`">
                                <v-list-item-title>{{ tt('View Details') }}</v-list-item-title>
                            </v-list-item>
                        </template>
                    </income-expense-overview-card>
                </v-col>

                <v-col cols="12" md="6">
                    <income-expense-overview-card
                        :loading="loadingOverview" :disabled="loadingOverview" :icon="mdiLayersTripleOutline"
                        :title="tt('This Year')"
                        :expense-amount="transactionOverview.thisYear && transactionOverview.thisYear.valid ? getDisplayExpenseAmount(transactionOverview.thisYear) : ''"
                        :income-amount="transactionOverview.thisYear && transactionOverview.thisYear.valid ? getDisplayIncomeAmount(transactionOverview.thisYear) : ''"
                        :datetime="displayDateRange?.thisYear?.displayTime || ''"
                    >
                        <template #menus>
                            <v-list-item :prepend-icon="mdiListBoxOutline" :to="`/transaction/list?${overviewStore.getTransactionListPageParams({ dateType: DateRange.ThisYear.type })}`">
                                <v-list-item-title>{{ tt('View Details') }}</v-list-item-title>
                            </v-list-item>
                        </template>
                    </income-expense-overview-card>
                </v-col>
            </v-row>
        </v-col>

        <v-col cols="12" md="6">
            <monthly-income-and-expense-card :data="monthlyIncomeAndExpenseData" :is-dark-mode="isDarkMode"
                                             :loading="loadingOverview" :disabled="loadingOverview"
                                             :enable-click-item="true" @click="clickMonthlyIncomeOrExpense" />
        </v-col>

        <!-- New Feature: Full Width Transaction Event Calendar (from User Snippet Translation) -->
        <v-col cols="12">
            <event-calendar />
        </v-col>
    </v-row>

    <snack-bar ref="snackbar" />
</template>

<script setup lang="ts">
import SnackBar from '@/components/desktop/SnackBar.vue';
import IncomeExpenseOverviewCard from './overview/cards/IncomeExpenseOverviewCard.vue';
import MonthlyIncomeAndExpenseCard, { type MonthlyIncomeAndExpenseCardClickEvent } from './overview/cards/MonthlyIncomeAndExpenseCard.vue';
import EventCalendar from '@/components/desktop/EventCalendar.vue';

import { ref, computed, useTemplateRef } from 'vue';
import { useRouter } from 'vue-router';
import { useTheme } from 'vuetify';

import { useI18n } from '@/locales/helpers.ts';
import { useHomePageBase } from '@/views/base/HomePageBase.ts';

import { useAccountsStore } from '@/stores/account.ts';
import { useTransactionCategoriesStore } from '@/stores/transactionCategory.ts';
import { useOverviewStore } from '@/stores/overview.ts';

import { type NumeralSystem } from '@/core/numeral.ts';
import { DateRange } from '@/core/datetime.ts';
import { ThemeType } from '@/core/theme.ts';
import {
    type TransactionMonthlyIncomeAndExpenseData,
    LATEST_12MONTHS_TRANSACTION_AMOUNTS_REQUEST_TYPES
} from '@/models/transaction.ts';

import { getUnixTimeBeforeUnixTime, getUnixTimeAfterUnixTime } from '@/lib/datetime.ts';
import { isUserLogined, isUserUnlocked } from '@/lib/userstate.ts';

import {
    mdiRefresh,
    mdiEyeOutline,
    mdiEyeOffOutline,
    mdiBankOutline,
    mdiCreditCardOutline,
    mdiPiggyBankOutline,
    mdiCalendarTodayOutline,
    mdiCalendarWeekOutline,
    mdiCalendarMonthOutline,
    mdiLayersTripleOutline,
    mdiListBoxOutline,
    mdiCalendarOutline,
    mdiArrowUp
} from '@mdi/js';

type SnackBarType = InstanceType<typeof SnackBar>;

const router = useRouter();
const theme = useTheme();

const { tt, getCurrentNumeralSystemType } = useI18n();
const {
    showAmountInHomePage,
    allAccounts,
    netAssets,
    totalAssets,
    totalLiabilities,
    displayDateRange,
    transactionOverview,
    getDisplayIncomeAmount,
    getDisplayExpenseAmount
} = useHomePageBase();

const accountsStore = useAccountsStore();
const transactionCategoriesStore = useTransactionCategoriesStore();
const overviewStore = useOverviewStore();

const snackbar = useTemplateRef<SnackBarType>('snackbar');

const loadingOverview = ref<boolean>(true);

const isDarkMode = computed<boolean>(() => theme.global.name.value === ThemeType.Dark);
const numeralSystem = computed<NumeralSystem>(() => getCurrentNumeralSystemType());

const displayAccountCount = computed<string>(() => allAccounts.value ? numeralSystem.value.formatNumber(allAccounts.value.length) : numeralSystem.value.digitZero);

function clickMonthlyIncomeOrExpense(e: MonthlyIncomeAndExpenseCardClickEvent): void {
    const minTime = e.monthStartTime;
    const maxTime = getUnixTimeBeforeUnixTime(getUnixTimeAfterUnixTime(minTime, 1, 'months'), 1, 'seconds');
    const type = e.transactionType;

    router.push(`/transaction/list?${overviewStore.getTransactionListPageParams({
        type: type,
        dateType: DateRange.Custom.type,
        minTime: minTime,
        maxTime: maxTime
    })}`);
}

const monthlyIncomeAndExpenseData = computed<TransactionMonthlyIncomeAndExpenseData[]>(() => {
    const data: TransactionMonthlyIncomeAndExpenseData[] = [];

    if (!transactionOverview.value || !transactionOverview.value.thisMonth || !transactionOverview.value.thisMonth.valid) {
        return data;
    }

    for (const amountRequestType of LATEST_12MONTHS_TRANSACTION_AMOUNTS_REQUEST_TYPES) {
        const dateRange = overviewStore.transactionDataRange[amountRequestType];

        if (!dateRange) {
            continue;
        }

        const item = transactionOverview.value[amountRequestType];

        data.push({
            monthStartTime: dateRange.startTime,
            incomeAmount: item?.incomeAmount || 0,
            expenseAmount: item?.expenseAmount || 0,
            incompleteIncomeAmount: item ? item.incompleteIncomeAmount : true,
            incompleteExpenseAmount: item ? item.incompleteExpenseAmount : true
        });
    }

    return data;
});

function reload(force: boolean): void {
    loadingOverview.value = true;

    const promises = [
        accountsStore.loadAllAccounts({ force: false }),
        transactionCategoriesStore.loadAllCategories({ force: false }),
        overviewStore.loadTransactionOverview({ force: force, loadLast11Months: true })
    ];

    Promise.all(promises).then(() => {
        loadingOverview.value = false;

        if (force) {
            snackbar.value?.showMessage('Data has been updated');
        }
    }).catch(error => {
        loadingOverview.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

if (isUserLogined() && isUserUnlocked()) {
    reload(false);
}
</script>

<!-- <style>
.overview-card-background {
    position: absolute;
    inline-size: 9rem;
    inset-block-end: 0;
    inset-inline-end: 0;
}

.overview-card-background-image {
    position: absolute;
    inline-size: 5rem;
    inset-block-end: 0.5rem;
    inset-inline-end: 1rem;
}
</style> -->
<style>
/* =========================================================
   SPILL CHILL EXECUTIVE HOME OVERVIEW
   Clean • Minimal • No Animations • No Flashy Effects
   ========================================================= */

.executive-hero-card {
    border-radius: 16px !important;
}

.overview-badge-pill {
    display: inline-flex;
    align-items: center;
    background: rgba(16, 185, 129, 0.1);
    border: 1px solid rgba(16, 185, 129, 0.22);
    border-radius: 9999px;
    padding: 3px 10px;
    font-size: 0.75rem;
    font-weight: 600;
    color: #10B981;
}

.overview-main-amount {
    font-family: 'Outfit', sans-serif !important;
    font-size: 2.15rem;
    font-weight: 700;
    letter-spacing: -0.02em;
    color: rgb(var(--v-theme-on-surface));
    line-height: 1.1;
}

.overview-income-box {
    background: rgba(16, 185, 129, 0.05);
    border: 1px solid rgba(16, 185, 129, 0.14);
    border-radius: 10px;
    padding: 6px 12px;
    display: inline-flex;
    align-items: center;
}

.overview-income-indicator {
    width: 22px;
    height: 22px;
    border-radius: 50%;
    background: rgba(16, 185, 129, 0.12);
    display: flex;
    align-items: center;
    justify-content: center;
}

.text-slate-muted {
    color: rgba(var(--v-theme-on-surface), 0.6) !important;
}

.text-emerald {
    color: #10B981 !important;
}

/* ---------------------------------------------------------
   Asset Summary Metric Tiles
   --------------------------------------------------------- */

.metric-tile {
    background: rgb(var(--v-theme-surface));
    border: 1px solid rgba(var(--v-border-color), var(--v-border-opacity));
    border-radius: 14px;
    transition: border-color 0.15s ease;
}

.metric-icon-box {
    width: 44px;
    height: 44px;
    border-radius: 11px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
}

.bg-emerald-subtle {
    background: rgba(16, 185, 129, 0.1) !important;
}

.bg-rose-subtle {
    background: rgba(244, 63, 94, 0.1) !important;
}

.bg-mint-subtle {
    background: rgba(52, 211, 153, 0.12) !important;
}
</style>
