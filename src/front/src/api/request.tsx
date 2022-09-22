import type { AxiosInstance, AxiosResponse } from 'axios';
import axios from 'axios';
// import { Message } from '@b-design/ui';
// import { handleError } from '../utils/errors';
// import { getToken } from '../utils/storage';
// import { authenticationRefreshToken } from './productionLink';
// import ResetLogin from '../utils/resetLogin';

type RetryRequests = (token: string) => void;
let isRefreshing = false;
let retryRequests: RetryRequests[] = [];

export const axiosInstance: AxiosInstance = axios.create({
  baseURL: process.env.BLOG_DOMAIN,
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
    if ([304].includes(response.status)){
        removeToken
        
        return
    }
    if (response.status === 200) {
      return response;
    }
  },

//   async (error: any) => {
//     if (!error.response) {
//       return Promise.reject(error);
//     }
//     const { data, config } = error.response;
//     if (data.BusinessCode === 12002) {
//       if (!isRefreshing) {
//         isRefreshing = true;
//         return getRefreshTokenFunc()
//           .then((res: any) => {
//             const refreshData = res && res.data;
//             if (refreshData && refreshData.accessToken) {
//               localStorage.setItem('token', refreshData.accessToken);
//               localStorage.setItem('refreshToken', refreshData.refreshToken);
//               config.headers.Authorization = localStorage.getItem("AccessToken");
//               retryRequests.forEach((cb) => {
//                 // cb(getToken());
//               });
//               retryRequests = [];
//               return axiosInstance(config);
//             }
//           })
//           .catch(() => {
//             // return ResetLogin.getInstance().reset;
//           })
//           .finally(() => {
//             isRefreshing = false;
//           });
//       } else {
//         return new Promise((resolve) => {
//           retryRequests.push((token: string) => {
//             config.headers.Authorization = 'Bearer ' + token;
//             resolve(axiosInstance(config));
//           });
//         });
//       }
//     } else if (data.BusinessCode === 12010 || data.BusinessCode === 12004) {
//     //   return ResetLogin.getInstance().reset;
//     } else {
//       return Promise.reject(error.response || error);
//     }
//   },
);

axiosInstance.interceptors.request.use(
  (config: any) => {
    if (getToken()) {
      config.headers.Authorization = 'Bearer ' + getToken();
    }
    return config;
  },
  (error: any) => {
    return Promise.reject(error);
  },
);

const handleAPIError = (err: any, customError: boolean) => {
  const { data, status } = err;
  if (customError) {
    throw data;
  } else if (data && data.BusinessCode) {
    // handleError(data);
  } else {
    // Message.error(getMessage(status));
  }
};

export const post = (url: string, params: any, customError?: boolean) => {
  return axiosInstance
    .post(url, params)
    .then((res) => {
      return res && res.data;
    })
    .catch((err) => {
      handleAPIError(err, params.customError || customError);
    });
};

export const get = (url: string, params: any) => {
  return axiosInstance
    .get(url, params)
    .then((res) => {
      return res && res.data;
    })
    .catch((err) => {
      handleAPIError(err, params.customError);
    });
};

export const rdelete = (url: string, params: any, customError?: boolean) => {
  return axiosInstance
    .delete(url, params)
    .then((res) => {
      return res && res.data;
    })
    .catch((err) => {
      handleAPIError(err, params.customError || customError);
    });
};

export const put = (url: string, params: any, customError?: boolean) => {
  return axiosInstance
    .put(url, params)
    .then((res) => {
      return res && res.data;
    })
    .catch((err) => {
      handleAPIError(err, params.customError || customError);
    });
};

async function getRefreshTokenFunc() {
  const refreshToken = localStorage.getItem('refreshToken') || '';
  return await axios({
    // url: `${baseURL}${authenticationRefreshToken}`,
    method: 'GET',
    headers: {
      RefreshToken: refreshToken,
    },
  });
}


const getToken = () => {
   return localStorage.getItem("AccessToken") || ''
}
const removeToken = () => {
    localStorage.removeItem("AccessToken")
    window.location.href="/signIn"
}

const getMessage = (status: number | string): string => {
    let message: string = '';
    switch (status) {
      case 400:
        message = 'BadRequest(400)';
        break;
      case 401:
        message = 'Unauthorized(401)';
        break;
      case 403:
        message = 'Forbidden(403)';
        break;
      case 404:
        message = 'NotFound(404)';
        break;
      case 408:
        message = 'TimeOut(408)';
        break;
      case 500:
        message = 'InternalServerError(500)';
        break;
      case 501:
        message = 'ServerNotImplements(501)';
        break;
      case 502:
        message = 'GatewayError(502)';
        break;
      case 503:
        message = 'InternalServerUnavailable(503)';
        break;
      case 504:
        message = 'GatewayTimeout(504)';
        break;
      default:
        message = `connect error(${status})!`;
    }
    return `${message}, please check the network or contact the administrator!`;
  };