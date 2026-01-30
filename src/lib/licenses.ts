export function getLicense(): string {
    return __SPLITCHILL_AI_LICENSE__;
}

export function getThirdPartyLicenses(): LicenseInfo[] {
    return __SPLITCHILL_AI_THIRD_PARTY_LICENSES__ || [];
}
