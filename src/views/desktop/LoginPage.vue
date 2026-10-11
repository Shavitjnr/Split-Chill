<template>
    <div class="clean-login-page">
        <div class="login-container">
            <!-- Brand Logo & Header matching landing.html -->
            <div class="login-brand mb-6 text-center">
                <router-link to="/" class="brand-link">
                    <div class="brand-icon-box">
                        <v-icon icon="mdi-call-split" class="brand-icon" size="24" color="white" />
                    </div>
                    <div class="d-flex flex-column text-start">
                        <div class="d-flex align-center gap-2">
                            <h1 class="login-brand-title">{{ tt('global.app.title') }}</h1>
                            <span class="brand-pulse-dot"></span>
                        </div>
                        <span class="brand-tagline">Expense Harmony</span>
                    </div>
                </router-link>
            </div>

            <!-- Centered Informational Box / Card -->
            <v-card class="login-card" elevation="0">
                <div class="d-flex justify-center mb-4">
                    <div class="login-badge-pill">
                        <span class="badge-dot"></span>
                        <span>Next-Gen Expense Intelligence</span>
                    </div>
                </div>

                <div class="login-card-header text-center mb-6">
                    <h2 class="login-title">{{ tt('Login') }}</h2>
                    <p class="login-subtitle">{{ tt('Please enter your credentials to access your account') }}</p>
                    <p class="mt-2 text-error font-weight-medium" v-if="tips">{{ tips }}</p>
                </div>

                <v-form @submit.prevent="login">
                    <div v-if="isInternalAuthEnabled()">
                        <div class="form-field-group mb-4">
                            <v-text-field
                                variant="outlined"
                                density="comfortable"
                                type="text"
                                autocomplete="username"
                                :autofocus="true"
                                :disabled="show2faInput || loggingInByPassword || loggingInByOAuth2 || verifying"
                                :label="tt('Username or Email')"
                                :placeholder="tt('Username or Email')"
                                prepend-inner-icon="mdi-account-outline"
                                v-model.trim="username"
                                @input="tempToken = ''"
                                @keyup.enter="passwordInput?.focus()"
                                class="clean-input"
                            />
                        </div>

                        <div class="form-field-group mb-2">
                            <v-text-field
                                variant="outlined"
                                density="comfortable"
                                autocomplete="current-password"
                                ref="passwordInput"
                                type="password"
                                :disabled="show2faInput || loggingInByPassword || loggingInByOAuth2 || verifying"
                                :label="tt('Password')"
                                :placeholder="tt('Password')"
                                prepend-inner-icon="mdi-lock-outline"
                                v-model="password"
                                @input="tempToken = ''"
                                @keyup.enter="login"
                                class="clean-input"
                            />
                        </div>
                    </div>

                    <!-- 2FA Fields -->
                    <div v-show="show2faInput" class="mb-4">
                        <v-text-field
                            variant="outlined"
                            density="comfortable"
                            type="number"
                            autocomplete="one-time-code"
                            ref="passcodeInput"
                            :disabled="loggingInByPassword || loggingInByOAuth2 || verifying"
                            :label="tt('Passcode')"
                            prepend-inner-icon="mdi-shield-check-outline"
                            :append-inner-icon="mdiHelpCircleOutline"
                            v-model="passcode"
                            @click:append-inner="twoFAVerifyType = 'backupcode'"
                            @keyup.enter="verify"
                            v-if="twoFAVerifyType === 'passcode'"
                            class="clean-input mb-3"
                        />
                        <v-text-field
                            variant="outlined"
                            density="comfortable"
                            type="text"
                            ref="backupCodeInput"
                            :disabled="loggingInByPassword || loggingInByOAuth2 || verifying"
                            :label="tt('Backup Code')"
                            prepend-inner-icon="mdi-key-outline"
                            :append-inner-icon="mdiOnepassword"
                            v-model="backupCode"
                            @click:append-inner="twoFAVerifyType = 'passcode'"
                            @keyup.enter="verify"
                            v-if="twoFAVerifyType === 'backupcode'"
                            class="clean-input mb-3"
                        />
                    </div>

                    <div class="d-flex align-center justify-end mb-6" v-if="isInternalAuthEnabled()">
                        <router-link class="forgot-password-link" to="/forgetpassword"
                                     :class="{ 'disabled': !isUserForgetPasswordEnabled() || loggingInByPassword || loggingInByOAuth2 || verifying }">
                            {{ tt('Forgot password?') }}
                        </router-link>
                    </div>

                    <!-- Submit Buttons -->
                    <v-btn block size="large" elevation="0" class="login-submit-btn mb-4"
                           :disabled="inputIsEmpty || loggingInByPassword || loggingInByOAuth2 || verifying"
                           @click="login" v-if="isInternalAuthEnabled() && !show2faInput">
                        <span>{{ tt('Log In') }}</span>
                        <v-progress-circular indeterminate size="20" width="2" class="ms-3 text-white" v-if="loggingInByPassword"></v-progress-circular>
                    </v-btn>

                    <v-btn block size="large" elevation="0" class="login-submit-btn mb-4"
                           :disabled="twoFAInputIsEmpty || loggingInByPassword || loggingInByOAuth2 || verifying"
                           @click="verify" v-else-if="isInternalAuthEnabled() && show2faInput">
                        <span>{{ tt('Verify & Continue') }}</span>
                        <v-progress-circular indeterminate size="20" width="2" class="ms-3 text-white" v-if="verifying"></v-progress-circular>
                    </v-btn>

                    <!-- OAuth2 -->
                    <template v-if="isOAuth2Enabled()">
                        <div class="d-flex align-center my-5">
                            <v-divider />
                            <span class="px-3 text-caption text-medium-emphasis text-uppercase">{{ tt('or') }}</span>
                            <v-divider />
                        </div>

                        <v-btn block variant="outlined" size="large" class="oauth-btn mb-4"
                               :disabled="show2faInput || loggingInByPassword || loggingInByOAuth2 || verifying" :href="oauth2LoginUrl"
                               @click="loggingInByOAuth2 = true">
                            <v-icon start icon="mdi-account-circle-outline" class="me-2" />
                            <span>{{ oauth2LoginDisplayName }}</span>
                            <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="loggingInByOAuth2"></v-progress-circular>
                        </v-btn>
                    </template>

                    <!-- Register Link -->
                    <div class="text-center mt-6" v-if="isInternalAuthEnabled()">
                        <span class="text-caption text-medium-emphasis">{{ tt('New to Split Chill AI?') }}</span>
                        <router-link class="create-account-link ms-2" to="/signup"
                                     :class="{ 'disabled': !isUserRegistrationEnabled() || loggingInByPassword || loggingInByOAuth2 || verifying }">
                            {{ tt('Create an account') }}
                        </router-link>
                    </div>
                </v-form>

                <!-- Footer Utilities -->
                <div class="card-footer-bar mt-8 pt-4">
                    <div class="d-flex align-center justify-space-between flex-wrap gap-2">
                        <language-select-button :disabled="loggingInByPassword || loggingInByOAuth2 || verifying" variant="text" size="small" />
                        <a href="javascript:void(0);"
                           class="mobile-app-link"
                           :class="{ 'disabled': loggingInByPassword || loggingInByOAuth2 || verifying }"
                           @click="showMobileQrCode = true">
                            {{ tt('Mobile App') }}
                        </a>
                    </div>
                    <div class="text-center mt-3 text-caption text-medium-emphasis">
                        <span>{{ tt('global.app.title') }} v{{ version }} &copy; {{ new Date().getFullYear() }}</span>
                    </div>
                </div>
            </v-card>

            <!-- Trust Badge & Back to Landing link matching landing.html -->
            <div class="login-trust-badge text-center mt-6">
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

        <switch-to-mobile-dialog v-model:show="showMobileQrCode" />
        <snack-bar ref="snackbar" />
    </div>
</template>

<script setup lang="ts">
import { VTextField } from 'vuetify/components/VTextField';
import SnackBar from '@/components/desktop/SnackBar.vue';
import SwitchToMobileDialog from '@/components/desktop/SwitchToMobileDialog.vue';
import LanguageSelectButton from '@/components/desktop/LanguageSelectButton.vue';

import { ref, useTemplateRef, nextTick } from 'vue';
import { useRouter } from 'vue-router';

import { useI18n } from '@/locales/helpers.ts';
import { useLoginPageBase } from '@/views/base/LoginPageBase.ts';

import { useRootStore } from '@/stores/index.ts';

import { APPLICATION_LOGO_PATH } from '@/consts/asset.ts';
import { KnownErrorCode } from '@/consts/api.ts';

import { generateRandomUUID } from '@/lib/misc.ts';
import {
    isUserRegistrationEnabled,
    isUserForgetPasswordEnabled,
    isUserVerifyEmailEnabled,
    isInternalAuthEnabled,
    isOAuth2Enabled
} from '@/lib/server_settings.ts';

import {
    mdiOnepassword,
    mdiHelpCircleOutline
} from '@mdi/js';

type SnackBarType = InstanceType<typeof SnackBar>;

const router = useRouter();

const { tt } = useI18n();

const rootStore = useRootStore();

const {
    version,
    username,
    password,
    passcode,
    backupCode,
    tempToken,
    twoFAVerifyType,
    oauth2ClientSessionId,
    loggingInByPassword,
    loggingInByOAuth2,
    verifying,
    inputIsEmpty,
    twoFAInputIsEmpty,
    oauth2LoginUrl,
    oauth2LoginDisplayName,
    tips,
    doAfterLogin
} = useLoginPageBase('desktop');

const passwordInput = useTemplateRef<VTextField>('passwordInput');
const passcodeInput = useTemplateRef<VTextField>('passcodeInput');
const snackbar = useTemplateRef<SnackBarType>('snackbar');

const show2faInput = ref<boolean>(false);
const showMobileQrCode = ref<boolean>(false);

function login(): void {
    if (!username.value) {
        snackbar.value?.showMessage('Username cannot be blank');
        return;
    }

    if (!password.value) {
        snackbar.value?.showMessage('Password cannot be blank');
        return;
    }

    if (tempToken.value) {
        show2faInput.value = true;
        return;
    }

    if (loggingInByPassword.value) {
        return;
    }

    loggingInByPassword.value = true;

    rootStore.authorize({
        loginName: username.value,
        password: password.value
    }).then(authResponse => {
        loggingInByPassword.value = false;

        if (authResponse.need2FA) {
            tempToken.value = authResponse.token;
            show2faInput.value = true;

            nextTick(() => {
                if (passcodeInput.value) {
                    passcodeInput.value.focus();
                    passcodeInput.value.select();
                }
            });

            return;
        }

        doAfterLogin(authResponse);
        router.replace('/');
    }).catch(error => {
        loggingInByPassword.value = false;

        if (isUserVerifyEmailEnabled() && error.error && error.error.errorCode === KnownErrorCode.UserEmailNotVerified && error.error.context && error.error.context.email) {
            router.push(`/verify_email?email=${encodeURIComponent(error.error.context.email)}&emailSent=${error.error.context.hasValidEmailVerifyToken || false}`);
            return;
        }

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

function verify(): void {
    if (twoFAInputIsEmpty.value || verifying.value) {
        return;
    }

    if (twoFAVerifyType.value === 'passcode' && !passcode.value) {
        snackbar.value?.showMessage('Passcode cannot be blank');
        return;
    } else if (twoFAVerifyType.value === 'backupcode' && !backupCode.value) {
        snackbar.value?.showMessage('Backup code cannot be blank');
        return;
    }

    verifying.value = true;

    rootStore.authorize2FA({
        token: tempToken.value,
        passcode: twoFAVerifyType.value === 'passcode' ? passcode.value : null,
        recoveryCode: twoFAVerifyType.value === 'backupcode' ? backupCode.value : null
    }).then(authResponse => {
        verifying.value = false;

        doAfterLogin(authResponse);
        router.replace('/');
    }).catch(error => {
        verifying.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}

oauth2ClientSessionId.value = generateRandomUUID();
</script>

<style lang="scss" scoped>
/* =========================================================
   SPILL CHILL LUXURY THEME (Aligned with landing.html)
   Emerald (#10B981) • Forest (#064E3B) • Glassmorphism
   ========================================================= */

@keyframes pulse {
    0%, 100% { opacity: 1; transform: scale(1); }
    50% { opacity: 0.5; transform: scale(0.85); }
}

.clean-login-page {
    min-height: 100vh;
    width: 100%;
    background:
        radial-gradient(at 15% 15%, rgba(16, 185, 129, 0.12) 0px, transparent 50%),
        radial-gradient(at 85% 20%, rgba(6, 78, 59, 0.12) 0px, transparent 55%),
        radial-gradient(at 50% 80%, rgba(52, 211, 153, 0.08) 0px, transparent 50%),
        #F8FAFC !important;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 40px 16px;
    box-sizing: border-box;
    position: relative;
    overflow-y: auto;
}

:deep(.v-theme--dark) .clean-login-page,
.v-theme--dark .clean-login-page {
    background:
        radial-gradient(at 15% 15%, rgba(16, 185, 129, 0.15) 0px, transparent 50%),
        radial-gradient(at 85% 20%, rgba(6, 78, 59, 0.28) 0px, transparent 55%),
        radial-gradient(at 50% 80%, rgba(2, 44, 34, 0.35) 0px, transparent 50%),
        #0B0F19 !important;
}

.login-container {
    width: 100%;
    max-width: 440px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
}

/* Brand Header */
.login-brand {
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

    .login-brand-title {
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

:deep(.v-theme--dark) .login-brand .login-brand-title {
    color: #ffffff !important;
}

/* Badge Pill */
.login-badge-pill {
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

:deep(.v-theme--dark) .login-badge-pill {
    background: rgba(16, 185, 129, 0.15);
    border-color: rgba(16, 185, 129, 0.4);
    color: #34d399;
}

/* Centered Glass Card */
.login-card {
    width: 100%;
    background: rgba(255, 255, 255, 0.92) !important;
    -webkit-backdrop-filter: blur(20px);
    backdrop-filter: blur(20px);
    border: 1px solid rgba(226, 232, 240, 0.9) !important;
    border-radius: 24px !important;
    padding: 38px 34px 28px !important;
    box-shadow: 0 20px 45px -12px rgba(6, 78, 59, 0.1), 0 0 1px 1px rgba(16, 185, 129, 0.12) !important;
}

:deep(.v-theme--dark) .login-card,
.v-theme--dark .login-card {
    background: rgba(17, 24, 39, 0.88) !important;
    border: 1px solid rgba(255, 255, 255, 0.08) !important;
    box-shadow: 0 20px 45px -12px rgba(0, 0, 0, 0.6) !important;
}

.login-card-header {
    .login-title {
        font-family: 'Outfit', sans-serif;
        font-size: 26px;
        font-weight: 800;
        letter-spacing: -0.02em;
        color: #0f172a;
        margin: 0 0 6px;
    }

    .login-subtitle {
        font-family: 'Plus Jakarta Sans', 'Inter', sans-serif;
        font-size: 14px;
        color: #64748b;
        margin: 0;
        line-height: 1.5;
    }
}

:deep(.v-theme--dark) .login-card-header .login-title {
    color: #ffffff;
}

/* Input Fields */
.clean-input {
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
        font-size: 14px;
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

:deep(.v-theme--dark) .clean-input {
    :deep(.v-field) {
        background: #1e293b !important;
        border-color: #334155 !important;
    }
    :deep(input) {
        color: #f8fafc;
    }
}

/* Submit Buttons with Emerald & Forest Gradient */
.login-submit-btn {
    background: linear-gradient(135deg, #10B981 0%, #059669 100%) !important;
    color: #ffffff !important;
    border-radius: 14px !important;
    height: 50px !important;
    font-family: 'Outfit', sans-serif !important;
    font-size: 15px !important;
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

.oauth-btn {
    border-radius: 14px !important;
    height: 48px !important;
    border: 1px solid #e2e8f0 !important;
    color: #0f172a !important;
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 14px !important;
    font-weight: 600 !important;
    letter-spacing: 0 !important;
    text-transform: none !important;

    &:hover:not(:disabled) {
        background: #ecfdf5 !important;
        border-color: #10B981 !important;
        color: #047857 !important;
    }
}

:deep(.v-theme--dark) .oauth-btn {
    border-color: #334155 !important;
    color: #f8fafc !important;
    background: transparent !important;

    &:hover:not(:disabled) {
        background: #064E3B !important;
        border-color: #10B981 !important;
    }
}

/* Links */
.forgot-password-link {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 13px;
    font-weight: 700;
    color: #059669;
    text-decoration: none;
    transition: color 0.15s ease;

    &:hover {
        color: #047857;
        text-decoration: underline;
    }
}

.create-account-link {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 13px;
    font-weight: 700;
    color: #059669;
    text-decoration: none;
    transition: color 0.15s ease;

    &:hover {
        color: #047857;
        text-decoration: underline;
    }
}

.mobile-app-link {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 12px;
    font-weight: 600;
    color: #64748b;
    text-decoration: none;

    &:hover {
        color: #10B981;
    }
}

.card-footer-bar {
    border-top: 1px solid #f1f5f9;
}

:deep(.v-theme--dark) .card-footer-bar {
    border-top-color: #1e293b;
}

/* Trust Badge & Back Link */
.login-trust-badge {
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

@media (max-width: 480px) {
    .login-card {
        padding: 28px 20px 20px !important;
        border-radius: 20px !important;
    }
}
</style>
