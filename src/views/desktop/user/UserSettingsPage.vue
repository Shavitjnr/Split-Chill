<template>
    <div class="user-settings-page-wrapper">
        <div class="settings-page-header mb-6">
            <h1 class="settings-page-title">{{ tt('User Settings') }}</h1>
            <p class="settings-page-subtitle">{{ tt('Manage your profile, security credentials, two-factor authentication, and data preferences') }}</p>
        </div>

        <div class="settings-tabs-container mb-6">
            <v-tabs show-arrows v-model="activeTab" class="settings-nav-tabs">
                <v-tab value="basicSetting" @click="pushRouter('basicSetting')" class="settings-tab-item">
                    <v-icon size="19" start :icon="mdiAccountOutline"/>
                    {{ tt('Basic') }}
                </v-tab>
                <v-tab value="securitySetting" @click="pushRouter('securitySetting')" class="settings-tab-item">
                    <v-icon size="19" start :icon="mdiLockOpenOutline"/>
                    {{ tt('Security') }}
                </v-tab>
                <v-tab value="twoFactorSetting" @click="pushRouter('twoFactorSetting')" class="settings-tab-item">
                    <v-icon size="19" start :icon="mdiOnepassword"/>
                    {{ tt('Two-Factor Authentication') }}
                </v-tab>
                <v-tab value="dataManagementSetting" @click="pushRouter('dataManagementSetting')" class="settings-tab-item">
                    <v-icon size="19" start :icon="mdiDatabaseCogOutline"/>
                    {{ tt('Data Management') }}
                </v-tab>
            </v-tabs>
        </div>

        <v-window class="settings-window-container disable-tab-transition" v-model="activeTab">
            <v-window-item value="basicSetting">
                <user-basic-setting-tab/>
            </v-window-item>

            <v-window-item value="securitySetting">
                <user-security-setting-tab/>
            </v-window-item>

            <v-window-item value="twoFactorSetting">
                <user-two-factor-auth-setting-tab ref="twoFactorSettingTab"/>
            </v-window-item>

            <v-window-item value="dataManagementSetting">
                <user-data-management-setting-tab/>
            </v-window-item>
        </v-window>
    </div>
</template>

<script setup lang="ts">
import UserBasicSettingTab from './settings/tabs/UserBasicSettingTab.vue';
import UserSecuritySettingTab from './settings/tabs/UserSecuritySettingTab.vue';
import UserTwoFactorAuthSettingTab from './settings/tabs/UserTwoFactorAuthSettingTab.vue';
import UserDataManagementSettingTab from './settings/tabs/UserDataManagementSettingTab.vue';

import { ref, useTemplateRef, watch } from 'vue';
import { useRouter, onBeforeRouteUpdate } from 'vue-router';

import { useI18n } from '@/locales/helpers.ts';

import {
    mdiAccountOutline,
    mdiLockOpenOutline,
    mdiOnepassword,
    mdiDatabaseCogOutline
} from '@mdi/js';

type TwoFactorSettingTabType = InstanceType<typeof UserTwoFactorAuthSettingTab>;

const props = defineProps<{
    initTab?: string;
}>();

const router = useRouter();

const { tt } = useI18n();

const ALL_TABS: string[] = [
    'basicSetting',
    'securitySetting',
    'twoFactorSetting',
    'dataManagementSetting'
];

const twoFactorSettingTab = useTemplateRef<TwoFactorSettingTabType>('twoFactorSettingTab');

const activeTab = ref<string>((() => {
    let queryActiveTab = props.initTab || 'basicSetting';

    if (ALL_TABS.indexOf(queryActiveTab) < 0) {
        queryActiveTab = 'basicSetting';
    }

    return queryActiveTab;
})());

const pushRouter = (tab: string) => {
    router.push(`/user/settings?tab=${tab}`);
};

onBeforeRouteUpdate((to) => {
    if (to.query && to.query['tab'] && ALL_TABS.indexOf(to.query['tab'] as string) >= 0) {
        activeTab.value = to.query['tab'] as string;
    } else {
        activeTab.value = 'basicSetting';
    }
});

watch(activeTab, (newValue, oldValue) => {
    if (oldValue === 'twoFactorSetting' && newValue !== 'twoFactorSetting') {
        twoFactorSettingTab.value?.reset();
    }
});
</script>

<style lang="scss" scoped>
/* =========================================================
   USER SETTINGS PAGE
   Modern Pill Tabs • Clean Layout
   ========================================================= */

.user-settings-page-wrapper {
    max-width: 1200px;
    margin: 0 auto;

    /* Shared card elevation & modern look for all tabs */
    :deep(.v-card) {
        border-radius: 20px !important;
        background: rgba(var(--v-theme-surface), 0.85) !important;
        backdrop-filter: blur(16px);
        border: 1px solid rgba(var(--v-border-color), 0.12) !important;
        box-shadow: 0 10px 30px -5px rgba(15, 23, 42, 0.05), 0 4px 12px -2px rgba(15, 23, 42, 0.02) !important;
    }

    :deep(.v-card-title),
    :deep(.v-card-item__title) {
        font-family: 'Outfit', sans-serif !important;
        font-weight: 700 !important;
        letter-spacing: -0.01em;
    }

    :deep(.v-field) {
        border-radius: 12px !important;
        background: rgba(var(--v-theme-on-surface), 0.02) !important;
        border: 1px solid rgba(var(--v-border-color), 0.12) !important;
        transition: all 0.2s ease;

        &:hover {
            border-color: rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.3) !important;
        }

        &.v-field--focused {
            border-color: rgb(var(--v-theme-primary)) !important;
            box-shadow: 0 0 0 3px rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.12) !important;
        }
    }

    :deep(.v-btn.bg-primary),
    :deep(.v-btn[type="submit"]),
    :deep(.settings-save-btn) {
        background: linear-gradient(135deg, #10B981 0%, #059669 100%) !important;
        color: #ffffff !important;
        border-radius: 12px !important;
        font-weight: 600 !important;
        box-shadow: 0 10px 25px -5px rgba(16, 185, 129, 0.25) !important;
    }

    :deep(.v-btn--variant-tonal:not([color="error"])) {
        border-radius: 10px !important;
    }
}

.settings-page-header {
    .settings-page-title {
        font-family: 'Outfit', sans-serif;
        font-size: 28px;
        font-weight: 700;
        letter-spacing: -0.02em;
        color: rgb(var(--v-theme-on-surface));
        margin: 0;
    }

    .settings-page-subtitle {
        font-family: 'Inter', sans-serif;
        font-size: 14px;
        color: rgba(var(--v-theme-on-surface), 0.65);
        margin: 4px 0 0;
    }
}

.settings-tabs-container {
    background: rgba(var(--v-theme-surface), 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(var(--v-border-color), 0.12);
    border-radius: 16px;
    padding: 6px;
    box-shadow: 0 4px 16px rgba(15, 23, 42, 0.03);

    .settings-nav-tabs {
        border-radius: 12px;

        :deep(.v-tab) {
            border-radius: 10px !important;
            font-family: 'Inter', sans-serif;
            font-size: 13.5px;
            font-weight: 500;
            text-transform: none;
            letter-spacing: 0;
            min-height: 42px;
            padding: 0 16px;
            transition: all 0.2s ease;
            color: rgba(var(--v-theme-on-surface), 0.75);

            &:hover {
                color: rgb(var(--v-theme-primary));
                background: rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.05);
            }
        }

        :deep(.v-tab--selected) {
            background: linear-gradient(135deg, rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.14) 0%, rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.08) 100%) !important;
            color: rgb(var(--v-theme-primary)) !important;
            font-weight: 600;
            box-shadow: 0 2px 8px rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.1);
        }
    }
}

.v-theme--dark {
    .user-settings-page-wrapper {
        :deep(.v-card) {
            background: rgba(30, 41, 59, 0.7) !important;
            border-color: rgba(255, 255, 255, 0.08) !important;
        }
    }

    .settings-tabs-container {
        background: rgba(30, 41, 59, 0.6);
        border-color: rgba(255, 255, 255, 0.08);

        .settings-nav-tabs {
            :deep(.v-tab--selected) {
                background: rgba(168, 85, 247, 0.18) !important;
                color: #c084fc !important;
            }
        }
    }
}
</style>
