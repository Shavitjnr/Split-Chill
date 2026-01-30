<template>
    <div class="layout-wrapper auth-wrapper">
        
        <router-link to="/" class="auth-logo animate-fade-in" style="text-decoration: none;">
            <div class="d-flex align-center gap-x-3">
                <img alt="logo" class="login-page-logo animate-float" :src="APPLICATION_LOGO_PATH" />
                <h1 class="font-weight-bold text-gradient" style="font-size: 1.8rem;">{{ tt('global.app.title') }}</h1>
            </div>
        </router-link>

        <v-row no-gutters class="h-100">
            
            <v-col cols="12" md="4" class="auth-image-background d-none d-md-flex align-center justify-center position-relative">
                <div class="hero-backdrop"></div>
                
                <HelixHero 
                    title="Join the Future"
                    description="Create your account and start your journey towards financial freedom today."
                />
            </v-col>

            
            <v-col cols="12" md="8" class="d-flex flex-column justify-center align-center auth-form-container px-6">
                <div class="d-flex align-center justify-center w-100 animate-scale-in">
                    <v-card class="glass-card w-100 px-10 py-12" max-width="900">
                        <v-card-text class="text-start px-0 mb-8">
                            <h2 class="text-h3 font-weight-bold mb-2" style="font-family: 'Outfit', sans-serif;">{{ tt('Create Account') }}</h2>
                            <p class="text-body-1 text-medium-emphasis">{{ tt('Fill in your details to get started with Split Chill AI') }}</p>
                        </v-card-text>

                        <v-card-text class="pa-0">
                            <steps-bar :steps="allSteps" :current-step="currentStep" :min-width="0" @step:change="switchToTab" class="mb-10" />

                            <v-window class="disable-tab-transition overflow-visible" v-model="currentStep">
                                <v-form @submit.prevent>
                                    <v-window-item value="basicSetting">
                                        <v-row>
                                            <v-col cols="12" md="6">
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

                            <div class="d-flex justify-space-between align-center mt-12 pt-8 border-t">
                                <v-btn variant="text" color="medium-emphasis"
                                       :disabled="currentStep === 'basicSetting' || currentStep === 'finalResult' || submitting || navigateToHomePage"
                                       :prepend-icon="mdiArrowLeft"
                                       @click="switchToPreviousTab"
                                       class="premium-btn-text">
                                    {{ tt('Back') }}
                                </v-btn>

                                <div class="d-flex align-center gap-4">
                                    <template v-if="currentStep === 'basicSetting'">
                                        <span class="text-body-2 text-medium-emphasis me-4 d-none d-sm-inline">{{ tt('Already have an account?') }} <router-link to="/login" class="text-primary font-weight-bold text-decoration-none h-link ms-1">{{ tt('Log In') }}</router-link></span>
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
                        </v-card-text>
                    </v-card>
                </div>
            </v-col>
        </v-row>

        <snack-bar ref="snackbar" @update:show="onSnackbarShowStateChanged" />
    </div>
</template>

<script setup lang="ts">
import HelixHero from '@/components/desktop/HelixHeroSimple.vue';
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
.layout-wrapper {
    overflow: hidden;
}

.auth-logo {
    position: absolute;
    top: 2.5rem;
    left: 2.5rem;
    z-index: 20;
    
    .login-page-logo {
        width: 48px;
        height: 48px;
        filter: drop-shadow(0 4px 10px rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.2));
    }
}

.auth-image-background {
    background: radial-gradient(circle at center, rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.05) 0%, transparent 70%);
}

.hero-backdrop {
    position: absolute;
    inset: 0;
    background: url('https://grainy-gradients.vercel.app/noise.svg');
    opacity: 0.03;
    pointer-events: none;
}

.modern-input {
    :deep(.v-field) {
        border-radius: 12px !important;
        background-color: rgba(var(--v-theme-on-surface), 0.02) !important;
        transition: all 0.2s ease;

        &.v-field--focused {
            background-color: #fff !important;
        }
    }

    :deep(.v-label) {
        font-family: 'Inter', sans-serif;
        font-weight: 500;
        opacity: 0.8;
    }
}

.premium-btn {
    border-radius: 12px !important;
    background: linear-gradient(135deg, var(--ebk-primary) 0%, #a855f7 100%) !important;
    font-family: 'Outfit', sans-serif !important;
    font-weight: 700 !important;
    letter-spacing: 0.5px !important;
    transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;

    &:hover:not(:disabled) {
        transform: translateY(-2px);
        
    }

    &:active:not(:disabled) {
        transform: translateY(0);
    }
}

.premium-btn-text {
    font-family: 'Outfit', sans-serif !important;
    font-weight: 600 !important;
    text-transform: none !important;
    letter-spacing: 0 !important;
}

.h-link {
    transition: all 0.2s ease;
    border-bottom: 2px solid transparent;

    &:hover {
        color: var(--ebk-primary) !important;
        border-bottom-color: var(--ebk-primary);
    }
}

.border-t {
    border-top: 1px solid rgba(var(--v-border-color), var(--v-border-opacity));
}

.stylish-scroll {
    &::-webkit-scrollbar {
        width: 6px;
    }
    &::-webkit-scrollbar-track {
        background: transparent;
    }
    &::-webkit-scrollbar-thumb {
        background: rgba(var(--v-theme-on-surface), 0.1);
        border-radius: 10px;
    }
}

.modern-panels {
    :deep(.v-expansion-panel) {
        background: rgba(var(--v-theme-on-surface), 0.02) !important;
        border: 1px solid rgba(var(--v-border-color), 0.05) !important;
        border-radius: 12px !important;
        margin-bottom: 8px !important;

        &::before {
            display: none !important;
        }

        &.v-expansion-panel--active {
            background: #fff !important;
            border-color: var(--ebk-primary) !important;
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05) !important;
        }
    }
}

.expand-panel-title-with-bg {
    font-family: 'Outfit', sans-serif;
}

.v-theme--dark {
    .modern-input :deep(.v-field--focused) {
        background-color: #1e293b !important;
    }
    .modern-panels :deep(.v-expansion-panel.v-expansion-panel--active) {
        background: #1e293b !important;
    }
}

.animate-scale-in {
    animation: scaleIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

@keyframes scaleIn {
    from { opacity: 0; transform: scale(0.95); }
    to { opacity: 1; transform: scale(1); }
}
</style>
