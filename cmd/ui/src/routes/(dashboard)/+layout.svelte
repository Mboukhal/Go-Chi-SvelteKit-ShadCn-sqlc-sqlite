<script lang="ts">
	import { onMount } from "svelte";

	import * as Avatar from '$lib/components/ui/avatar/index.js';
	import * as HoverCard from '$lib/components/ui/hover-card/index.js';
	import { Button } from "@/lib/components/ui/button";
	import ThemeButton from "@/lib/components/theme-button.svelte";
	import { sessionStore } from "@/lib/hooks/get-user-info.svelte";
	import { goto } from "$app/navigation";

	let loading = $state(true);

	onMount(async() => {

		sessionStore.init();
		// sleep for 40 seconds 
		// console.log("sessionStore:", sessionStore.get());
		// await new Promise((resolve) => setTimeout(resolve, 400000));

		if (!sessionStore.isAuthenticated()) {
			goto('/sign-in');
		}
		loading = false;
	});

</script>

{#if loading}
	<div class="flex gap-4 h-screen w-screen items-center justify-center animate-pulse">
		<img src="/logo.png" alt="logo" class="h-16 w-32" />
		<p class="text-5xl font-semibold">{import.meta.env.APP_NAME}</p>
	</div>
{:else}
	<header
		class="flex h-14 items-center gap-2 ease-linear justify-between border-b shadow-sm border-dashed px-4"
	>
		<div class="flex items-center gap-2">
			<img src="/logo.png" alt="Logo" class="h-8 w-16" />
			<h2>{import.meta.env.APP_NAME}</h2>
		</div>
		<div class="flex items-center gap-2">
			{#if sessionStore.get()}
				<HoverCard.Root openDelay={100} closeDelay={100}>
					<HoverCard.Trigger>
						<Avatar.Root class="cursor-pointer">
							<Avatar.Image class="h-8 w-8" src={sessionStore.get()?.picture} />
							<Avatar.Fallback class="h-8 w-8 border p-2"
								>{(sessionStore.get()?.name || '').charAt(0).toUpperCase() ?? ''}</Avatar.Fallback
							>
						</Avatar.Root>
					</HoverCard.Trigger>
					<HoverCard.Content class="w-80" side="bottom" align="end" sideOffset={-8}>
						<div class="flex justify-between space-x-4">
							<div class="space-y-1">
								<h4 class="text-sm font-semibold">{sessionStore.get()?.name}</h4>
								<p class="text-sm">{sessionStore.get()?.email}</p>
								<div class="flex items-center pt-4">
									<Button
										class=""
										onclick={async () => {
											sessionStore.clear();
											goto('/api/v1/auth/logout');
											// window.location.href = '/';
										}}
									>
										Sign Out
									</Button>
								</div>
							</div>
							<div class="flex flex-col items-end justify-between space-y-2">
								<ThemeButton />
								<Button variant="ghost" href="/settings">Settings</Button>
							</div>
						</div>
					</HoverCard.Content>
				</HoverCard.Root>
			{/if}
		</div>
	</header>
	<!-- svelte-ignore slot_element_deprecated -->
	<div class="flex flex-1 flex-col gap-4 p-4">
		<slot />
	</div>
{/if}
