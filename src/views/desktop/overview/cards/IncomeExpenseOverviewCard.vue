<template>
    <v-card class="period-overview-card" :class="{ 'disabled': disabled }">
        <v-card-text class="d-flex align-center pb-2">
            <div class="period-icon-box bg-emerald-subtle me-3">
                <v-icon size="20" color="primary" :icon="icon" />
            </div>
            <span class="font-weight-bold text-subtitle-1">{{ title }}</span>
            <v-spacer/>
            <v-btn density="comfortable" color="default" variant="text" class="ms-2" :icon="true">
                <v-icon :icon="mdiDotsVertical" size="18" />
                <v-menu activator="parent">
                    <v-list>
                        <slot name="menus"></slot>
                    </v-list>
                </v-menu>
            </v-btn>
        </v-card-text>

        <v-card-text class="pt-2 pb-2">
            <template v-if="loading && (!incomeAmount && !expenseAmount)">
                <v-skeleton-loader class="skeleton-no-margin mb-2" type="text" width="100%" :loading="true"></v-skeleton-loader>
                <v-skeleton-loader class="skeleton-no-margin" type="text" width="80%" :loading="true"></v-skeleton-loader>
            </template>
            <template v-else-if="incomeAmount || expenseAmount">
                <div class="d-flex align-center justify-space-between mb-2" v-if="incomeAmount">
                    <span class="metric-label-pill income-pill">
                        <v-icon :icon="mdiArrowUp" size="13" class="me-1" />
                        <span>{{ tt('Income') }}</span>
                    </span>
                    <span class="text-subtitle-1 font-weight-bold text-income text-truncate ms-2">{{ incomeAmount }}</span>
                </div>
                <div class="d-flex align-center justify-space-between" v-if="expenseAmount">
                    <span class="metric-label-pill expense-pill">
                        <v-icon :icon="mdiArrowDown" size="13" class="me-1" />
                        <span>{{ tt('Expense') }}</span>
                    </span>
                    <span class="text-subtitle-1 font-weight-bold text-expense text-truncate ms-2">{{ expenseAmount }}</span>
                </div>
            </template>
            <template v-else>
                <div class="d-flex align-center justify-center py-3 text-slate-muted text-caption font-weight-medium">
                    <span>{{ tt('No data') }}</span>
                </div>
            </template>
        </v-card-text>

        <v-card-text class="pt-2 pb-4">
            <div class="period-date-pill">
                <v-icon :icon="mdiClockOutline" size="13" class="me-1" />
                <span class="text-truncate">{{ datetime }}</span>
            </div>
        </v-card-text>
    </v-card>
</template>

<script setup lang="ts">
import { useI18n } from '@/locales/helpers.ts';

import {
    mdiDotsVertical,
    mdiArrowUp,
    mdiArrowDown,
    mdiClockOutline
} from '@mdi/js';

defineProps<{
    loading: boolean;
    disabled: boolean;
    icon: string;
    title: string;
    expenseAmount: string;
    incomeAmount: string;
    datetime: string;
}>();

const { tt } = useI18n();
</script>

<style scoped>
.period-overview-card {
    border-radius: 16px !important;
    transition: border-color 0.15s ease;
}

.period-icon-box {
    width: 38px;
    height: 38px;
    border-radius: 10px;
    background: rgba(16, 185, 129, 0.1);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
}

.metric-label-pill {
    display: inline-flex;
    align-items: center;
    padding: 2px 8px;
    border-radius: 6px;
    font-size: 0.72rem;
    font-weight: 600;
}

.income-pill {
    background: rgba(16, 185, 129, 0.1);
    color: #10B981;
}

.expense-pill {
    background: rgba(244, 63, 94, 0.1);
    color: #F43F5E;
}

.period-date-pill {
    display: inline-flex;
    align-items: center;
    padding: 3px 9px;
    background: rgba(var(--v-theme-on-surface), 0.06);
    border-radius: 9999px;
    font-size: 0.72rem;
    font-weight: 500;
    color: rgba(var(--v-theme-on-surface), 0.7);
    max-width: 100%;
}

.text-slate-muted {
    color: rgba(var(--v-theme-on-surface), 0.6) !important;
}
</style>
