<template>
    <div class="layout-wrapper layout-nav-type-vertical layout-navbar-static layout-footer-static layout-content-width-fluid"
         :class="{ 'layout-overlay-nav': mdAndDown }">
        <div class="layout-vertical-nav" :class="{'visible': showVerticalOverlayMenu, 'scrolled': isVerticalNavScrolled, 'overlay-nav': mdAndDown}">
            <div class="nav-header">
                <router-link to="/" class="brand-logo-link d-flex align-center">
                    <div class="brand-icon-box">
                        <v-icon :icon="mdiCallSplit" size="20" color="white" class="brand-icon-split" />
                    </div>
                    <div class="d-flex flex-column ms-3">
                        <span class="brand-title-text d-flex align-center">
                            Spill Chill
                            <span class="brand-dot-pulse"></span>
                        </span>
                        <span class="brand-subtitle-text">Expense Harmony</span>
                    </div>
                </router-link>
            </div>
            <perfect-scrollbar
                tag="ul" class="nav-items"
                :options="{ wheelPropagation: false }"
                @ps-scroll-y="handleNavScroll"
            >
                <li class="nav-link home-link">
                    <router-link to="/">
                        <v-icon class="nav-item-icon" :icon="mdiHomeOutline"/>
                        <span class="nav-item-title">{{ tt('Overview') }}</span>
                    </router-link>
                </li>
                <li class="nav-section-title">
                    <div class="title-wrapper">
                        <span class="title-text">{{ tt('Transaction Data') }}</span>
                    </div>
                </li>
                <li class="nav-link">
                    <router-link to="/transaction/list?pageType=0&dateType=7">
                        <v-icon class="nav-item-icon" :icon="mdiListBoxOutline"/>
                        <span class="nav-item-title d-inline-block">{{ tt('Transaction Details') }}</span>
                        <v-btn density="compact" color="secondary" variant="text" size="22"
                               class="ms-1" :icon="true" v-if="showAddTransactionButtonInDesktopNavbar"
                               @click="showAddDialogInTransactionListPage">
                            <v-icon :icon="mdiPlusCircle" size="22" />
                            <v-tooltip activator="parent">{{ tt('Add Transaction') }}</v-tooltip>
                        </v-btn>
                    </router-link>
                </li>
                <li class="nav-link">
                    <router-link to="/statistics/transaction">
                        <v-icon class="nav-item-icon" :icon="mdiChartPieOutline"/>
                        <span class="nav-item-title">{{ tt('Statistics & Analysis') }}</span>
                    </router-link>
                </li>
                <li class="nav-link">
                    <router-link to="/insights/explorer">
                        <v-icon class="nav-item-icon" :icon="mdiCompassOutline"/>
                        <span class="nav-item-title">{{ tt('Insights Explorer') }}</span>
                    </router-link>
                </li>
                <li class="nav-section-title">
                    <div class="title-wrapper">
                        <span class="title-text">{{ tt('Basis Data') }}</span>
                    </div>
                </li>
                <li class="nav-link">
                    <router-link to="/account/list">
                        <v-icon class="nav-item-icon" :icon="mdiCreditCardOutline"/>
                        <span class="nav-item-title">{{ tt('Accounts') }}</span>
                    </router-link>
                </li>
                <li class="nav-link">
                    <router-link to="/category/list">
                        <v-icon class="nav-item-icon" :icon="mdiViewDashboardOutline"/>
                        <span class="nav-item-title">{{ tt('Transaction Categories') }}</span>
                    </router-link>
                </li>
                <li class="nav-link">
                    <router-link to="/tag/list">
                        <v-icon class="nav-item-icon" :icon="mdiTagOutline"/>
                        <span class="nav-item-title">{{ tt('Transaction Tags') }}</span>
                    </router-link>
                </li>
                <li class="nav-link">
                    <router-link to="/template/list">
                        <v-icon class="nav-item-icon" :icon="mdiClipboardTextOutline"/>
                        <span class="nav-item-title">{{ tt('Transaction Templates') }}</span>
                    </router-link>
                </li>
                <li class="nav-link" v-if="isUserScheduledTransactionEnabled()">
                    <router-link to="/schedule/list">
                        <v-icon class="nav-item-icon" :icon="mdiClipboardTextClockOutline"/>
                        <span class="nav-item-title">{{ tt('Scheduled Transactions') }}</span>
                    </router-link>
                </li>
                <li class="nav-section-title">
                    <div class="title-wrapper">
                        <span class="title-text">{{ tt('Miscellaneous') }}</span>
                    </div>
                </li>
                <li class="nav-link">
                    <router-link to="/exchange_rates">
                        <v-icon class="nav-item-icon" :icon="mdiSwapHorizontal"/>
                        <span class="nav-item-title">{{ tt('Exchange Rates Data') }}</span>
                    </router-link>
                </li>
                <li class="nav-link">
                    <a href="javascript:void(0);" @click="showMobileQrCode = true">
                        <v-icon class="nav-item-icon" :icon="mdiCellphone"/>
                        <span class="nav-item-title">{{ tt('Use on Mobile Device') }}</span>
                    </a>
                </li>
                <li class="nav-link">
                    <router-link to="/about">
                        <v-icon class="nav-item-icon" :icon="mdiInformationOutline"/>
                        <span class="nav-item-title">{{ tt('About') }}</span>
                    </router-link>
                </li>
            </perfect-scrollbar>
        </div>

        <div class="layout-content-wrapper">
            <div class="layout-navbar navbar-blur">
                <div class="navbar-content-container">
                    <div class="d-flex h-100 align-center">
                        <v-btn class="ms-n3 me-2 d-lg-none" color="default" variant="text"
                               :icon="true" @click="showVerticalOverlayMenu = true">
                            <v-icon :icon="mdiMenu" size="24" />
                        </v-btn>
                        <div class="brand-logo-link d-flex align-center ms-2" v-if="mdAndDown">
                            <div class="brand-icon-box brand-icon-box-sm">
                                <v-icon :icon="mdiCallSplit" size="16" color="white" class="brand-icon-split" />
                            </div>
                            <div class="d-flex flex-column ms-2">
                                <span class="brand-title-text brand-title-sm d-flex align-center">
                                    Spill Chill
                                </span>
                            </div>
                        </div>
                        <v-spacer />
                        <v-btn color="primary" variant="text" class="me-2"
                               :icon="true" @click="(currentTheme === 'light' ? currentTheme = 'dark' : (currentTheme === 'dark' ? currentTheme = 'auto' : currentTheme = 'light'))">
                            <v-icon :icon="(currentTheme === 'light' ? mdiWeatherSunny : (currentTheme === 'dark' ? mdiWeatherNight : mdiThemeLightDark))" size="24" />
                        </v-btn>
                        <v-avatar class="cursor-pointer" variant="tonal"
                                  :color="currentUserAvatar ? 'rgba(0,0,0,0)' : 'primary'">
                            <v-img :src="currentUserAvatar" v-if="currentUserAvatar">
                                <template #placeholder>
                                    <div class="d-flex align-center justify-center fill-height bg-light-primary">
                                        <v-icon color="primary" :icon="mdiAccount"/>
                                    </div>
                                </template>
                            </v-img>
                            <v-icon :icon="mdiAccount" v-else-if="!currentUserAvatar"/>
                            <v-menu activator="parent" width="230" location="bottom end" offset="14px">
                                <v-list>
                                    <v-list-item>
                                        <template #prepend>
                                            <v-list-item-action>
                                                <v-avatar variant="tonal"
                                                          :color="currentUserAvatar ? 'rgba(0,0,0,0)' : 'primary'">
                                                    <v-img :src="currentUserAvatar" v-if="currentUserAvatar">
                                                        <template #placeholder>
                                                            <div class="d-flex align-center justify-center fill-height bg-light-primary">
                                                                <v-icon color="primary" :icon="mdiAccount"/>
                                                            </div>
                                                        </template>
                                                    </v-img>
                                                    <v-icon :icon="mdiAccount" v-else-if="!currentUserAvatar"/>
                                                </v-avatar>
                                            </v-list-item-action>
                                        </template>
                                        <v-list-item-title class="ms-2">
                                            {{ currentNickName }}
                                        </v-list-item-title>
                                    </v-list-item>
                                    <v-divider class="my-2"/>
                                    <v-list-item :prepend-icon="mdiAccountCogOutline"
                                                 :title="tt('User Settings')"
                                                 to="/user/settings"></v-list-item>
                                    <v-list-item :prepend-icon="mdiCogOutline"
                                                 :title="tt('Application Settings')"
                                                 to="/app/settings"></v-list-item>
                                    <v-divider class="my-2"/>
                                    <v-list-item :prepend-icon="mdiLockOutline"
                                                 :title="tt('Lock Application')"
                                                 v-if="isEnableApplicationLock"
                                                 @click="lock"></v-list-item>
                                    <v-list-item :disabled="logouting"
                                                 :prepend-icon="mdiLogout"
                                                 :title="tt('Log Out')"
                                                 @click="logout"></v-list-item>
                                </v-list>
                            </v-menu>
                        </v-avatar>
                    </div>
                </div>
            </div>
            <div class="layout-page-content">
                <div class="page-content-container">
                    <router-view :key="currentRoutePath" />
                </div>
            </div>
        </div>

        <switch-to-mobile-dialog v-model:show="showMobileQrCode" />

        <div class="layout-overlay" :class="{ 'visible': showVerticalOverlayMenu }" @click="showVerticalOverlayMenu = false"></div>

        <v-overlay class="justify-center align-center" :persistent="true" v-model="showLoading">
            <v-progress-circular indeterminate></v-progress-circular>
        </v-overlay>

        <snack-bar ref="snackbar" />
    </div>
</template>

<script setup lang="ts">
import SnackBar from '@/components/desktop/SnackBar.vue';

import { ref, computed, useTemplateRef } from 'vue';

import { useDisplay, useTheme } from 'vuetify';
import { useRoute, useRouter } from 'vue-router';

import { useI18n } from '@/locales/helpers.ts';

import { useRootStore } from '@/stores/index.ts';
import { useSettingsStore } from '@/stores/setting.ts';
import { useUserStore } from '@/stores/user.ts';
import { useDesktopPageStore } from '@/stores/desktopPage.ts';

import { ThemeType } from '@/core/theme.ts';
import { isUserScheduledTransactionEnabled } from '@/lib/server_settings.ts';
import { getSystemTheme, setExpenseAndIncomeAmountColor } from '@/lib/ui/common.ts';

import {
    mdiMenu,
    mdiHomeOutline,
    mdiListBoxOutline,
    mdiPlusCircle,
    mdiCreditCardOutline,
    mdiViewDashboardOutline,
    mdiTagOutline,
    mdiClipboardTextOutline,
    mdiClipboardTextClockOutline,
    mdiChartPieOutline,
    mdiCompassOutline,
    mdiSwapHorizontal,
    mdiCogOutline,
    mdiCellphone,
    mdiInformationOutline,
    mdiThemeLightDark,
    mdiWeatherSunny,
    mdiWeatherNight,
    mdiAccount,
    mdiAccountCogOutline,
    mdiLockOutline,
    mdiLogout,
    mdiCallSplit
} from '@mdi/js';

type SnackBarType = InstanceType<typeof SnackBar>;

const display = useDisplay();
const theme = useTheme();
const route = useRoute();
const router = useRouter();

const { tt, initLocale } = useI18n();

const rootStore = useRootStore();
const settingsStore = useSettingsStore();
const userStore = useUserStore();
const desktopPageStore = useDesktopPageStore();

const snackbar = useTemplateRef<SnackBarType>('snackbar');

const logouting = ref<boolean>(false);
const isVerticalNavScrolled = ref<boolean>(false);
const showVerticalOverlayMenu = ref<boolean>(false);
const showLoading = ref<boolean>(false);
const showMobileQrCode = ref<boolean>(false);

const mdAndDown = computed<boolean>(() => display.mdAndDown.value);
const currentRoutePath = computed<string>(() => route.path);

const currentNickName = computed<string>(() => userStore.currentUserNickname || tt('User'));
const currentUserAvatar = computed<string | null>(() => userStore.getUserAvatarUrl(userStore.currentUserBasicInfo, true));

const currentTheme = computed<string>({
    get: () => {
        return settingsStore.appSettings.theme;
    },
    set: (value: string) => {
        if (value !== settingsStore.appSettings.theme) {
            settingsStore.setTheme(value);

            if (value === ThemeType.Light || value === ThemeType.Dark) {
                theme.change(value);
            } else {
                theme.change(getSystemTheme());
            }
        }
    }
});

const showAddTransactionButtonInDesktopNavbar = computed<boolean>(() => settingsStore.appSettings.showAddTransactionButtonInDesktopNavbar);
const isEnableApplicationLock = computed<boolean>(() => settingsStore.appSettings.applicationLock);

function handleNavScroll(e: Event): void {
    isVerticalNavScrolled.value = (e.target as HTMLElement).scrollTop > 0;
}

function lock(): void {
    rootStore.lock();
    router.replace('/unlock');
}

function logout(): void {
    logouting.value = true;
    showLoading.value = true;

    rootStore.logout().then(() => {
        logouting.value = false;
        showLoading.value = false;

        settingsStore.clearAppSettings();

        const localeDefaultSettings = initLocale(userStore.currentUserLanguage, settingsStore.appSettings.timeZone);
        settingsStore.updateLocalizedDefaultSettings(localeDefaultSettings);

        setExpenseAndIncomeAmountColor(userStore.currentUserExpenseAmountColor, userStore.currentUserIncomeAmountColor);

        router.replace('/login');
    }).catch(error => {
        logouting.value = false;
        showLoading.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function showAddDialogInTransactionListPage(): void {
    desktopPageStore.setShowAddTransactionDialogInTransactionList();
}
</script>

<!-- <style>
.main-logo {
    width: 1.75rem;
    height: 1.75rem;
}

.nav-link.home-link > a:not(.router-link-exact-active):hover::before {
    opacity: calc(var(--v-hover-opacity)* var(--v-theme-overlay-multiplier));
}
</style> -->

<style>
/* =========================================================
   SPILL CHILL EXECUTIVE NAVIGATION & BRANDING
   Clean • Minimal • Pure Emerald Fintech SaaS
   ========================================================= */

.brand-logo-link {
    text-decoration: none;
    padding: 2px 0;
}

.brand-icon-box {
    width: 40px;
    height: 40px;
    border-radius: 11px;
    background: linear-gradient(135deg, #10B981 0%, #059669 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.22);
    flex-shrink: 0;
}

.brand-icon-split {
    transform: rotate(-45deg);
}

.brand-title-text {
    font-family: 'Outfit', sans-serif !important;
    font-weight: 800;
    font-size: 1.25rem;
    letter-spacing: -0.02em;
    color: #0F172A;
    line-height: 1.1;
}

.brand-dot-pulse {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: #10B981;
    display: inline-block;
    margin-left: 6px;
}

.brand-subtitle-text {
    font-family: 'Outfit', sans-serif !important;
    font-size: 9px;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    color: #10B981;
    line-height: 1;
    margin-top: 3px;
}

.brand-icon-box-sm {
    width: 32px;
    height: 32px;
    border-radius: 8px;
}

.brand-title-sm {
    font-size: 1.05rem;
}

/* ---------------------------------------------------------
   Sidebar & Nav Links
   --------------------------------------------------------- */

.layout-vertical-nav {
    border-right: 1px solid #E2E8F0;
    background-color: #FFFFFF !important;
}

.nav-link > a {
    position: relative;
    border-radius: 10px;
    font-family: 'Inter', sans-serif;
    font-weight: 500;
    font-size: 0.875rem;
    color: #475569 !important;
    transition: color 0.15s ease, background-color 0.15s ease;
}

.nav-link > a:hover {
    color: #0F172A !important;
    background: rgba(16, 185, 129, 0.06) !important;
}

.nav-link > a.router-link-exact-active,
.nav-link > a.router-link-active:not([href="/"]) {
    color: #065F46 !important;
    background: rgba(16, 185, 129, 0.10) !important;
    font-weight: 600;
    border-left: 3px solid #10B981 !important;

    .nav-item-icon {
        color: #10B981 !important;
    }
}

.layout-navbar {
    border-bottom: 1px solid #E2E8F0;
    background: rgba(255, 255, 255, 0.85);
    backdrop-filter: blur(12px);
}

/* ---------------------------------------------------------
   Dark Mode
   --------------------------------------------------------- */

.v-theme--dark {
    .layout-vertical-nav {
        background-color: #111827 !important;
        border-right: 1px solid rgba(255, 255, 255, 0.08);
    }

    .brand-title-text {
        color: #F8FAFC;
    }

    .nav-link > a {
        color: #94A3B8 !important;
    }

    .nav-link > a:hover {
        color: #F8FAFC !important;
        background: rgba(16, 185, 129, 0.10) !important;
    }

    .nav-link > a.router-link-exact-active,
    .nav-link > a.router-link-active:not([href="/"]) {
        color: #34D399 !important;
        background: rgba(16, 185, 129, 0.16) !important;
        border-left: 3px solid #34D399 !important;

        .nav-item-icon {
            color: #34D399 !important;
        }
    }

    .layout-navbar {
        border-bottom: 1px solid rgba(255, 255, 255, 0.08);
        background: rgba(17, 24, 39, 0.85);
    }
}
</style>