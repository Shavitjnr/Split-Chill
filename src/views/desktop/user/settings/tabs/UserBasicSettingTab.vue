<template>
    <div class="user-basic-settings-container">
        <v-card class="user-settings-card" :class="{ 'disabled': loading || saving }">
            <!-- Card Hero Header -->
            <div class="card-hero-header d-flex align-center justify-space-between flex-wrap gap-2">
                <div class="d-flex align-center">
                    <span class="settings-card-title">{{ tt('Basic Settings') }}</span>
                    <v-progress-circular indeterminate size="20" class="ms-3 text-primary" v-if="loading"></v-progress-circular>
                </div>
                <div class="settings-card-badge">
                    <span>{{ tt('Profile & Preferences') }}</span>
                </div>
            </div>

            <!-- Profile Hero Banner -->
            <div class="profile-hero-section">
                <div class="d-flex flex-column flex-sm-row align-center align-sm-start gap-5">
                    <div class="avatar-ring-container">
                        <v-avatar rounded="xl" size="104" class="user-profile-avatar-icon"
                                  :class="{ 'cursor-pointer': avatarProvider === 'internal', 'user-profile-avatar-icon-modifiable': avatarProvider === 'internal' }"
                                  :color="currentUserAvatar ? 'rgba(0,0,0,0)' : 'primary'">
                            <v-img :src="currentUserAvatar" v-if="currentUserAvatar" cover>
                                <template #placeholder>
                                    <div class="d-flex align-center justify-center fill-height bg-light-primary">
                                        <v-icon color="primary" size="48" class="user-profile-avatar-placeholder" :icon="mdiAccount"/>
                                    </div>
                                </template>
                            </v-img>
                            <v-icon size="52" class="user-profile-avatar-placeholder" :icon="mdiAccount" v-else-if="!currentUserAvatar"/>
                            <div class="avatar-edit-overlay" v-if="avatarProvider === 'internal'">
                                <v-icon size="26" :icon="mdiCameraOutline"/>
                                <span class="edit-badge-text">{{ tt('Change') }}</span>
                            </div>
                            <v-menu activator="parent" width="200" location="bottom" offset="12px" v-if="avatarProvider === 'internal'">
                                <v-list class="avatar-menu-list">
                                    <v-list-item :disabled="saving" :prepend-icon="mdiAccountEditOutline" :title="tt('Update Avatar')" @click="showOpenAvatarDialog"></v-list-item>
                                    <v-list-item :disabled="!currentUserAvatar || saving" :title="tt('Remove Avatar')" @click="removeAvatar"></v-list-item>
                                </v-list>
                            </v-menu>
                        </v-avatar>
                    </div>

                    <div class="profile-meta-info d-flex flex-column justify-center flex-grow-1 text-center text-sm-start">
                        <div class="d-flex flex-column flex-sm-row align-center align-sm-baseline gap-2 mb-2">
                            <span class="user-hero-username" v-if="!loading">{{ oldProfile.username }}</span>
                            <v-skeleton-loader class="skeleton-no-margin" type="text" style="width: 140px" :loading="true" v-if="loading"></v-skeleton-loader>
                            <span class="user-role-chip" v-if="!loading">{{ tt('Active Account') }}</span>
                        </div>

                        <div class="email-status-row d-flex flex-wrap align-center justify-center justify-sm-start gap-2">
                            <div class="status-pill status-pill-verified" v-if="!loading && emailVerified">
                                <v-icon size="16" :icon="mdiCheckDecagram"/>
                                <span>{{ tt('Email address is verified') }}</span>
                            </div>
                            <div class="status-pill status-pill-unverified" v-if="!loading && !emailVerified">
                                <v-icon size="16" :icon="mdiAlertCircleOutline"/>
                                <span>{{ tt('Email address is not verified') }}</span>
                            </div>
                            <v-btn class="resend-verification-btn" size="small" variant="tonal" :disabled="loading || resending"
                                   @click="resendVerifyEmail" v-if="isUserVerifyEmailEnabled() && !loading && !emailVerified">
                                {{ tt('Resend Validation Email') }}
                                <v-progress-circular indeterminate size="16" class="ms-2" v-if="resending"></v-progress-circular>
                            </v-btn>
                            <v-skeleton-loader class="skeleton-no-margin" type="text" style="width: 180px" :loading="true" v-if="loading"></v-skeleton-loader>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Form Content -->
            <v-form class="settings-form-body">
                <!-- Section 1: Profile & Account -->
                <div class="settings-group-section">
                    <div class="section-heading-bar d-flex align-center gap-3 mb-5">
                        <div class="section-icon-badge">
                            <v-icon size="20" :icon="mdiAccountOutline" />
                        </div>
                        <div>
                            <h3 class="section-title">{{ tt('Profile & Account') }}</h3>
                            <p class="section-subtitle">{{ tt('Personal details and transaction editing bounds') }}</p>
                        </div>
                    </div>

                    <v-row class="settings-input-grid">
                        <v-col cols="12" md="6">
                            <v-text-field
                                class="modern-field"
                                type="text"
                                autocomplete="nickname"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Nickname')"
                                :placeholder="tt('Your nickname')"
                                v-model="newProfile.nickname"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-text-field
                                class="modern-field"
                                type="email"
                                autocomplete="email"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('E-mail')"
                                :placeholder="tt('Your email address')"
                                v-model="newProfile.email"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <two-column-select primary-key-field="id" primary-value-field="category"
                                               primary-title-field="name"
                                               primary-icon-field="icon" primary-icon-type="account"
                                               primary-sub-items-field="accounts"
                                               :primary-title-i18n="true"
                                               secondary-key-field="id" secondary-value-field="id"
                                               secondary-title-field="name"
                                               secondary-icon-field="icon" secondary-icon-type="account" secondary-color-field="color"
                                               :disabled="loading || saving || !allVisibleAccounts.length"
                                               :enable-filter="true" :filter-placeholder="tt('Find account')" :filter-no-items-text="tt('No available account')"
                                               :label="tt('Default Account')"
                                               :placeholder="tt('Default Account')"
                                               :items="allVisibleCategorizedAccounts"
                                               :no-item-text="Account.findAccountNameById(allAccounts, newProfile.defaultAccountId, tt('Unspecified'))"
                                               v-model="newProfile.defaultAccountId">
                            </two-column-select>
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Editable Transaction Range')"
                                :placeholder="tt('Editable Transaction Range')"
                                :items="allTransactionEditScopeTypes"
                                v-model="newProfile.transactionEditScope"
                            />
                        </v-col>
                    </v-row>
                </div>

                <!-- Section 2: Regional & Currency -->
                <div class="settings-group-section">
                    <div class="section-heading-bar d-flex align-center gap-3 mb-5">
                        <div class="section-icon-badge">
                            <v-icon size="20" :icon="mdiTranslate" />
                        </div>
                        <div>
                            <h3 class="section-title">{{ tt('Regional & Currency') }}</h3>
                            <p class="section-subtitle">{{ tt('Primary currency, language, and calendar week cycle') }}</p>
                        </div>
                    </div>

                    <v-row class="settings-input-grid">
                        <v-col cols="12" md="6">
                            <language-select :disabled="loading || saving"
                                             :label="languageTitle"
                                             :placeholder="languageTitle"
                                             :include-system-default="true"
                                             :use-model-value="true" v-model="newProfile.language" />
                        </v-col>

                        <v-col cols="12" md="6">
                            <currency-select :disabled="loading || saving"
                                             :label="tt('Default Currency')"
                                             :placeholder="tt('Default Currency')"
                                             v-model="newProfile.defaultCurrency" />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('First Day of Week')"
                                :placeholder="tt('First Day of Week')"
                                :items="allWeekDays"
                                v-model="newProfile.firstDayOfWeek"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <fiscal-year-start-select
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Fiscal Year Start Date')"
                                :placeholder="tt('Fiscal Year Start Date')"
                                :numeral-system="newProfile.numeralSystem"
                                v-model="newProfile.fiscalYearStart"
                            />
                        </v-col>
                    </v-row>
                </div>

                <!-- Section 3: Date & Time Formats -->
                <div class="settings-group-section">
                    <div class="section-heading-bar d-flex align-center gap-3 mb-5">
                        <div class="section-icon-badge">
                            <v-icon size="20" :icon="mdiCalendarClockOutline" />
                        </div>
                        <div>
                            <h3 class="section-title">{{ tt('Date & Time Formats') }}</h3>
                            <p class="section-subtitle">{{ tt('Presentation styles for dates, times, and fiscal periods') }}</p>
                        </div>
                    </div>

                    <v-row class="settings-input-grid">
                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Calendar Display Type')"
                                :placeholder="tt('Calendar Display Type')"
                                :items="allCalendarDisplayTypes"
                                v-model="newProfile.calendarDisplayType"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Date Display Type')"
                                :placeholder="tt('Date Display Type')"
                                :items="allDateDisplayTypes"
                                v-model="newProfile.dateDisplayType"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Long Date Format')"
                                :placeholder="tt('Long Date Format')"
                                :items="allLongDateFormats"
                                v-model="newProfile.longDateFormat"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Short Date Format')"
                                :placeholder="tt('Short Date Format')"
                                :items="allShortDateFormats"
                                v-model="newProfile.shortDateFormat"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Long Time Format')"
                                :placeholder="tt('Long Time Format')"
                                :items="allLongTimeFormats"
                                v-model="newProfile.longTimeFormat"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Short Time Format')"
                                :placeholder="tt('Short Time Format')"
                                :items="allShortTimeFormats"
                                v-model="newProfile.shortTimeFormat"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Fiscal Year Format')"
                                :placeholder="tt('Fiscal Year Format')"
                                :items="allFiscalYearFormats"
                                v-model="newProfile.fiscalYearFormat"
                            />
                        </v-col>
                    </v-row>
                </div>

                <!-- Section 4: Numbers & Currency -->
                <div class="settings-group-section">
                    <div class="section-heading-bar d-flex align-center gap-3 mb-5">
                        <div class="section-icon-badge">
                            <v-icon size="20" :icon="mdiFormatListNumbered" />
                        </div>
                        <div>
                            <h3 class="section-title">{{ tt('Numbers & Currency Styling') }}</h3>
                            <p class="section-subtitle">{{ tt('Grouping delimiters, decimal points, and display modes') }}</p>
                        </div>
                    </div>

                    <v-row class="settings-input-grid">
                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Currency Display Mode')"
                                :placeholder="tt('Currency Display Mode')"
                                :items="allCurrencyDisplayTypes"
                                v-model="newProfile.currencyDisplayType"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Numeral System')"
                                :placeholder="tt('Numeral System')"
                                :items="allNumeralSystemTypes"
                                v-model="newProfile.numeralSystem"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Digit Grouping')"
                                :placeholder="tt('Digit Grouping')"
                                :items="allDigitGroupingTypes"
                                v-model="newProfile.digitGrouping"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving || !supportDigitGroupingSymbol"
                                :label="tt('Digit Grouping Symbol')"
                                :placeholder="tt('Digit Grouping Symbol')"
                                :items="allDigitGroupingSymbols"
                                v-model="newProfile.digitGroupingSymbol"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Decimal Separator')"
                                :placeholder="tt('Decimal Separator')"
                                :items="allDecimalSeparators"
                                v-model="newProfile.decimalSeparator"
                            />
                        </v-col>
                    </v-row>
                </div>

                <!-- Section 5: Geographic Location -->
                <div class="settings-group-section">
                    <div class="section-heading-bar d-flex align-center gap-3 mb-5">
                        <div class="section-icon-badge">
                            <v-icon size="20" :icon="mdiMapMarkerOutline" />
                        </div>
                        <div>
                            <h3 class="section-title">{{ tt('Geographic Location') }}</h3>
                            <p class="section-subtitle">{{ tt('Formatting format for coordinates and maps') }}</p>
                        </div>
                    </div>

                    <v-row class="settings-input-grid">
                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Geographic Location Format')"
                                :placeholder="tt('Geographic Location Format')"
                                :items="allCoordinateDisplayTypes"
                                v-model="newProfile.coordinateDisplayType"
                            />
                        </v-col>
                    </v-row>
                </div>

                <!-- Section 6: Amount Colors -->
                <div class="settings-group-section">
                    <div class="section-heading-bar d-flex align-center gap-3 mb-5">
                        <div class="section-icon-badge">
                            <v-icon size="20" :icon="mdiPaletteOutline" />
                        </div>
                        <div>
                            <h3 class="section-title">{{ tt('Amount Visuals') }}</h3>
                            <p class="section-subtitle">{{ tt('Custom color coding for expenses and incomes') }}</p>
                        </div>
                    </div>

                    <v-row class="settings-input-grid">
                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Expense Amount Color')"
                                :placeholder="tt('Expense Amount Color')"
                                :items="allExpenseAmountColorTypes"
                                v-model="newProfile.expenseAmountColor"
                            />
                        </v-col>

                        <v-col cols="12" md="6">
                            <v-select
                                class="modern-field"
                                item-title="displayName"
                                item-value="type"
                                persistent-placeholder
                                :disabled="loading || saving"
                                :label="tt('Income Amount Color')"
                                :placeholder="tt('Income Amount Color')"
                                :items="allIncomeAmountColorTypes"
                                v-model="newProfile.incomeAmountColor"
                            />
                        </v-col>
                    </v-row>
                </div>

                <!-- Action Buttons Footer -->
                <div class="settings-actions-footer d-flex flex-wrap align-center gap-4 pt-4">
                    <v-btn
                        class="settings-save-btn"
                        size="large"
                        :disabled="inputIsNotChanged || inputIsInvalid || saving"
                        @click="save"
                    >
                        <v-icon :icon="mdiContentSaveCheckOutline" class="me-2" size="20" v-if="!saving"/>
                        {{ tt('Save Changes') }}
                        <v-progress-circular indeterminate size="20" class="ms-2" v-if="saving"></v-progress-circular>
                    </v-btn>

                    <v-btn
                        class="settings-reset-btn"
                        size="large"
                        variant="tonal"
                        @click="reset"
                    >
                        <v-icon :icon="mdiRestart" class="me-2" size="20"/>
                        {{ tt('Reset') }}
                    </v-btn>
                </div>
            </v-form>
        </v-card>

        <confirm-dialog ref="confirmDialog"/>
        <snack-bar ref="snackbar" />
        <input ref="avatarInput" type="file" style="display: none" :accept="SUPPORTED_IMAGE_EXTENSIONS" @change="updateAvatar($event)" />
    </div>
</template>

<script setup lang="ts">
import ConfirmDialog from '@/components/desktop/ConfirmDialog.vue';
import SnackBar from '@/components/desktop/SnackBar.vue';

import { ref, computed, useTemplateRef } from 'vue';

import { useI18n } from '@/locales/helpers.ts';
import { useUserProfilePageBase } from '@/views/base/users/UserProfilePageBase.ts';

import { useRootStore } from '@/stores/index.ts';
import { useUserStore } from '@/stores/user.ts';
import { useAccountsStore } from '@/stores/account.ts';

import { SUPPORTED_IMAGE_EXTENSIONS } from '@/consts/file.ts';
import type { UserProfileResponse } from '@/models/user.ts';
import { Account } from '@/models/account.ts';

import { generateRandomUUID } from '@/lib/misc.ts';
import { isUserVerifyEmailEnabled } from '@/lib/server_settings.ts';

import {
    mdiAccount,
    mdiAccountOutline,
    mdiAccountEditOutline,
    mdiCameraOutline,
    mdiTranslate,
    mdiCalendarClockOutline,
    mdiFormatListNumbered,
    mdiMapMarkerOutline,
    mdiPaletteOutline,
    mdiCheckDecagram,
    mdiAlertCircleOutline,
    mdiContentSaveCheckOutline,
    mdiRestart
} from '@mdi/js';

type ConfirmDialogType = InstanceType<typeof ConfirmDialog>;
type SnackBarType = InstanceType<typeof SnackBar>;

const { tt } = useI18n();

const {
    newProfile,
    oldProfile,
    emailVerified,
    loading,
    resending,
    saving,
    allAccounts,
    allVisibleAccounts,
    allVisibleCategorizedAccounts,
    allWeekDays,
    allCalendarDisplayTypes,
    allDateDisplayTypes,
    allLongDateFormats,
    allShortDateFormats,
    allLongTimeFormats,
    allShortTimeFormats,
    allFiscalYearFormats,
    allCurrencyDisplayTypes,
    allNumeralSystemTypes,
    allDecimalSeparators,
    allDigitGroupingSymbols,
    allDigitGroupingTypes,
    allCoordinateDisplayTypes,
    allExpenseAmountColorTypes,
    allIncomeAmountColorTypes,
    allTransactionEditScopeTypes,
    languageTitle,
    supportDigitGroupingSymbol,
    inputIsNotChangedProblemMessage,
    inputInvalidProblemMessage,
    langAndRegionInputInvalidProblemMessage,
    extendInputInvalidProblemMessage,
    inputIsNotChanged,
    inputIsInvalid,
    setCurrentUserProfile,
    reset,
    doAfterProfileUpdate
} = useUserProfilePageBase();

const rootStore = useRootStore();
const userStore = useUserStore();
const accountsStore = useAccountsStore();

const confirmDialog = useTemplateRef<ConfirmDialogType>('confirmDialog');
const snackbar = useTemplateRef<SnackBarType>('snackbar');
const avatarInput = useTemplateRef<HTMLInputElement>('avatarInput');

const avatarUrl = ref<string>('');
const avatarProvider = ref<string | undefined>('');
const avatarNoCacheId = ref<string>('');

const currentUserAvatar = computed<string | null>(() => userStore.getUserAvatarUrl(avatarUrl.value, avatarNoCacheId.value));

function init(): void {
    loading.value = true;

    const promises = [
        accountsStore.loadAllAccounts({ force: false }),
        userStore.getCurrentUserProfile()
    ];

    Promise.all(promises).then(responses => {
        const profile = responses[1] as UserProfileResponse;
        setCurrentUserProfile(profile);
        avatarUrl.value = profile.avatar;
        avatarProvider.value = profile.avatarProvider;
        loading.value = false;
    }).catch(error => {
        oldProfile.value.nickname = '';
        oldProfile.value.email = '';
        newProfile.value.nickname = '';
        newProfile.value.email = '';
        loading.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function save(): void {
    const problemMessage = inputIsNotChangedProblemMessage.value || inputInvalidProblemMessage.value || extendInputInvalidProblemMessage.value || langAndRegionInputInvalidProblemMessage.value;

    if (problemMessage) {
        snackbar.value?.showMessage(problemMessage);
        return;
    }

    saving.value = true;

    rootStore.updateUserProfile(newProfile.value.toProfileUpdateRequest()).then(response => {
        saving.value = false;

        doAfterProfileUpdate(response.user);
        snackbar.value?.showMessage('Your profile has been successfully updated');
    }).catch(error => {
        saving.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function updateAvatar(event: Event): void {
    if (!event || !event.target) {
        return;
    }

    const el = event.target as HTMLInputElement;

    if (!el.files || !el.files.length || !el.files[0]) {
        return;
    }

    const avatarFile = el.files[0] as File;

    el.value = '';

    saving.value = true;

    userStore.updateUserAvatar({ avatarFile }).then(response => {
        saving.value = false;

        if (response) {
            avatarUrl.value = response.avatar;
            avatarProvider.value = response.avatarProvider;
            avatarNoCacheId.value = generateRandomUUID();
            setCurrentUserProfile(response);
        }

        snackbar.value?.showMessage('Your avatar has been successfully updated');
    }).catch(error => {
        saving.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function removeAvatar(): void {
    confirmDialog.value?.open('Are you sure you want to remove avatar?').then(() => {
        saving.value = true;

        userStore.removeUserAvatar().then(response => {
            saving.value = false;

            if (response) {
                avatarUrl.value = response.avatar;
                avatarProvider.value = response.avatarProvider;
                setCurrentUserProfile(response);
            }

            snackbar.value?.showMessage('Your profile has been successfully updated');
        }).catch(error => {
            saving.value = false;

            if (!error.processed) {
                snackbar.value?.showError(error);
            }
        });
    });
}

function resendVerifyEmail(): void {
    resending.value = true;

    rootStore.resendVerifyEmailByLoginedUser().then(() => {
        resending.value = false;
        snackbar.value?.showMessage('Validation email has been sent');
    }).catch(error => {
        resending.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function showOpenAvatarDialog(): void {
    avatarInput.value?.click();
}

init();
</script>

<style lang="scss" scoped>
/* =========================================================
   PREMIUM USER BASIC SETTINGS
   Clean • Modern SaaS • Luxury Purple Theme
   ========================================================= */

.user-basic-settings-container {
    width: 100%;
}

.user-settings-card {
    position: relative;
    border-radius: 20px !important;
    background: rgba(var(--v-theme-surface), 0.85) !important;
    backdrop-filter: blur(16px);
    border: 1px solid rgba(var(--v-border-color), 0.12) !important;
    box-shadow: 0 10px 30px -5px rgba(15, 23, 42, 0.05), 0 4px 12px -2px rgba(15, 23, 42, 0.02) !important;
    overflow: hidden;
    transition: box-shadow 0.25s ease;

    &.disabled {
        pointer-events: none;
        opacity: 0.85;
    }
}

/* Card Hero Header */
.card-hero-header {
    padding: 24px 28px 20px;
    border-bottom: 1px solid rgba(var(--v-border-color), 0.08);

    .settings-card-title {
        font-family: 'Outfit', sans-serif;
        font-size: 20px;
        font-weight: 700;
        letter-spacing: -0.01em;
        color: rgb(var(--v-theme-on-surface));
    }

    .settings-card-badge {
        font-family: 'Inter', sans-serif;
        font-size: 12px;
        font-weight: 600;
        padding: 4px 12px;
        border-radius: 999px;
        background: rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.09);
        color: rgb(var(--v-theme-primary));
        border: 1px solid rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.2);
    }
}

/* Profile Hero Section */
.profile-hero-section {
    padding: 28px;
    background:
        radial-gradient(circle at 10% 20%, rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.08) 0%, transparent 50%),
        linear-gradient(135deg, rgba(var(--v-theme-surface), 0.6) 0%, rgba(var(--v-theme-surface), 0.95) 100%);
    border-bottom: 1px solid rgba(var(--v-border-color), 0.08);

    .avatar-ring-container {
        position: relative;
        padding: 4px;
        border-radius: 22px;
        background: linear-gradient(135deg, rgba(16, 185, 129, 0.6), rgba(6, 78, 59, 0.4));
        box-shadow: 0 4px 18px rgba(16, 185, 129, 0.15);
        transition: transform 0.25s ease, box-shadow 0.25s ease;

        &:hover {
            transform: scale(1.02);
            box-shadow: 0 8px 24px rgba(147, 51, 234, 0.25);
        }
    }

    .user-profile-avatar-icon {
        border-radius: 18px !important;
        position: relative;
        overflow: hidden;

        .avatar-edit-overlay {
            position: absolute;
            inset: 0;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            background: rgba(15, 23, 42, 0.65);
            backdrop-filter: blur(4px);
            color: #ffffff;
            opacity: 0;
            transition: opacity 0.2s ease;
            cursor: pointer;

            .edit-badge-text {
                font-family: 'Inter', sans-serif;
                font-size: 11px;
                font-weight: 600;
                margin-top: 4px;
                text-transform: uppercase;
                letter-spacing: 0.04em;
            }
        }

        &:hover .avatar-edit-overlay {
            opacity: 1;
        }

        .user-profile-avatar-placeholder {
            opacity: 0.8;
        }
    }

    .user-hero-username {
        font-family: 'Outfit', sans-serif;
        font-size: 24px;
        font-weight: 700;
        letter-spacing: -0.02em;
        color: rgb(var(--v-theme-on-surface));
    }

    .user-role-chip {
        font-family: 'Inter', sans-serif;
        font-size: 11px;
        font-weight: 600;
        padding: 2px 10px;
        border-radius: 6px;
        background: rgba(var(--v-theme-on-surface), 0.06);
        color: rgba(var(--v-theme-on-surface), 0.7);
        text-transform: uppercase;
        letter-spacing: 0.05em;
    }

    .status-pill {
        display: inline-flex;
        align-items: center;
        gap: 6px;
        font-family: 'Inter', sans-serif;
        font-size: 13px;
        font-weight: 600;
        padding: 5px 14px;
        border-radius: 999px;

        &.status-pill-verified {
            background: rgba(16, 185, 129, 0.1);
            color: #10b981;
            border: 1px solid rgba(16, 185, 129, 0.25);
        }

        &.status-pill-unverified {
            background: rgba(245, 158, 11, 0.1);
            color: #f59e0b;
            border: 1px solid rgba(245, 158, 11, 0.25);
        }
    }

    .resend-verification-btn {
        border-radius: 999px !important;
        font-family: 'Inter', sans-serif;
        font-size: 12px !important;
        font-weight: 600;
        text-transform: none;
        letter-spacing: 0;
        border: 1px solid rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.2);
    }
}

/* Form Sections */
.settings-form-body {
    padding: 28px;
}

.settings-group-section {
    padding: 24px;
    margin-bottom: 24px;
    border-radius: 16px;
    background: rgba(var(--v-theme-surface), 0.45);
    border: 1px solid rgba(var(--v-border-color), 0.08);
    transition: border-color 0.2s ease, background 0.2s ease;

    &:hover {
        border-color: rgba(var(--v-border-color), 0.16);
        background: rgba(var(--v-theme-surface), 0.65);
    }

    .section-heading-bar {
        .section-icon-badge {
            width: 38px;
            height: 38px;
            border-radius: 10px;
            display: flex;
            align-items: center;
            justify-content: center;
            background: rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.1);
            color: rgb(var(--v-theme-primary));
            border: 1px solid rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.15);
        }

        .section-title {
            font-family: 'Outfit', sans-serif;
            font-size: 17px;
            font-weight: 600;
            color: rgb(var(--v-theme-on-surface));
            margin: 0;
            line-height: 1.25;
        }

        .section-subtitle {
            font-family: 'Inter', sans-serif;
            font-size: 13px;
            color: rgba(var(--v-theme-on-surface), 0.6);
            margin: 2px 0 0;
        }
    }
}

/* Fields Styling */
:deep(.v-field) {
    border-radius: 12px !important;
    background: rgba(var(--v-theme-on-surface), 0.02) !important;
    border: 1px solid rgba(var(--v-border-color), 0.12) !important;
    transition: all 0.2s ease;

    &:hover {
        border-color: rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.3) !important;
        background: rgba(var(--v-theme-on-surface), 0.035) !important;
    }

    &.v-field--focused {
        border-color: rgb(var(--v-theme-primary)) !important;
        background: rgb(var(--v-theme-surface)) !important;
        box-shadow: 0 0 0 3px rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.12) !important;
    }
}

:deep(.v-label) {
    font-family: 'Inter', sans-serif;
    font-size: 14px;
    font-weight: 500;
    opacity: 0.75;
}

:deep(input),
:deep(.v-select__selection-text) {
    font-family: 'Inter', sans-serif;
    font-weight: 500;
    font-size: 14px;
}

/* Actions Footer */
.settings-actions-footer {
    border-top: 1px solid rgba(var(--v-border-color), 0.08);
    margin-top: 8px;

    .settings-save-btn {
        background: linear-gradient(135deg, #10B981 0%, #059669 100%) !important;
        color: #ffffff !important;
        border-radius: 12px !important;
        font-family: 'Inter', sans-serif;
        font-size: 14px;
        font-weight: 600;
        letter-spacing: 0;
        text-transform: none;
        padding: 0 28px !important;
        height: 48px;
        box-shadow: 0 10px 25px -5px rgba(16, 185, 129, 0.25) !important;
        transition: transform 0.2s ease, box-shadow 0.2s ease;

        &:not(:disabled):hover {
            transform: translateY(-1px);
            box-shadow: 0 6px 22px rgba(16, 185, 129, 0.45);
        }

        &:disabled {
            opacity: 0.55;
            box-shadow: none;
        }
    }

    .settings-reset-btn {
        border-radius: 12px !important;
        font-family: 'Inter', sans-serif;
        font-size: 14px;
        font-weight: 600;
        letter-spacing: 0;
        text-transform: none;
        padding: 0 24px !important;
        height: 48px;
        border: 1px solid rgba(var(--v-border-color), 0.15);
    }
}

/* Dark Mode Fine-tuning */
.v-theme--dark {
    .user-settings-card {
        background: rgba(30, 41, 59, 0.7) !important;
        border-color: rgba(255, 255, 255, 0.08) !important;
        box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3) !important;
    }

    .profile-hero-section {
        background:
            radial-gradient(circle at 10% 20%, rgba(168, 85, 247, 0.12) 0%, transparent 50%),
            linear-gradient(135deg, rgba(30, 41, 59, 0.5) 0%, rgba(15, 23, 42, 0.9) 100%);
    }

    .settings-group-section {
        background: rgba(15, 23, 42, 0.35);
        border-color: rgba(255, 255, 255, 0.05);

        &:hover {
            background: rgba(15, 23, 42, 0.5);
            border-color: rgba(255, 255, 255, 0.1);
        }
    }
}

@media (max-width: 600px) {
    .card-hero-header {
        padding: 18px 20px;
    }

    .profile-hero-section {
        padding: 20px;
    }

    .settings-form-body {
        padding: 16px;
    }

    .settings-group-section {
        padding: 16px;
        margin-bottom: 16px;
    }

    .settings-actions-footer {
        flex-direction: column;
        align-items: stretch;

        .settings-save-btn,
        .settings-reset-btn {
            width: 100%;
        }
    }
}
</style>
