<script lang="ts">
  import './layout.css';
  import { ModeWatcher } from 'mode-watcher';
  import { Toaster } from "$lib/components/ui/sonner/index.js";

	let { children } = $props();
</script>

<svelte:head>
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<title>{import.meta.env.APP_NAME}</title>
	<meta name="description" content={import.meta.env.APP_DESCRIPTION} />
	<meta name="author" content="MBoukhal" />
</svelte:head>

<ModeWatcher />

<!-- <QueryClientProvider client={queryClient}> -->
{@render children()}
<Toaster position="top-center" richColors />
<!-- </QueryClientProvider> -->
