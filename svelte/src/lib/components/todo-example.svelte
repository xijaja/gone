<script lang="ts">
	import { flip } from 'svelte/animate';
	import { quintOut } from 'svelte/easing';
	import { crossfade } from 'svelte/transition';
	import { toast } from 'svelte-sonner';
	import { Button } from './ui/button';
	import { Checkbox } from './ui/checkbox';
	import Input from './ui/input/input.svelte';
	import { Label } from './ui/label';
	import { X } from 'lucide-svelte';

	// 定义 todo 的类型
	interface Todo {
		id: string; // 唯一标识
		content: string; // 内容
		done: boolean; // 是否完成
	}

	// 定义 todo 的初始值
	const initialTodos: Todo[] = [
		{ id: '1', content: '君子藏器于身，待时而动', done: true },
		{ id: '2', content: '人有冲天之志，非运不能自通', done: true },
		{ id: '3', content: '择一业，谋食养命', done: true },
		{ id: '4', content: '等一运，扭转乾坤', done: false },
		{ id: '5', content: '会当凌绝顶，一览众山小', done: false },
		{ id: '6', content: '树深时见鹿，溪午不闻钟', done: false },
		{ id: '7', content: '长风破浪会有时，直挂云帆济沧海', done: false }
	];

	// 新的 todo 内容
	let newTodo = $state('');

	// todos 列表
	let todos = $state<Todo[]>(initialTodos);

	// 添加对 input 元素的引用
	let inputElement = $state<HTMLInputElement | null>(null);

	// 添加 todo 的方法
	function addTodo(e: SubmitEvent) {
		e.preventDefault(); // 阻止表单默认提交行为
		if (newTodo.trim() === '') {
			toast('Todo content cannot be empty');
			return;
		}
		todos = [...todos, { id: crypto.randomUUID(), content: newTodo, done: false }];
		newTodo = ''; // 清空输入框
		if (inputElement) inputElement.focus(); // 自动获取光标
	}

	// 删除 todo 的方法
	function deleteTodo(id: string) {
		todos = todos.filter((todo) => todo.id !== id);
	}

	// 创建 crossfade 动画
	const [send, receive] = crossfade({
		duration: 400,
		fallback(node, params) {
			const style = getComputedStyle(node); // 获取节点的样式
			const transform = style.transform === 'none' ? '' : style.transform; // 获取 transform 属性
			return {
				duration: 400, // 动画持续时间
				easing: quintOut, // 动画缓动函数
				css: (t) => `
					transform: ${transform} scale(${t});
					opacity: ${t}
				` // 动画样式，t 是动画进度
			};
		}
	});
</script>

<form class="mb-10 flex w-full gap-2" onsubmit={addTodo}>
	<Input id="new-todo" name="new-todo" type="text" bind:value={newTodo} bind:ref={inputElement} />
	<Button type="submit" class="w-1/5">Add</Button>
</form>

<div class="grid grid-cols-2 gap-4">
	<div>
		<h2 class="text-lg font-bold">Todo List</h2>
		<div class="mt-4 flex flex-col gap-2">
			{#each todos.filter((todo) => !todo.done).reverse() as todo (todo.id)}
				<div
					class="hover:bg-muted group flex items-center gap-2 rounded-md p-2 transition-colors"
					in:receive={{ key: todo.id }}
					out:send={{ key: todo.id }}
					animate:flip
				>
					<Checkbox id={todo.id} bind:checked={todo.done} />
					<Label
						id={todo.id}
						for={todo.id}
						class="break-all text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					>
						{todo.content}
					</Label>
					<button
						class="text-muted-foreground ml-auto hidden size-4 cursor-pointer hover:text-red-500 group-hover:block"
						onclick={() => deleteTodo(todo.id)}
					>
						<X />
					</button>
				</div>
			{/each}
		</div>
	</div>

	<div>
		<h2 class="text-lg font-bold">Done List</h2>
		<div class="mt-4 flex flex-col gap-2">
			{#each todos.filter((todo) => todo.done).reverse() as todo (todo.id)}
				<div
					class="hover:bg-muted group flex items-center gap-2 rounded-md p-2 transition-colors"
					in:receive={{ key: todo.id }}
					out:send={{ key: todo.id }}
					animate:flip
				>
					<Checkbox id={todo.id} bind:checked={todo.done} />
					<Label
						id={todo.id}
						for={todo.id}
						class="break-all text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					>
						{todo.content}
					</Label>
					<button
						class="text-muted-foreground ml-auto hidden size-4 cursor-pointer hover:text-red-500 group-hover:block"
						onclick={() => deleteTodo(todo.id)}
					>
						<X />
					</button>
				</div>
			{/each}
		</div>
	</div>
</div>
