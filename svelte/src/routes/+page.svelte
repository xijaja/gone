<script>
	import { Button } from '$lib/components/ui/button';
	import { toast } from 'svelte-sonner';
</script>

<h1>Welcome to GONE</h1>
<p>
	Visit <Button variant="link" href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</Button> to read
	the documentation
</p>

<Button onclick={() => toast('Hello world')}>Show toast</Button>
