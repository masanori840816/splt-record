declare global {
    interface Window {
        RegisterPage: RegisterPageApi,
    };
}
export interface RegisterPageApi {
    init: (baseUrl: string) => void,
    save: (baseUrl: string) => void,
}