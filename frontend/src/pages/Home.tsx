import Layout from "../layouts/Layout";
import { getTodos, postTodo, deleteTodoById } from "../libs/apis";

type todo = {
  id?: number;
  title: string;
  done?: boolean;
}

export default function HomePage() {
  // 获取
  const [todos, { refetch }] = createResource(getTodos); // 获取全部待办
  // 新增
  const [title, setTitle] = createSignal<string>(""); // 输入框的值
  const [todo, setTodo] = createSignal<todo>(); // 新增待办
  const [addTodo] = createResource(todo, postTodo); // 新增待办
  // 删除
  const [id, setId] = createSignal<number>(); // 待删除的 id
  const [deleteTodo] = createResource(id, deleteTodoById); // 删除待办

  // 监听新增待办的变化
  createEffect(() => {
    if (addTodo()) {
      refetch(); // 重新获取全部待办数据
      setTitle(""); // 清空 title 的值
      const input = document.querySelector("input");
      if (input) input.value = ""; // 把输入框的内容也清空
    }
    if (deleteTodo()) {
      refetch(); // 重新获取全部待办数据
    }
  })

  return (
    <Layout>
      <div class="flex gap-4 justify-center">
        <input
          type="text"
          required // 必填项
          placeholder="还有什么事是待办的呢？"
          class="input input-bordered w-full max-w-xs"
          onChange={(e) => setTitle(e.currentTarget.value)}
        // onInput={(e) => setTitle(e.currentTarget.value)}
        />
        <button class="btn" onclick={() => setTodo({ title: title() })}>
          添加事项
        </button>
      </div>

      <div class="mx-2 my-8 p-2">
        {todos.loading && <div>加载中...</div>}
        {todos.error && <div>加载失败</div>}
        <Show when={todos()}>
          <ul class="space-y-2">
            <For each={todos().data.list as todo[]}>
              {(todo) => (
                <li class="p-2 border dark:border-gray-600 rounded-lg flex justify-between items-center">
                  <div class="flex items-center gap-4">
                    <span class="badge badge-lg">{todo.id}</span>
                    <span>{todo.title}</span>
                  </div>
                  <button class="btn btn-sm" onclick={() => setId(todo.id!)}>删除</button>
                </li>
              )}
            </For>
          </ul>
        </Show>
      </div>
    </Layout>
  );
}