// 请求路由
const baseUrl = "http://127.0.0.1:3030/api";

export const apiPath = {
  // 获取全部待办事项
  getAllTodos: `${baseUrl}/todos/all`,
  // 更新或添加
  postOneTodo: `${baseUrl}/todos/one`,
  // 删除待办事项
  deleteTodoById: `${baseUrl}/todos/`
}

// 请求返回的数据
export interface FetchResponse {
  code: number; // 状态码 2000 表示成功
  data: any; // 返回的数据
  msg: string; // 返回信息
  rid: string; // 请求唯一标识
}

export async function callApi(url: string, options?: RequestInit): Promise<FetchResponse> {
  // 添加默认的请求头
  options = {
    headers: {'Content-Type': 'application/json'}, // 默认请求头
    credentials: 'include', // 携带 cookie
    mode: 'cors', // 允许跨域请求
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
    document.cookie = ''; // 清空 cookie
    window.location.href = '/login'; // 重定向跳转登录页
  }
  // 返回数据
  return jsonData;
}

// 添加 GET 方法
callApi.get = async function (url: string, params?: Record<string, any>): Promise<FetchResponse> {
  const queryString = params ? '?' + new URLSearchParams(params).toString() : '';
  return callApi(url + queryString, {method: 'GET'});
};

// 添加 POST 方法
callApi.post = async function (url: string, body?: Record<string, any>): Promise<FetchResponse> {
  return callApi(url, {method: 'POST', body: JSON.stringify(body)});
};

// 添加 DELETE 方法
callApi.delete = async function (url: string): Promise<FetchResponse> {
  return callApi(url, {method: 'DELETE'});
};