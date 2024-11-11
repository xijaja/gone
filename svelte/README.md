# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## shadcn-svelte

[shadcn-svelte](https://github.com/shadcn-svelte/ui) 是一个基于 ShadcnUI 的 Svelte 组件库。

安装组件：

```bash
pnpm dlx shadcn-svelte@next add button
```

使用组件：

```svelte
<script lang="ts">
	import { Button } from '$lib/components/ui/button';
</script>

<Button>Click me</Button>
```
