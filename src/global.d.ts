declare const __SPLITCHILL_AI_IS_PRODUCTION__: boolean;
declare const __SPLITCHILL_AI_VERSION__: string;
declare const __SPLITCHILL_AI_BUILD_UNIX_TIME__: string;
declare const __SPLITCHILL_AI_BUILD_COMMIT_HASH__: string;
declare const __SPLITCHILL_AI_LICENSE__: string;
declare const __SPLITCHILL_AI_CONTRIBUTORS__: ContributorInfo;
declare const __SPLITCHILL_AI_THIRD_PARTY_LICENSES__: LicenseInfo[];

declare interface ContributorInfo {
    code: string[];
    translators: Record<string, string[]>;
}

declare interface LicenseInfo {
    name: string;
    copyright?: string;
    url?: string;
    license?: string;
    licenseUrl?: string;
}

interface Window {
    SPLITCHILL_AI_SERVER_SETTINGS?: {
        [key: string]: string | number | boolean | undefined | null;
    };
}

interface Navigator {
    browserLanguage?: string;
}

interface Credential {
    rawId: ArrayBuffer;
    response: {
        clientDataJSON: ArrayBuffer;
        attestationObject: ArrayBuffer;
        userHandle: ArrayBuffer;
    };
}
