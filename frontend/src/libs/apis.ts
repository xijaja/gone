// 请求路由
const baseUrl = "http://127.0.0.1:3030/api";

// 请求获取全部待办事项
export async function getTodos() {
  const response = await fetch(`${baseUrl}/todos/all`);
  const jsonStr = await response.json();
  console.log(jsonStr);
  console.log(jsonStr.data.list);
  return jsonStr;
}

// 更新或添加
export async function postTodo(data: {
  id?: number;
  title: string;
  done?: boolean;
}) {
  const response = await fetch(`${baseUrl}/todos/one`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    mode: "cors",
    body: JSON.stringify(data),
  });
  return await response.json();
}

// 删除待办事项
export async function deleteTodoById(id: number) {
  const response = await fetch(`${baseUrl}/todos/${id}`, {
    method: "DELETE",
  });
  return await response.json();
}
