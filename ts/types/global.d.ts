declare global {
    interface Window {
        RegisterPage: RegisterPageApi,
        SearchPage: SearchPageApi,
    };
}
export interface RegisterPageApi {
    init: (baseUrl: string) => void,
    save: (baseUrl: string) => void,
}
export interface SearchPageApi {
    search: (baseUrl: string) => void,
}