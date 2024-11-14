<script lang="ts">
	import { X } from 'lucide-svelte';
	import { flip } from 'svelte/animate';
	import { quintOut } from 'svelte/easing';
	import { crossfade } from 'svelte/transition';
	import { Checkbox } from './ui/checkbox';
	import { Label } from './ui/label';
	import type { Todo } from '$lib/todo-type';

	// 定义 props 的类型
	interface Props {
		todos: Todo[];
		deleteTodo: (id: string) => void;
	}

	// todo list
	// 使用 $bindable 绑定 props，这样可以在组件内部修改 todos 的值
	let { todos = $bindable<Todo[]>([]), deleteTodo }: Props = $props();

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

{#if todos && todos.length > 0}
	<div class="mt-4 flex flex-col gap-2">
		{#each todos.reverse() as todo (todo.id)}
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
					<X size={16} />
				</button>
			</div>
		{/each}
	</div>
{/if}
