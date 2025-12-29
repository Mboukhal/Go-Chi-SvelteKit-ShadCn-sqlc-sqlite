<script lang="ts">
 import * as Empty from "$lib/components/ui/empty/index.js";
	import { Button } from "@/lib/components/ui/button";
	import Theme from "@/lib/components/ui/theme/theme.svelte";
//  import * as InputGroup from "$lib/components/ui/input-group/index.js";
//  import * as Kbd from "$lib/components/ui/kbd/index.js";
//  import SearchIcon from "@lucide/svelte/icons/search";
    import ArrowLeft from "@lucide/svelte/icons/arrow-left";
</script>

<svelte:head>
	<title>{import.meta.env.APP_NAME} | 404</title>
	<meta name="description" content="Dashboard overview and metrics" />
</svelte:head>

<main class="h-screen w-screen flex items-center justify-center p-4">
	<div class="absolute top-4 right-4 md:top-6 md:right-6">
		<Theme />
	</div>
	<Empty.Root>
		<Empty.Header>
			<Empty.Title class="text-5xl font-bold animate-pulse">404 - Not Found</Empty.Title>
			<Empty.Description>
				The page you're looking for doesn't exist. Try searching for what you need below.
			</Empty.Description>
		</Empty.Header>
		<Empty.Content>
			<button on:click={() => window.history.back()}>
				<Button><ArrowLeft /> Go back home</Button>
			</button>
			<!-- <InputGroup.Root class="sm:w-3/4">
				<InputGroup.Input placeholder="Try searching for pages..." />
				<InputGroup.Addon>
					<SearchIcon />
				</InputGroup.Addon>
				<InputGroup.Addon align="inline-end">
					<Kbd.Root>/</Kbd.Root>
				</InputGroup.Addon>
			</InputGroup.Root>
			<Empty.Description>
				Need help? <a href="#/">Contact support</a>
			</Empty.Description> -->
		</Empty.Content>
	</Empty.Root>
</main>
