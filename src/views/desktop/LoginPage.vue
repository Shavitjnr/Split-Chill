<template>
    <div class="layout-wrapper auth-wrapper">
        
        <router-link to="/" class="auth-logo animate-fade-in" style="text-decoration: none;">
            <div class="d-flex align-center gap-x-3">
                <img alt="logo" class="login-page-logo animate-float" :src="APPLICATION_LOGO_PATH" />
                <h1 class="font-weight-bold text-gradient" style="font-size: 1.8rem;">{{ tt('global.app.title') }}</h1>
            </div>
        </router-link>

        <v-row no-gutters class="h-100">
            
            <v-col cols="12" md="8" class="auth-image-background d-none d-md-flex align-center justify-center position-relative">
                <div class="hero-backdrop"></div>
                
                <HelixHero 
                    title="Split Chill AI"
                    description="The future of personal finance management. Track, analyze, and optimize your wealth with intelligent automation and stunning visualizations."
                />
            </v-col>

            
            <v-col cols="12" md="4" class="d-flex flex-column justify-center align-center px-6">
                <div class="d-flex align-center justify-center w-100 animate-scale-in">
                    <v-card class="glass-card w-100 px-8 pt-10 pb-6" max-width="460">
                        <v-card-text class="text-start px-0">
                            <h2 class="text-h3 font-weight-bold mb-1" style="font-family: 'Outfit', sans-serif;">{{ tt('Login') }}</h2>
                            <p class="text-body-1 text-medium-emphasis mb-8">{{ tt('Please enter your credentials to access your account') }}</p>
                            
                            <p class="mt-1 mb-4 text-error font-weight-medium animate-fade-in" v-if="tips">{{ tips }}</p>
                        </v-card-text>

                        <v-card-text class="pa-0">
                            <v-form @submit.prevent="login">
                                <v-row dense>
                                    <v-col cols="12" v-if="isInternalAuthEnabled()">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            density="default"
                                            type="text"
                                            autocomplete="username"
                                            :autofocus="true"
                                            :disabled="show2faInput || loggingInByPassword || loggingInByOAuth2 || verifying"
                                            :label="tt('Username or Email')"
                                            prepend-inner-icon="mdi-account-outline"
                                            v-model.trim="username"
                                            @input="tempToken = ''"
                                            @keyup.enter="passwordInput?.focus()"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    <v-col cols="12" v-if="isInternalAuthEnabled()">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            density="default"
                                            autocomplete="current-password"
                                            ref="passwordInput"
                                            type="password"
                                            :disabled="show2faInput || loggingInByPassword || loggingInByOAuth2 || verifying"
                                            :label="tt('Password')"
                                            prepend-inner-icon="mdi-lock-outline"
                                            v-model="password"
                                            @input="tempToken = ''"
                                            @keyup.enter="login"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    
                                    <v-col cols="12" v-show="show2faInput" class="animate-fade-in">
                                        <v-text-field
                                            variant="outlined"
                                            density="default"
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
                                            class="modern-input"
                                        />
                                        <v-text-field
                                            variant="outlined"
                                            density="default"
                                            type="text"
                                            :disabled="loggingInByPassword || loggingInByOAuth2 || verifying"
                                            :label="tt('Backup Code')"
                                            prepend-inner-icon="mdi-key-outline"
                                            :append-inner-icon="mdiOnepassword"
                                            v-model="backupCode"
                                            @click:append-inner="twoFAVerifyType = 'passcode'"
                                            @keyup.enter="verify"
                                            v-if="twoFAVerifyType === 'backupcode'"
                                            class="modern-input"
                                        />
                                    </v-col>

                                    <v-col cols="12" class="d-flex align-center justify-space-between mb-6">
                                        <router-link class="text-decoration-none text-body-2 text-primary font-weight-bold h-link" to="/forgetpassword"
                                                     :class="{ 'disabled': !isUserForgetPasswordEnabled() || loggingInByPassword || loggingInByOAuth2 || verifying }">
                                            {{ tt('Forgot password?') }}
                                        </router-link>
                                    </v-col>

                                    <v-col cols="12">
                                        <v-btn block color="primary" size="x-large" elevation="0" class="premium-btn text-none"
                                               :disabled="inputIsEmpty || loggingInByPassword || loggingInByOAuth2 || verifying"
                                               @click="login" v-if="isInternalAuthEnabled() && !show2faInput">
                                            <span>{{ tt('Log In') }}</span>
                                            <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="loggingInByPassword"></v-progress-circular>
                                        </v-btn>
                                        <v-btn block color="primary" size="x-large" elevation="0" class="premium-btn text-none"
                                               :disabled="twoFAInputIsEmpty || loggingInByPassword || loggingInByOAuth2 || verifying"
                                               @click="verify" v-else-if="isInternalAuthEnabled() && show2faInput">
                                            <span>{{ tt('Verify & Continue') }}</span>
                                            <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="verifying"></v-progress-circular>
                                        </v-btn>

                                        <template v-if="isOAuth2Enabled()">
                                            <div class="d-flex align-center my-8">
                                                <v-divider />
                                                <span class="px-4 text-caption text-medium-emphasis text-uppercase">{{ tt('or') }}</span>
                                                <v-divider />
                                            </div>

                                            <v-btn block variant="outlined" color="primary" size="x-large" class="premium-btn-outline text-none"
                                                   :disabled="show2faInput || loggingInByPassword || loggingInByOAuth2 || verifying" :href="oauth2LoginUrl"
                                                   @click="loggingInByOAuth2 = true">
                                                <v-icon start icon="mdi-account-circle-outline" />
                                                <span>{{ oauth2LoginDisplayName }}</span>
                                                <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="loggingInByOAuth2"></v-progress-circular>
                                            </v-btn>
                                        </template>
                                    </v-col>

                                    <v-col cols="12" class="text-center text-body-1 mt-8" v-if="isInternalAuthEnabled()">
                                        <span class="text-medium-emphasis">{{ tt('New to Split Chill AI?') }}</span>
                                        <router-link class="text-primary font-weight-bold text-decoration-none ms-2 h-link" to="/signup"
                                                     :class="{ 'disabled': !isUserRegistrationEnabled() || loggingInByPassword || loggingInByOAuth2 || verifying }">
                                            {{ tt('Create an account') }}
                                        </router-link>
                                    </v-col>
                                </v-row>
                            </v-form>
                        </v-card-text>
                        
                        <div class="mt-10 pt-4 border-t">
                             <v-row no-gutters align="center">
                                <v-col cols="6">
                                    <language-select-button :disabled="loggingInByPassword || loggingInByOAuth2 || verifying" variant="text" size="small" />
                                </v-col>
                                <v-col cols="6" class="text-end">
                                    <a href="javascript:void(0);"
                                       class="text-decoration-none text-caption text-primary font-weight-medium h-link"
                                       :class="{ 'disabled': loggingInByPassword || loggingInByOAuth2 || verifying }"
                                       @click="showMobileQrCode = true">
                                        {{ tt('Mobile App') }}
                                    </a>
                                </v-col>
                                <v-col cols="12" class="text-center text-caption text-medium-emphasis mt-4">
                                     <span>{{ tt('global.app.title') }} v{{ version }} &copy; {{ new Date().getFullYear() }}</span>
                                </v-col>
                             </v-row>
                        </div>
                    </v-card>
                </div>
            </v-col>
        </v-row>

        <switch-to-mobile-dialog v-model:show="showMobileQrCode" />
        <snack-bar ref="snackbar" />
    </div>
</template>

<script setup lang="ts">
import { VTextField } from 'vuetify/components/VTextField';
import SnackBar from '@/components/desktop/SnackBar.vue';
import HelixHero from '@/components/desktop/HelixHeroSimple.vue';

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
            box-shadow: 0 0 0 4px var(--ebk-primary-glow) !important;
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
        box-shadow: 0 10px 25px rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.4) !important;
    }

    &:active:not(:disabled) {
        transform: translateY(0);
    }
}

.premium-btn-outline {
    border-radius: 12px !important;
    border: 2px solid rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.1) !important;
    font-family: 'Outfit', sans-serif !important;
    font-weight: 600 !important;
    transition: all 0.2s ease !important;

    &:hover:not(:disabled) {
        border-color: var(--ebk-primary) !important;
        background-color: rgba(var(--ebk-primary-h), var(--ebk-primary-s), var(--ebk-primary-l), 0.02) !important;
    }
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
    border-top: 1px solid var(--ebk-border-color);
}

.v-theme--dark {
    .modern-input :deep(.v-field--focused) {
        background-color: #1e293b !important;
    }
}
</style>
