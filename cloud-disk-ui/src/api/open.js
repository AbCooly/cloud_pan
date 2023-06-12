//不需要获取授权的接口
import http from '../utils/http'

/**
 *  @parms resquest 请求地址 例如：http://197.82.15.15:8088/request/...
 *  @param '/testIp'代表vue-cil中config，index.js中配置的代理
 */
let resquest = "/api"

//登录
export function login(params){
    return http.post(`${resquest}/login`,params)
}

// 获取用户数据
export function registerCode(params){
    return http.get(`${resquest}/registerCode`,params)
}

export function register(params){
    return http.post(`${resquest}/register`,params)
}

