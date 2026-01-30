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
                    title="Account Recovery"
                    description="Lost your key? No worries. We'll help you regain access to your account in a few simple steps."
                />
            </v-col>

            
            <v-col cols="12" md="4" class="d-flex flex-column justify-center align-center px-6">
                <div class="d-flex align-center justify-center w-100 animate-scale-in">
                    <v-card class="glass-card w-100 px-8 pt-12 pb-8" max-width="480">
                         <v-card-text class="text-start px-0">
                            <h2 class="text-h3 font-weight-bold mb-2" style="font-family: 'Outfit', sans-serif;">{{ tt('Recovery') }}</h2>
                            <p class="text-body-1 text-medium-emphasis mb-10">{{ tt('Enter your registered email to receive a password reset link.') }}</p>
                        </v-card-text>

                        <v-card-text class="pa-0">
                            <v-form @submit.prevent="requestResetPassword">
                                <v-row dense>
                                    <v-col cols="12">
                                        <v-text-field
                                            variant="outlined"
                                            color="primary"
                                            type="email"
                                            autocomplete="email"
                                            :autofocus="true"
                                            :disabled="requesting"
                                            :label="tt('E-mail Address')"
                                            prepend-inner-icon="mdi-email-outline"
                                            v-model="email"
                                            @keyup.enter="requestResetPassword"
                                            class="modern-input mb-4"
                                        />
                                    </v-col>

                                    <v-col cols="12">
                                        <v-btn block color="primary" size="x-large" elevation="0" class="premium-btn text-none"
                                               :disabled="!email || requesting" @click="requestResetPassword">
                                            <span>{{ tt('Send Recovery Link') }}</span>
                                            <v-progress-circular indeterminate size="20" width="2" class="ms-3" v-if="requesting"></v-progress-circular>
                                        </v-btn>
                                    </v-col>

                                    <v-col cols="12" class="text-center mt-8">
                                        <router-link class="text-decoration-none text-body-1 text-primary font-weight-bold h-link d-inline-flex align-center" to="/login"
                                                     :class="{ 'disabled': requesting }">
                                            <v-icon :icon="mdiChevronLeft" class="me-1"/>
                                            <span>{{ tt('Return to login') }}</span>
                                        </router-link>
                                    </v-col>
                                </v-row>
                            </v-form>
                        </v-card-text>

                         <div class="mt-12 pt-6 border-t">
                             <v-row no-gutters align="center">
                                <v-col cols="12" class="text-center">
                                    <language-select-button :disabled="requesting" variant="text" size="small" />
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

        <snack-bar ref="snackbar" />
    </div>
</template>

<script setup lang="ts">
import HelixHero from '@/components/desktop/HelixHeroSimple.vue';
import SnackBar from '@/components/desktop/SnackBar.vue';

import { ref, useTemplateRef } from 'vue';
import { useI18n } from '@/locales/helpers.ts';

import { useRootStore } from '@/stores/index.ts';

import { APPLICATION_LOGO_PATH } from '@/consts/asset.ts';
import { getClientDisplayVersion } from '@/lib/version.ts';

import {
    mdiChevronLeft
} from '@mdi/js';

type SnackBarType = InstanceType<typeof SnackBar>;

const { tt } = useI18n();
const rootStore = useRootStore();

const version = `${getClientDisplayVersion()}`;
const snackbar = useTemplateRef<SnackBarType>('snackbar');

const email = ref<string>('');
const requesting = ref<boolean>(false);

function requestResetPassword(): void {
    if (!email.value) {
        snackbar.value?.showMessage('Email address cannot be blank');
        return;
    }

    requesting.value = true;

    rootStore.requestResetPassword({
        email: email.value
    }).then(() => {
        requesting.value = false;
        snackbar.value?.showMessage('Password reset email has been sent');
    }).catch(error => {
        requesting.value = false;

        if (!error.processed) {
            snackbar.value?.showError(error);
        }
    });
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

.v-theme--dark {
    .modern-input :deep(.v-field--focused) {
        background-color: #1e293b !important;
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
