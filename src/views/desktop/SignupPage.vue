<template>
    <div class="clean-signup-page">
        <div class="signup-container">
            <!-- Brand Logo & Header matching landing.html & LoginPage.vue -->
            <div class="signup-brand mb-6 text-center">
                <router-link to="/" class="brand-link">
                    <div class="brand-icon-box">
                        <v-icon icon="mdi-call-split" class="brand-icon" size="24" color="white" />
                    </div>
                    <div class="d-flex flex-column text-start">
                        <div class="d-flex align-center gap-2">
                            <h1 class="signup-brand-title">{{ tt('global.app.title') }}</h1>
                            <span class="brand-pulse-dot"></span>
                        </div>
                        <span class="brand-tagline">Expense Harmony</span>
                    </div>
                </router-link>
            </div>

            <!-- Centered Luxury Registration Card -->
            <v-card class="signup-card" elevation="0">
                <div class="d-flex justify-center mb-4">
                    <div class="signup-badge-pill">
                        <span class="badge-dot"></span>
                        <span>Seamless Onboarding • Split Chill 2.0</span>
                    </div>
                </div>

                <div class="signup-card-header text-center mb-6">
                    <h2 class="signup-title">{{ tt('Create Account') }}</h2>
                    <p class="signup-subtitle">{{ tt('Fill in your details to get started with Split Chill AI') }}</p>
                </div>

                <div class="pa-0">
                    <steps-bar :steps="allSteps" :current-step="currentStep" :min-width="0" @step:change="switchToTab" class="mb-8" />

                    <v-window class="disable-tab-transition overflow-visible" v-model="currentStep">
                        <v-form @submit.prevent>
                            <v-window-item value="basicSetting">
                                <v-row dense>
                                    <v-col cols="12" sm="6">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            type="text"
                                            autocomplete="username"
                                            :disabled="submitting || navigateToHomePage"
                                            :label="tt('Username')"
                                            :prepend-inner-icon="mdiAccountOutline"
                                            v-model="user.username"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    <v-col cols="12" md="6">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            type="text"
                                            autocomplete="nickname"
                                            :disabled="submitting || navigateToHomePage"
                                            :label="tt('Nickname')"
                                            :prepend-inner-icon="mdiFaceManOutline"
                                            v-model="user.nickname"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    <v-col cols="12">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            type="email"
                                            autocomplete="email"
                                            :disabled="submitting || navigateToHomePage"
                                            :label="tt('E-mail Address')"
                                            :prepend-inner-icon="mdiEmailOutline"
                                            v-model="user.email"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    <v-col cols="12" md="6">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            autocomplete="new-password"
                                            type="password"
                                            :disabled="submitting || navigateToHomePage"
                                            :label="tt('Password')"
                                            :prepend-inner-icon="mdiLockOutline"
                                            v-model="user.password"
                                            class="modern-input"
                                        />
                                    </v-col>
                                    <v-col cols="12" md="6">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            autocomplete="new-password"
                                            type="password"
                                            :disabled="submitting || navigateToHomePage"
                                            :label="tt('Confirm Password')"
                                            :prepend-inner-icon="mdiLockCheckOutline"
                                            v-model="user.confirmPassword"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    <v-col cols="12">
                                        <v-divider class="my-4" />
                                    </v-col>

                                    <v-col cols="12" md="6">
                                        <LanguageSelect :disabled="submitting || navigateToHomePage"
                                                       :label="languageTitle"
                                                       :prependInnerIcon="mdiTranslate"
                                                       :use-model-value="true" v-model="currentLocale" />
                                    </v-col>

                                    <v-col cols="12" md="6">
                                        <CurrencySelect :disabled="submitting || navigateToHomePage"
                                                       :label="tt('Default Currency')"
                                                       :prependInnerIcon="mdiCurrencyUsd"
                                                       v-model="user.defaultCurrency" />
                                    </v-col>
                                </v-row>
                            </v-window-item>

                            <v-window-item value="presetCategories" class="signup-preset-categories">
                                <div class="d-flex align-center justify-space-between mb-6">
                                    <div>
                                        <h4 class="text-h5 font-weight-bold mb-1">{{ tt('Transaction Categories') }}</h4>
                                        <p class="text-body-2 text-medium-emphasis">{{ tt('Recommended categories for a quick start') }}</p>
                                    </div>
                                    <v-switch :disabled="submitting || navigateToHomePage"
                                              color="primary"
                                              inset
                                              :label="tt('Use Preset')"
                                              v-model="usePresetCategories"/>
                                </div>

                                <div class="overflow-y-auto px-1 stylish-scroll" :class="{ 'disabled': !usePresetCategories || submitting || navigateToHomePage }" style="max-height: 400px">
                                    <v-row :key="categoryType" v-for="(categories, categoryType) in allPresetCategories">
                                        <v-col cols="12">
                                            <h5 class="text-subtitle-1 font-weight-bold mb-3 opacity-70">{{ getCategoryTypeName(parseInt(categoryType)) }}</h5>

                                            <v-expansion-panels class="modern-panels" variant="accordion" multiple>
                                                <v-expansion-panel :key="idx" v-for="(category, idx) in categories" elevation="0">
                                                    <v-expansion-panel-title class="py-4 px-6 expand-panel-title-with-bg">
                                                        <ItemIcon icon-type="category" :icon-id="category.icon" :color="category.color"></ItemIcon>
                                                        <span class="ms-4 font-weight-medium">{{ category.name }}</span>
                                                    </v-expansion-panel-title>
                                                    <v-expansion-panel-text v-if="category.subCategories.length">
                                                        <v-list rounded density="comfortable" class="pa-0">
                                                            <template :key="subIdx"
                                                                      v-for="(subCategory, subIdx) in category.subCategories">
                                                                <v-list-item class="py-3 px-6">
                                                                    <template #prepend>
                                                                        <ItemIcon icon-type="category" :icon-id="subCategory.icon" :color="subCategory.color"></ItemIcon>
                                                                    </template>
                                                                    <v-list-item-title class="ms-4">{{ subCategory.name }}</v-list-item-title>
                                                                </v-list-item>
                                                                <v-divider v-if="subIdx !== category.subCategories.length - 1" class="opacity-30"/>
                                                            </template>
                                                        </v-list>
                                                    </v-expansion-panel-text>
                                                </v-expansion-panel>
                                            </v-expansion-panels>
                                        </v-col>
                                    </v-row>
                                </div>
                            </v-window-item>

                            <v-window-item value="finalResult" v-if="finalResultMessage" class="text-center py-10">
                                <v-icon :icon="mdiCheckCircle" color="success" size="80" class="mb-6 animate-scale-in" />
                                <h2 class="text-h4 font-weight-bold mb-4">{{ tt('All Set!') }}</h2>
                                <p class="text-body-1 text-medium-emphasis mb-10 mx-auto" style="max-width: 500px">{{ finalResultMessage }}</p>
                            </v-window-item>
                        </v-form>
                    </v-window>

                    <div class="d-flex justify-space-between align-center mt-10 pt-6 border-t flex-wrap gap-4">
                        <v-btn variant="text" color="medium-emphasis"
                               :disabled="currentStep === 'basicSetting' || currentStep === 'finalResult' || submitting || navigateToHomePage"
                               :prepend-icon="mdiArrowLeft"
                               @click="switchToPreviousTab"
                               class="premium-btn-text">
                            {{ tt('Back') }}
                        </v-btn>

                        <div class="d-flex align-center gap-4">
                            <template v-if="currentStep === 'basicSetting'">
                                <span class="text-body-2 text-medium-emphasis me-4 d-none d-sm-inline">{{ tt('Already have an account?') }} <router-link to="/login" class="text-emerald-link font-weight-bold text-decoration-none ms-1">{{ tt('Log In') }}</router-link></span>
                                <v-btn color="primary" size="large" elevation="0"
                                       :disabled="submitting || navigateToHomePage"
                                       :append-icon="mdiArrowRight"
                                       @click="switchToNextTab"
                                       class="premium-btn px-10">
                                    {{ tt('Continue') }}
                                </v-btn>
                            </template>

                            <v-btn color="primary" size="large" elevation="0"
                                   :disabled="submitting || navigateToHomePage"
                                   class="premium-btn px-10"
                                   @click="submit"
                                   v-if="currentStep === 'presetCategories'">
                                <span>{{ tt('Complete Setup') }}</span>
                                <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="submitting"></v-progress-circular>
                            </v-btn>

                            <v-btn color="primary" size="large" elevation="0"
                                   :append-icon="mdiArrowRight"
                                   @click="navigateToLogin"
                                   class="premium-btn px-10"
                                   v-if="currentStep === 'finalResult'">
                                {{ tt('Get Started') }}
                            </v-btn>
                        </div>
                    </div>
                </div>
            </v-card>

            <!-- Trust Badge & Back to Landing link matching landing.html -->
            <div class="signup-trust-badge text-center mt-6">
                <v-icon icon="mdi-shield-check" size="16" class="text-emerald me-1" />
                <span>256-Bit Bank-Grade Encryption • Zero Ad-Tracking</span>
            </div>
            <div class="text-center mt-3">
                <a href="/landing.html" class="back-to-landing-link">
                    <v-icon icon="mdi-arrow-left" size="14" class="me-1" />
                    <span>Back to Spill Chill Home</span>
                </a>
            </div>
        </div>

        <snack-bar ref="snackbar" @update:show="onSnackbarShowStateChanged" />
    </div>
</template>

<script setup lang="ts">
import ItemIcon from '@/components/desktop/ItemIcon.vue';
import StepsBar from '@/components/desktop/StepsBar.vue';
import LanguageSelect from '@/components/desktop/LanguageSelect.vue';
import CurrencySelect from '@/components/desktop/CurrencySelect.vue';
import SnackBar from '@/components/desktop/SnackBar.vue';

import type { StepBarItem } from '@/components/desktop/StepsBar.vue';

import { ref, computed, useTemplateRef } from 'vue';
import { useRouter } from 'vue-router';

import { useI18n } from '@/locales/helpers.ts';
import { useSignupPageBase } from '@/views/base/SignupPageBase.ts';

import { useRootStore } from '@/stores/index.ts';

import type { LocalizedPresetCategory } from '@/core/category.ts';
import { APPLICATION_LOGO_PATH } from '@/consts/asset.ts';

import { categorizedArrayToPlainArray } from '@/lib/common.ts';
import { isUserLogined } from '@/lib/userstate.ts';

import {
    mdiArrowLeft,
    mdiArrowRight,
    mdiAccountOutline,
    mdiFaceManOutline,
    mdiEmailOutline,
    mdiLockOutline,
    mdiLockCheckOutline,
    mdiTranslate,
    mdiCurrencyUsd,
    mdiCheckCircle
} from '@mdi/js';


type SnackBarType = InstanceType<typeof SnackBar>;

const router = useRouter();

const { tt, getAllTransactionDefaultCategories } = useI18n();

const {
    user,
    submitting,
    languageTitle,
    currentLocale,
    inputEmptyProblemMessage,
    inputInvalidProblemMessage,
    getCategoryTypeName,
    doAfterSignupSuccess
} = useSignupPageBase();

const rootStore = useRootStore();

const snackbar = useTemplateRef<SnackBarType>('snackbar');

const currentStep = ref<string>('basicSetting');
const usePresetCategories = ref<boolean>(false);
const finalResultMessage = ref<string | null>(null);
const navigateToHomePage = ref<boolean>(false);

const allPresetCategories = computed<Record<string, LocalizedPresetCategory[]>>(() => getAllTransactionDefaultCategories(0, currentLocale.value));

const allSteps = computed<StepBarItem[]>(() => {
    const allSteps = [
        {
            name: 'basicSetting',
            title: tt('User Information'),
            subTitle: tt('Basic Information')
        },
        {
            name: 'presetCategories',
            title: tt('Transaction Categories'),
            subTitle: tt('Preset Categories')
        }
    ];

    if (finalResultMessage.value) {
        allSteps.push({
            name: 'finalResult',
            title: tt('Complete'),
            subTitle: tt('Registration Completed')
        });
    }

    return allSteps;
});

function switchToTab(tabName: string): void {
    if (submitting.value || currentStep.value === 'finalResult' || navigateToHomePage.value) {
        return;
    }

    if (tabName === 'basicSetting') {
        currentStep.value = 'basicSetting';
    } else if (tabName === 'presetCategories') {
        const problemMessage = inputEmptyProblemMessage.value || inputInvalidProblemMessage.value;

        if (problemMessage) {
            snackbar.value?.showMessage(problemMessage);
            return;
        }

        currentStep.value = 'presetCategories';
    }
}

function switchToPreviousTab(): void {
    switchToTab('basicSetting');
}

function switchToNextTab(): void {
    switchToTab('presetCategories');
}

function submit(): void {
    const problemMessage = inputEmptyProblemMessage.value || inputInvalidProblemMessage.value;

    if (problemMessage) {
        snackbar.value?.showMessage(problemMessage);
        return;
    }

    navigateToHomePage.value = false;
    submitting.value = true;

    let presetCategories: LocalizedPresetCategory[] = [];

    if (usePresetCategories.value) {
        presetCategories = categorizedArrayToPlainArray(allPresetCategories.value);
    }

    rootStore.register({
        user: user.value,
        presetCategories: presetCategories
    }).then(response => {
        if (!isUserLogined()) {
            submitting.value = false;

            if (usePresetCategories.value && !response.presetCategoriesSaved) {
                finalResultMessage.value = tt('You have been successfully registered, but there was an failure when adding preset categories. You can re-add preset categories in settings page anytime.');
                currentStep.value = 'finalResult';
            } else if (response.needVerifyEmail) {
                finalResultMessage.value = tt('You have been successfully registered. An account activation link has been sent to your email address, please activate your account first.');
                currentStep.value = 'finalResult';
            } else {
                snackbar.value?.showMessage('You have been successfully registered');
                navigateToHomePage.value = true;
            }

            return;
        }

        doAfterSignupSuccess(response);
        submitting.value = false;

        if (usePresetCategories.value && !response.presetCategoriesSaved) {
            snackbar.value?.showMessage('You have been successfully registered, but there was an failure when adding preset categories. You can re-add preset categories in settings page anytime.');
        } else {
            snackbar.value?.showMessage('You have been successfully registered');
            router.replace('/');
        }

        navigateToHomePage.value = true;
    }).catch(error => {
        submitting.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function navigateToLogin(): void {
    router.push('/');
}

function onSnackbarShowStateChanged(newValue: boolean): void {
    if (!newValue && navigateToHomePage.value) {
        router.replace('/');
    }
}
</script>

<style lang="scss" scoped>
/* =========================================================
   SPILL CHILL LUXURY SIGNUP THEME (Aligned with landing.html & LoginPage.vue)
   Emerald (#10B981) • Forest (#064E3B) • Glassmorphism
   ========================================================= */

@keyframes pulse {
    0%, 100% { opacity: 1; transform: scale(1); }
    50% { opacity: 0.5; transform: scale(0.85); }
}

@keyframes scaleIn {
    from { opacity: 0; transform: scale(0.92); }
    to { opacity: 1; transform: scale(1); }
}

.animate-scale-in {
    animation: scaleIn 0.5s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.clean-signup-page {
    min-height: 100vh;
    width: 100%;
    background:
        radial-gradient(at 15% 15%, rgba(16, 185, 129, 0.12) 0px, transparent 50%),
        radial-gradient(at 85% 20%, rgba(6, 78, 59, 0.12) 0px, transparent 55%),
        radial-gradient(at 50% 80%, rgba(52, 211, 153, 0.08) 0px, transparent 50%),
        #F8FAFC !important;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 48px 16px;
    box-sizing: border-box;
    position: relative;
    overflow-y: auto;
}

:deep(.v-theme--dark) .clean-signup-page,
.v-theme--dark .clean-signup-page {
    background:
        radial-gradient(at 15% 15%, rgba(16, 185, 129, 0.15) 0px, transparent 50%),
        radial-gradient(at 85% 20%, rgba(6, 78, 59, 0.28) 0px, transparent 55%),
        radial-gradient(at 50% 80%, rgba(2, 44, 34, 0.35) 0px, transparent 50%),
        #0B0F19 !important;
}

.signup-container {
    width: 100%;
    max-width: 620px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
}

/* Brand Header */
.signup-brand {
    display: flex;
    justify-content: center;
    margin-bottom: 24px;

    .brand-link {
        display: inline-flex;
        align-items: center;
        gap: 12px;
        text-decoration: none;
    }

    .brand-icon-box {
        width: 44px;
        height: 44px;
        border-radius: 14px;
        background: linear-gradient(135deg, #10B981 0%, #059669 100%);
        display: flex;
        align-items: center;
        justify-content: center;
        box-shadow: 0 4px 14px rgba(16, 185, 129, 0.22);
        transition: transform 0.2s ease;

        &:hover {
            transform: scale(1.05);
        }
    }

    .signup-brand-title {
        font-family: 'Outfit', sans-serif;
        font-size: 24px;
        font-weight: 800;
        letter-spacing: -0.02em;
        color: #0f172a !important;
        margin: 0;
    }

    .brand-pulse-dot {
        width: 8px;
        height: 8px;
        border-radius: 50%;
        background: #10B981;
        display: inline-block;
        animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
    }

    .brand-tagline {
        font-family: 'Plus Jakarta Sans', sans-serif;
        font-size: 10.5px;
        font-weight: 700;
        text-transform: uppercase;
        letter-spacing: 0.08em;
        color: #059669;
    }
}

:deep(.v-theme--dark) .signup-brand .signup-brand-title {
    color: #ffffff !important;
}

/* Badge Pill */
.signup-badge-pill {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    padding: 5px 14px;
    border-radius: 9999px;
    background: rgba(16, 185, 129, 0.08);
    border: 1px solid rgba(16, 185, 129, 0.25);
    color: #047857;
    font-size: 11.5px;
    font-weight: 700;
    letter-spacing: 0.02em;

    .badge-dot {
        width: 7px;
        height: 7px;
        border-radius: 50%;
        background: #10B981;
    }
}

:deep(.v-theme--dark) .signup-badge-pill {
    background: rgba(16, 185, 129, 0.15);
    border-color: rgba(16, 185, 129, 0.4);
    color: #34d399;
}

/* Centered Luxury Glass Card - Custom Size */
.signup-card {
    width: 100%;
    max-width: 620px;
    background: rgba(255, 255, 255, 0.94) !important;
    -webkit-backdrop-filter: blur(20px);
    backdrop-filter: blur(20px);
    border: 1px solid rgba(226, 232, 240, 0.9) !important;
    border-radius: 24px !important;
    padding: 36px 36px 28px !important;
    box-shadow: 0 20px 45px -12px rgba(6, 78, 59, 0.1), 0 0 1px 1px rgba(16, 185, 129, 0.12) !important;
}

:deep(.v-theme--dark) .signup-card,
.v-theme--dark .signup-card {
    background: rgba(17, 24, 39, 0.90) !important;
    border: 1px solid rgba(255, 255, 255, 0.08) !important;
    box-shadow: 0 20px 45px -12px rgba(0, 0, 0, 0.6) !important;
}

.signup-card-header {
    .signup-title {
        font-family: 'Outfit', sans-serif;
        font-size: 26px;
        font-weight: 800;
        letter-spacing: -0.02em;
        color: #0f172a;
        margin: 0 0 6px;
    }

    .signup-subtitle {
        font-family: 'Plus Jakarta Sans', 'Inter', sans-serif;
        font-size: 13.5px;
        color: #64748b;
        margin: 0;
        line-height: 1.5;
    }
}

:deep(.v-theme--dark) .signup-card-header .signup-title {
    color: #ffffff;
}

/* Stepper Bar Tweaks */
:deep(.slide-group-with-stepper) {
    .step-number {
        font-family: 'Outfit', sans-serif !important;
        font-weight: 800 !important;
        font-size: 18px !important;
        color: #0f172a;
    }

    .step-title {
        font-family: 'Plus Jakarta Sans', sans-serif !important;
        font-weight: 700 !important;
        font-size: 13px !important;
        color: #1e293b;
    }

    .step-subtitle {
        font-family: 'Plus Jakarta Sans', sans-serif !important;
        font-size: 11px !important;
        color: #64748b;
    }

    .slide-group-stepper-link-line {
        background-color: #10B981 !important;
    }

    .slide-group-stepper-indicator {
        border-color: #10B981 !important;
    }
}

:deep(.v-theme--dark) .slide-group-with-stepper {
    .step-number {
        color: #f8fafc !important;
    }

    .step-title {
        color: #f1f5f9 !important;
    }

    .step-subtitle {
        color: #94a3b8 !important;
    }
}

/* Modern Input Fields */
.modern-input {
    :deep(.v-field) {
        border-radius: 14px !important;
        background: #ffffff !important;
        border: 1px solid #e2e8f0 !important;
        transition: border-color 0.15s ease, box-shadow 0.15s ease;

        &:hover {
            border-color: #cbd5e1 !important;
        }

        &.v-field--focused {
            border-color: #10B981 !important;
            box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.18) !important;
        }
    }

    :deep(.v-label) {
        font-family: 'Plus Jakarta Sans', 'Inter', sans-serif;
        font-size: 13.5px;
        font-weight: 500;
        color: #64748b;
    }

    :deep(input) {
        font-family: 'Plus Jakarta Sans', 'Inter', sans-serif;
        font-size: 14px;
        font-weight: 500;
        color: #0f172a;
    }

    :deep(.v-field__prepend-inner) {
        color: #10B981 !important;
    }
}

:deep(.v-theme--dark) .modern-input {
    :deep(.v-field) {
        background: #1e293b !important;
        border-color: #334155 !important;
    }
    :deep(input) {
        color: #f8fafc;
    }
}

/* Modern Expansion Panels */
.modern-panels {
    :deep(.v-expansion-panel) {
        border-radius: 14px !important;
        border: 1px solid #e2e8f0 !important;
        margin-bottom: 8px !important;
        background: #ffffff !important;
        transition: all 0.2s ease;

        &::after {
            display: none !important;
        }

        &.v-expansion-panel--active {
            border-color: #10B981 !important;
            box-shadow: 0 4px 14px rgba(16, 185, 129, 0.1) !important;
        }
    }
}

:deep(.v-theme--dark) .modern-panels {
    :deep(.v-expansion-panel) {
        background: #1e293b !important;
        border-color: #334155 !important;

        &.v-expansion-panel--active {
            border-color: #10B981 !important;
        }
    }
}

.expand-panel-title-with-bg {
    font-family: 'Outfit', sans-serif;
    font-weight: 600;
}

.border-t {
    border-top: 1px solid #e2e8f0;
}

:deep(.v-theme--dark) .border-t {
    border-top-color: #334155;
}

/* Stylish Scroll */
.stylish-scroll {
    scrollbar-width: thin;
    scrollbar-color: #cbd5e1 transparent;

    &::-webkit-scrollbar {
        width: 6px;
    }
    &::-webkit-scrollbar-track {
        background: transparent;
    }
    &::-webkit-scrollbar-thumb {
        background: #cbd5e1;
        border-radius: 9999px;
    }
}

:deep(.v-theme--dark) .stylish-scroll {
    scrollbar-color: #334155 transparent;
    &::-webkit-scrollbar-thumb {
        background: #334155;
    }
}

/* Buttons */
.premium-btn {
    background: linear-gradient(135deg, #10B981 0%, #059669 100%) !important;
    color: #ffffff !important;
    border-radius: 14px !important;
    height: 48px !important;
    font-family: 'Outfit', sans-serif !important;
    font-size: 14.5px !important;
    font-weight: 700 !important;
    letter-spacing: 0.01em !important;
    text-transform: none !important;
    box-shadow: 0 10px 25px -5px rgba(16, 185, 129, 0.25) !important;
    transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1) !important;

    &:hover:not(:disabled) {
        transform: translateY(-2px);
        box-shadow: 0 12px 28px -4px rgba(16, 185, 129, 0.35) !important;
    }

    &:active:not(:disabled) {
        transform: translateY(0);
    }

    &:disabled {
        opacity: 0.55;
        box-shadow: none !important;
    }
}

.premium-btn-text {
    border-radius: 14px !important;
    height: 48px !important;
    font-family: 'Outfit', sans-serif !important;
    font-size: 14px !important;
    font-weight: 600 !important;
    letter-spacing: 0 !important;
    text-transform: none !important;
}

.text-emerald-link {
    color: #059669;
    transition: color 0.15s ease;

    &:hover {
        color: #047857;
        text-decoration: underline;
    }
}

/* Trust Badge & Back Link */
.signup-trust-badge {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 12px;
    font-weight: 600;
    color: #64748b;
    display: flex;
    align-items: center;
    justify-content: center;

    .text-emerald {
        color: #10B981 !important;
    }
}

.back-to-landing-link {
    display: inline-flex;
    align-items: center;
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 12.5px;
    font-weight: 600;
    color: #64748b;
    text-decoration: none;
    transition: color 0.15s ease;

    &:hover {
        color: #10B981;
    }
}

@media (max-width: 600px) {
    .signup-card {
        padding: 24px 18px 20px !important;
        border-radius: 20px !important;
    }
}
</style>