import { post } from "./request"

export function postSignIn(param: { username: string, password: string }) {
    return post("/signIn", param)
}

export function postSignUp(param: { username: string, password: string }) {
    return post("/signUp", param)
}

export function postChangePwd(param: { oldPwd: string, newPwd: string }) {
    return post("/changePwd", param)
}

export function  postUserInfo() {
    return post("/info", {})
}