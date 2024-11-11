<script lang="ts">
	import '../app.css';
	import { Toaster } from '$lib/components/ui/sonner';
	import { ModeWatcher } from 'mode-watcher';
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';

	let { children } = $props();
	let mounted = $state(false);
	onMount(() => {
		mounted = true;
	});
</script>

<svelte:head>
	<title>GONE</title>
</svelte:head>

<!-- 模式切换 -->
{#if browser && mounted}
	<ModeWatcher />
{/if}

<!-- 吐司 -->
<Toaster />

<!-- 子页面渲染 -->
{@render children()}
