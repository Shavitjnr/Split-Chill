export interface ForgetPasswordRequest {
    readonly email: string;
}

export interface PasswordResetByOTPRequest {
    readonly email: string;
    readonly otp: string;
    readonly password: string;
}

