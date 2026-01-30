import { isEnableDebug } from './settings.ts';

function logDebug(msg: string, obj?: unknown): void {
    if (isEnableDebug()) {
        if (obj) {
            console.debug('[Split Chill AI Debug] ' + msg, obj);
        } else {
            console.debug('[Split Chill AI Debug] ' + msg);
        }
    }
}

function logInfo(msg: string, obj?: unknown): void {
    if (obj) {
        console.info('[Split Chill AI Info] ' + msg, obj);
    } else {
        console.info('[Split Chill AI Info] ' + msg);
    }
}

function logWarn(msg: string, obj?: unknown): void {
    if (obj) {
        console.warn('[Split Chill AI Warn] ' + msg, obj);
    } else {
        console.warn('[Split Chill AI Warn] ' + msg);
    }
}

function logError(msg: string, obj?: unknown): void {
    if (obj) {
        console.error('[Split Chill AI Error] ' + msg, obj);
    } else {
        console.error('[Split Chill AI Error] ' + msg);
    }
}

export default {
    debug: logDebug,
    info: logInfo,
    warn: logWarn,
    error: logError
};
