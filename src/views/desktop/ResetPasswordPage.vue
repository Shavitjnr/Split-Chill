<template>
    <div class="clean-recovery-page">
        <div class="recovery-container">
            <!-- Brand Logo & Header matching landing.html & LoginPage.vue -->
            <div class="recovery-brand mb-6 text-center">
                <router-link to="/" class="brand-link">
                    <div class="brand-icon-box">
                        <v-icon icon="mdi-call-split" class="brand-icon" size="24" color="white" />
                    </div>
                    <div class="d-flex flex-column text-start">
                        <div class="d-flex align-center gap-2">
                            <h1 class="recovery-brand-title">{{ tt('global.app.title') }}</h1>
                            <span class="brand-pulse-dot"></span>
                        </div>
                        <span class="brand-tagline">Expense Harmony</span>
                    </div>
                </router-link>
            </div>

            <!-- Centered Luxury Card -->
            <v-card class="recovery-card" elevation="0">
                <div class="d-flex justify-center mb-4">
                    <div class="recovery-badge-pill">
                        <span class="badge-dot"></span>
                        <span>Direct Token Reset • Split Chill 2.0</span>
                    </div>
                </div>

                <div class="recovery-card-header text-center mb-6">
                    <h2 class="recovery-title">{{ tt('Reset Password') }}</h2>
                    <p class="recovery-subtitle">{{ tt('Confirm your email address and enter your new secure password.') }}</p>
                </div>

                <v-form @submit.prevent="resetPassword">
                    <div class="mb-4">
                        <label class="input-field-label d-block mb-1">
                            <span>{{ tt('E-mail Address') }}</span>
                        </label>
                        <v-text-field
                            variant="outlined"
                            color="primary"
                            type="email"
                            autocomplete="email"
                            :autofocus="true"
                            :disabled="updating"
                            :placeholder="tt('yourname@example.com')"
                            :prepend-inner-icon="mdiEmailOutline"
                            v-model.trim="email"
                            class="modern-input"
                        />
                    </div>

                    <div class="mb-4">
                        <label class="input-field-label d-block mb-1">
                            <span>{{ tt('New Password') }}</span>
                        </label>
                        <v-text-field
                            variant="outlined"
                            color="primary"
                            :type="showPassword ? 'text' : 'password'"
                            autocomplete="new-password"
                            :disabled="updating"
                            :placeholder="tt('At least 6 characters')"
                            :prepend-inner-icon="mdiLockOutline"
                            :append-inner-icon="showPassword ? mdiEyeOffOutline : mdiEyeOutline"
                            @click:append-inner="showPassword = !showPassword"
                            v-model="newPassword"
                            class="modern-input"
                        />
                    </div>

                    <div class="mb-5">
                        <label class="input-field-label d-block mb-1">
                            <span>{{ tt('Confirm New Password') }}</span>
                        </label>
                        <v-text-field
                            variant="outlined"
                            color="primary"
                            :type="showConfirmPassword ? 'text' : 'password'"
                            autocomplete="new-password"
                            :disabled="updating"
                            :placeholder="tt('Re-enter new password')"
                            :prepend-inner-icon="mdiLockOutline"
                            :append-inner-icon="showConfirmPassword ? mdiEyeOffOutline : mdiEyeOutline"
                            @click:append-inner="showConfirmPassword = !showConfirmPassword"
                            v-model="confirmPassword"
                            @keyup.enter="resetPassword"
                            class="modern-input"
                        />
                    </div>

                    <v-btn block size="large" elevation="0" class="premium-btn text-none"
                           :disabled="!email || !newPassword || !confirmPassword || updating" @click="resetPassword">
                        <span>{{ tt('Update Password') }}</span>
                        <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="updating"></v-progress-circular>
                    </v-btn>

                    <div class="text-center mt-6">
                        <router-link class="return-to-login-link d-inline-flex align-center" to="/login"
                                     :class="{ 'disabled': updating }">
                            <v-icon :icon="mdiArrowLeft" size="16" class="me-1"/>
                            <span>{{ tt('Back to login page') }}</span>
                        </router-link>
                    </div>

                    <div class="card-footer-bar mt-6 pt-5 d-flex align-center justify-space-between flex-wrap gap-2">
                        <LanguageSelectButton :disabled="updating" variant="text" size="small" />
                        <span class="footer-version">{{ tt('global.app.title') }} v{{ version }}</span>
                    </div>
                </v-form>
            </v-card>

            <!-- Trust Badge & Back to Landing link matching landing.html -->
            <div class="recovery-trust-badge text-center mt-6">
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
import LanguageSelectButton from '@/components/desktop/LanguageSelectButton.vue';
import SnackBar from '@/components/desktop/SnackBar.vue';

import { ref, computed, useTemplateRef } from 'vue';
import { useRouter } from 'vue-router';

import { useI18n } from '@/locales/helpers.ts';

import { useRootStore } from '@/stores/index.ts';

import { getClientDisplayVersion } from '@/lib/version.ts';

import {
    mdiEmailOutline,
    mdiArrowLeft,
    mdiLockOutline,
    mdiEyeOutline,
    mdiEyeOffOutline
} from '@mdi/js';

type SnackBarType = InstanceType<typeof SnackBar>;

const props = defineProps<{
    token: string;
}>();

const router = useRouter();
const { tt } = useI18n();
const rootStore = useRootStore();

const version = `${getClientDisplayVersion()}`;
const snackbar = useTemplateRef<SnackBarType>('snackbar');

const email = ref<string>('');
const newPassword = ref<string>('');
const confirmPassword = ref<string>('');
const showPassword = ref<boolean>(false);
const showConfirmPassword = ref<boolean>(false);
const updating = ref<boolean>(false);
const passwordChanged = ref<boolean>(false);

const inputProblemMessage = computed<string | null>(() => {
    if (!email.value) {
        return 'Email address cannot be blank';
    } else if (!newPassword.value && !confirmPassword.value) {
        return 'Nothing has been modified';
    } else if (!newPassword.value && confirmPassword.value) {
        return 'New password cannot be blank';
    } else if (newPassword.value && !confirmPassword.value) {
        return 'Password confirmation cannot be blank';
    } else if (newPassword.value && confirmPassword.value && newPassword.value !== confirmPassword.value) {
        return 'Password and password confirmation do not match';
    } else {
        return null;
    }
});

function onSnackbarShowStateChanged(newValue: boolean): void {
    if (!newValue && passwordChanged.value) {
        router.replace('/login');
    }
}

function resetPassword(): void  {
    passwordChanged.value = false;

    const problemMessage = inputProblemMessage.value;

    if (problemMessage) {
        snackbar.value?.showMessage(problemMessage);
        return;
    }

    updating.value = true;

    rootStore.resetPassword({
        token: props.token,
        email: email.value,
        password: newPassword.value
    }).then(() => {
        updating.value = false;
        passwordChanged.value = true;
        snackbar.value?.showMessage('Password has been updated');
    }).catch(error => {
        updating.value = false;
        passwordChanged.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
}
</script>

<style lang="scss" scoped>
/* =========================================================
   SPILL CHILL LUXURY RECOVERY THEME (Aligned with LoginPage & SignupPage)
   Emerald (#10B981) • Forest (#064E3B) • Glassmorphism
   ========================================================= */

@keyframes pulse {
    0%, 100% { opacity: 1; transform: scale(1); }
    50% { opacity: 0.5; transform: scale(0.85); }
}

.clean-recovery-page {
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

:deep(.v-theme--dark) .clean-recovery-page,
.v-theme--dark .clean-recovery-page {
    background:
        radial-gradient(at 15% 15%, rgba(16, 185, 129, 0.15) 0px, transparent 50%),
        radial-gradient(at 85% 20%, rgba(6, 78, 59, 0.28) 0px, transparent 55%),
        radial-gradient(at 50% 80%, rgba(2, 44, 34, 0.35) 0px, transparent 50%),
        #0B0F19 !important;
}

.recovery-container {
    width: 100%;
    max-width: 440px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
}

/* Brand Header */
.recovery-brand {
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

    .recovery-brand-title {
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

:deep(.v-theme--dark) .recovery-brand .recovery-brand-title {
    color: #ffffff !important;
}

/* Badge Pill */
.recovery-badge-pill {
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

:deep(.v-theme--dark) .recovery-badge-pill {
    background: rgba(16, 185, 129, 0.15);
    border-color: rgba(16, 185, 129, 0.4);
    color: #34d399;
}

/* Centered Glass Card */
.recovery-card {
    width: 100%;
    background: rgba(255, 255, 255, 0.94) !important;
    -webkit-backdrop-filter: blur(20px);
    backdrop-filter: blur(20px);
    border: 1px solid rgba(226, 232, 240, 0.9) !important;
    border-radius: 24px !important;
    padding: 36px 32px 28px !important;
    box-shadow: 0 20px 45px -12px rgba(6, 78, 59, 0.1), 0 0 1px 1px rgba(16, 185, 129, 0.12) !important;
}

:deep(.v-theme--dark) .recovery-card,
.v-theme--dark .recovery-card {
    background: rgba(17, 24, 39, 0.90) !important;
    border: 1px solid rgba(255, 255, 255, 0.08) !important;
    box-shadow: 0 20px 45px -12px rgba(0, 0, 0, 0.6) !important;
}

.recovery-card-header {
    .recovery-title {
        font-family: 'Outfit', sans-serif;
        font-size: 26px;
        font-weight: 800;
        letter-spacing: -0.02em;
        color: #0f172a;
        margin: 0 0 6px;
    }

    .recovery-subtitle {
        font-family: 'Plus Jakarta Sans', 'Inter', sans-serif;
        font-size: 13.5px;
        color: #64748b;
        margin: 0;
        line-height: 1.5;
    }
}

:deep(.v-theme--dark) .recovery-card-header .recovery-title {
    color: #ffffff;
}

.input-field-label {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 12.5px;
    font-weight: 600;
    color: #475569;
}

:deep(.v-theme--dark) .input-field-label {
    color: #94a3b8;
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

:deep(.v-theme--dark) .modern-input {
    :deep(.v-field) {
        background: #1e293b !important;
        border-color: #334155 !important;
    }
    :deep(input) {
        color: #f8fafc;
    }
}

/* Submit Buttons with Emerald & Forest Gradient */
.premium-btn {
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

.return-to-login-link {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 13.5px;
    font-weight: 700;
    color: #059669;
    text-decoration: none;
    transition: color 0.15s ease;

    &:hover {
        color: #047857;
        text-decoration: underline;
    }
}

.text-emerald {
    color: #10B981 !important;
}

.card-footer-bar {
    border-top: 1px solid #f1f5f9;

    .footer-version {
        font-family: 'Plus Jakarta Sans', sans-serif;
        font-size: 12px;
        color: #94a3b8;
    }
}

:deep(.v-theme--dark) .card-footer-bar {
    border-top-color: #1e293b;
}

/* Trust Badge & Back Link */
.recovery-trust-badge {
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 12px;
    font-weight: 600;
    color: #64748b;
    display: flex;
    align-items: center;
    justify-content: center;
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
    .recovery-card {
        padding: 28px 20px 20px !important;
        border-radius: 20px !important;
    }
}
</style>
