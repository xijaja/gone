/* eslint-disable @typescript-eslint/no-explicit-any */
// 请求路由
const baseUrl = import.meta.env.MODE === "production" ? "/api" : "http://127.0.0.1:3030/api";

// 请求路由
export const apiUrl = {
  // CSRF token
  csrfToken: `${baseUrl}/csrf-token`,

  // 示例：登录
  login: `${baseUrl}/v1/login`,

  // 示例：用户信息
  userInfo: (uid: string) => `${baseUrl}/v1/user/${uid}/info`,
}

// 请求返回的数据
export interface FetchResponse {
  code: number; // 状态码 2000 表示成功
  data: any; // 返回的数据
  msg: string; // 返回信息
  rid: string; // 请求唯一标识
}

// 存储CSRF token
let csrfToken: string | null = null;

// 获取CSRF token
async function getCsrfToken(): Promise<string | null> {
  try {
    const response = await fetch(apiUrl.csrfToken, {
      method: 'GET',
      credentials: 'include'
    });
    const data = await response.json();
    if (data.code === 2000 && data.data.token) {
      return data.data.token;
    }
    console.error('获取CSRF token失败:', data);
    return null;
  } catch (error) {
    console.error('获取CSRF token出错:', error);
    return null;
  }
}

// 请求方法
export async function apiFetch(url: string, options?: RequestInit): Promise<FetchResponse> {
  // 获取CSRF token(如果需要且未缓存)
  if (!csrfToken && options?.method !== 'GET') {
    csrfToken = await getCsrfToken();
  }

  // 合并默认配置和自定义配置
  const defaultHeaders: Record<string, string> = {
    Authorization: `Bearer ${localStorage.getItem("token")}`
  };

  // 如果不是 FormData，添加 JSON content-type
  if (!(options?.body instanceof FormData)) {
    defaultHeaders['Content-Type'] = 'application/json';
  }

  // 如果有CSRF token且不是GET请求，添加到请求头
  if (csrfToken && options?.method !== 'GET') {
    defaultHeaders['X-CSRF-Token'] = csrfToken;
  }

  options = {
    credentials: 'include',
    mode: 'cors',
    headers: {
      ...defaultHeaders,
      ...options?.headers,
    },
    ...options,
  };

  // 发送请求
  const response = await fetch(url, options);
  const jsonData = await response.json();
  // 非生产环境打印请求信息
  if (import.meta.env.MODE !== "production") {
    console.log("PATH:", url);
    console.log("BODY:", options);
    console.log("RESP:", jsonData);
    console.log("----- ----- -----");
  }
  // 返回拦截
  if (jsonData.code === 3000) {
    localStorage.removeItem("token"); // 清空本地存储中的 token
    localStorage.removeItem("role"); // 清空本地存储中的 role
    window.location.href = '/login'; // 重定向跳转登录页
  }
  // 返回数据
  return jsonData;
}

// 添加 GET 方法
apiFetch.get = async function (url: string, params?: Record<string, any>): Promise<FetchResponse> {
  const queryString = params ? '?' + new URLSearchParams(params).toString() : '';
  return apiFetch(url + queryString, { method: 'GET' });
};

// 添加 POST 方法
apiFetch.post = async function (url: string, body?: Record<string, any> | FormData): Promise<FetchResponse> {
  const options: RequestInit = {
    method: 'POST',
    body: body instanceof FormData ? body : JSON.stringify(body)
  };

  return apiFetch(url, options);
};

// 添加 DELETE 方法
apiFetch.delete = async function (url: string, body?: Record<string, any>): Promise<FetchResponse> {
  return apiFetch(url, { method: 'DELETE', body: JSON.stringify(body) });
};
