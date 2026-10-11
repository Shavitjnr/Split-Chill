<template>
    <v-select
        item-title="displayName"
        item-value="currencyCode"
        no-auto-scroll
        persistent-placeholder
        :disabled="disabled"
        :label="label"
        :placeholder="placeholder"
        :prepend-inner-icon="prependInnerIcon"
        :items="filteredCurrencies"
        :menu-props="{ maxHeight: 380, contentClass: 'currency-select-menu' }"
        v-model="currentCurrencyValue"
        v-model:menu="isMenuOpen"
        @update:menu="onMenuUpdate"
    >
        <template #selection>
            <span class="text-truncate">{{ currentCurrencyDisplayName }}</span>
        </template>

        <template #append-inner>
            <small class="text-field-append-text smaller me-1"
                   v-if="currentCurrencyValue">{{ currentCurrencyValue }}</small>
        </template>

        <template #prepend-item>
            <div class="px-3 pt-2 pb-2 currency-search-container" @click.stop @mousedown.stop>
                <v-text-field
                    ref="searchInput"
                    density="compact"
                    variant="outlined"
                    hide-details
                    single-line
                    clearable
                    :prepend-inner-icon="mdiMagnify"
                    :placeholder="tt('Search currency...')"
                    v-model="searchQuery"
                    @click.stop
                    @mousedown.stop
                    @keydown.stop
                    @keydown.esc="isMenuOpen = false"
                />
            </div>
        </template>

        <template #no-data>
            <div class="pa-4 text-center text-medium-emphasis">
                {{ tt('No results') }}
            </div>
        </template>

        <template #item="{ props: itemProps, item }">
            <v-list-item v-bind="itemProps" :active="currentCurrencyValue === item.raw.currencyCode">
                <template #title>
                    <div class="d-flex align-center">
                        <span>{{ item.raw.displayName }}</span>
                        <v-spacer style="min-width: 40px" />
                        <v-icon :icon="mdiCheck" v-if="currentCurrencyValue === item.raw.currencyCode" />
                        <small class="text-field-append-text" v-if="currentCurrencyValue !== item.raw.currencyCode">{{ item.raw.currencyCode }}</small>
                    </div>
                </template>
            </v-list-item>
        </template>
    </v-select>
</template>

<script setup lang="ts">
import { ref, computed, nextTick } from 'vue';
import { useI18n } from '@/locales/helpers.ts';
import type { LocalizedCurrencyInfo } from '@/core/currency.ts';
import {
    mdiCheck,
    mdiMagnify
} from '@mdi/js';

const props = defineProps<{
    disabled?: boolean;
    label?: string;
    placeholder?: string;
    prependInnerIcon?: string;
    modelValue?: string;
}>();

const emit = defineEmits<{
    (e: 'update:modelValue', value: string): void;
}>();

const { tt, getAllCurrencies, getCurrencyName } = useI18n();

const isMenuOpen = ref<boolean>(false);
const searchQuery = ref<string>('');
const searchInput = ref<{ focus: () => void } | null>(null);

const allCurrencies = computed<LocalizedCurrencyInfo[]>(() => getAllCurrencies());

const filteredCurrencies = computed<LocalizedCurrencyInfo[]>(() => {
    const q = searchQuery.value.trim().toLowerCase();
    if (!q) {
        return allCurrencies.value;
    }
    return allCurrencies.value.filter(currency =>
        currency.displayName.toLowerCase().includes(q) ||
        currency.currencyCode.toLowerCase().includes(q)
    );
});

const currentCurrencyDisplayName = computed<string>(() => {
    if (!props.modelValue) {
        return '';
    }
    return getCurrencyName(props.modelValue);
});

const currentCurrencyValue = computed<string>({
    get: () => props.modelValue ?? '',
    set: (value: string | null) => {
        if (value === null) {
            emit('update:modelValue', '');
        } else {
            emit('update:modelValue', value);
        }
    }
});

function onMenuUpdate(open: boolean): void {
    isMenuOpen.value = open;
    if (open) {
        searchQuery.value = '';
        nextTick(() => {
            searchInput.value?.focus();
        });
    } else {
        searchQuery.value = '';
    }
}
</script>

<style scoped>
.currency-search-container {
    position: sticky;
    top: 0;
    z-index: 10;
    background-color: rgb(var(--v-theme-surface));
    border-bottom: 1px solid rgba(var(--v-border-color), var(--v-border-opacity));
}
</style>
