import api from './axios';
export function loginapi(request: { email?: string; phone?: string; password: string }) {
    return api.post("/login", request);
}
export function register(email:string, password:string, phone:string, first_name:string, last_name:string) {
    return api.post("/register", {email, password, phone, first_name, last_name});
}
export function checkEmail(email:string) {
    return api.get("/check/email?email=" + email);
}
export function checkPhone(phone:string) {
    return api.get("/check/phone?phone=" + phone);
}