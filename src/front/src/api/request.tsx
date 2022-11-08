import type {AxiosInstance, AxiosResponse} from 'axios';
import axios from 'axios';
import {MessageData, showMessage} from '../pages/component/message';

export const axiosInstance: AxiosInstance = axios.create({
    baseURL: process.env.REACT_APP_BLOG_DOMAIN,
    timeout: 30000,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
    },
    transformRequest: [
        function (data) {
            return JSON.stringify(data);
        },
    ],
});

axiosInstance.interceptors.response.use(
    (response: AxiosResponse) => {
        console.log("response",response,response.headers.Authorization)
        if (response.headers.Authorization) {
            setToken(response.headers.Authorization)
        }
        if (response.status === 200) {
            return response;
        }
    },

    async (error: any) => {
        if (!error.response) {
            return Promise.reject(error)
        }
        const data = error.response
        if ([403, 401].includes(data.status)) {
            removeToken()
            return Promise.reject(error)
        }

    }
);

axiosInstance.interceptors.request.use(
    (config: any) => {
        if (getToken()) {
            config.headers.Authorization = getToken();
        }
        return config;
    },
    (error: any) => {
        return Promise.reject(error);
    },
);

const handleAPIError = (err: any) => {
    showMessage(new MessageData(true, "error", "请求失败", err))
};

export const post = (url: string, params: any) => {
    return axiosInstance
        .post(url, params)
        .then((res) => {
            if (res.data.code !== 0) {
                return Promise.reject(res.data.msg)
            }
            return res.data;
        })
        .catch((err) => {
            handleAPIError(err);
        });
};

export const get = (url: string, params: any) => {
    return axiosInstance
        .get(url, params)
        .then((res) => {
            return res && res.data;
        })
        .catch((err) => {
            handleAPIError(err);
        });
};

export const rdelete = (url: string, params: any) => {
    return axiosInstance
        .delete(url, params)
        .then((res) => {
            return res && res.data;
        })
        .catch((err) => {
            handleAPIError(err);
        });
};

export const put = (url: string, params: any) => {
    return axiosInstance
        .put(url, params)
        .then((res) => {
            return res && res.data;
        })
        .catch((err) => {
            handleAPIError(err);
        });
};

const setToken = (token: string) => {
    localStorage.setItem("AccessToken", token)
}
const getToken = () => {
    return localStorage.getItem("AccessToken") || ''
}
const removeToken = () => {
    localStorage.removeItem("AccessToken")
    window.location.href = "/signIn"
}

export const postGraphql = (url: string, params: any) => {
    return axiosInstance
        .post(url, params)
        .then(res => {
            if (res.data.errors) {
                console.log("aaa")
                return Promise.reject(res.data.errors[0].message)
            }
            const pgdata = res.data.data[Object.keys(res.data.data)[0]]
            if (pgdata.code !== 0) {
                return Promise.reject(pgdata.msg)
            }
            return pgdata
        })
        .catch((err) => {
            handleAPIError(err);
            return new Promise((resolve, reject) => {})
        });
};