<script lang="ts">
	import { goto } from "$app/navigation";
	import Button from "$lib/components/ui/button/button.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import Theme from "$lib/components/ui/theme/theme.svelte";
	import { toast } from "svelte-sonner";

//   let messages = $state()



	let loading = $state(false);
	let error = $state('');
	let success = $state('');

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		
		loading = true;
		error = '';
		success = '';

		const formDataObj = new FormData(e.currentTarget as HTMLFormElement);
		const email = formDataObj.get("email") as string;

		// console.log("Submitting email:", email);
		if (!email) {
			toast.error("Please enter a valid email address");
			loading = false;
			return;
		}

		try {
			const response = await fetch("/api/v1/auth/magic-link/request", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify({
					email
				})
			});

			const result = await response.json();
			if (!response.ok) {
				error = result.error || "Login failed";
				toast.error(result.error || "Login failed");
				// throw new Error(errorData || "Registration failed");
				return;
			}
			success = result.message;
			toast.success(result.message);
		} catch (err) {
			// error = err instanceof Error ? err.message : "An error occurred";
			toast.error("An error occurred");
			// console.error("Registration error:", err);
		} finally {
			loading = false;
		}
	}
 

</script>

<svelte:head>
	<title>Sign In - {import.meta.env.APP_NAME}</title>
	<meta name="description" content="Sign in to your account on {import.meta.env.APP_NAME}." />
</svelte:head>

<img
	src="/login.webp"
	alt="{import.meta.env.APP_NAME} Logo"
	class="h-screen w-screen fixed inset-0 z-0 dark:brightness-75"
/>

<section class="flex h-screen px-4 md:py-32 z-50">
	<div class="absolute top-4 right-4 md:top-6 md:right-6">
		<Theme />
	</div>
	<div class="m-auto h-fit w-full max-w-sm overflow-hidden rounded-xl border shadow z-50">
		<div class="bg-card/80 -m-px rounded-xl border p-4 pt-12 space-y-8">
			<div class="text-center space-y-6">
				<div class="w-fit mx-auto flex items-center gap-2">
					<img src="/logo.png" alt="{import.meta.env.APP_NAME} Logo" class="h-10 w-20" />
					<h2>
						{import.meta.env.APP_NAME}
					</h2>
					<!-- <Factory size="40" /> -->
				</div>
				<!-- <h1 class="mb-1 mt-4 text-xl font-semibold">Sign In</h1> -->
				<p class="text-sm">Welcome back! Sign in to continue</p>
			</div>

			<!-- if param email-success  -->
			{#if success}
				<div
					class="p-2 mb-4 text-sm rounded-lg bg-accent/80 text-green-700 dark:text-green-400"
					role="alert"
				>
					{success}
				</div>
			{:else if error}
				<div
					class="p-4 mb-4 text-sm text-red-600 rounded-lg bg-accent/80 dark:text-red-400"
					role="alert"
				>
					<strong class="font-medium">⚠️ {error}</strong>
				</div>
			{:else}
				<div class="px-2">
					<form class="flex gap-4 items-center" onsubmit={handleSubmit}>
						<Input type="email" name="email" placeholder="Email" required autocomplete="email" />
						<!-- bind:value={formData.email} -->
						<Button type="submit">Sign In</Button>
					</form>

					<!-- <div class="mt-6 grid grid-cols-2 gap-3"> -->
					<Button
						class="w-full my-4 py-6 flex items-center justify-center gap-4"
						type="button"
						variant="outline"
						onclick={async () => {
							loading = true;
							await goto('/api/v1/auth/google/login');
							loading = false;
						}}
						disabled={loading}
					>
						<!-- href="/api/v1/auth/google/login" -->
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="0.98em"
							height="1em"
							viewBox="0 0 256 262"
						>
							<path
								fill="#4285f4"
								d="M255.878 133.451c0-10.734-.871-18.567-2.756-26.69H130.55v48.448h71.947c-1.45 12.04-9.283 30.172-26.69 42.356l-.244 1.622l38.755 30.023l2.685.268c24.659-22.774 38.875-56.282 38.875-96.027"
							></path>
							<path
								fill="#34a853"
								d="M130.55 261.1c35.248 0 64.839-11.605 86.453-31.622l-41.196-31.913c-11.024 7.688-25.82 13.055-45.257 13.055c-34.523 0-63.824-22.773-74.269-54.25l-1.531.13l-40.298 31.187l-.527 1.465C35.393 231.798 79.49 261.1 130.55 261.1"
							></path>
							<path
								fill="#fbbc05"
								d="M56.281 156.37c-2.756-8.123-4.351-16.827-4.351-25.82c0-8.994 1.595-17.697 4.206-25.82l-.073-1.73L15.26 71.312l-1.335.635C5.077 89.644 0 109.517 0 130.55s5.077 40.905 13.925 58.602z"
							></path>
							<path
								fill="#eb4335"
								d="M130.55 50.479c24.514 0 41.05 10.589 50.479 19.438l36.844-35.974C195.245 12.91 165.798 0 130.55 0C79.49 0 35.393 29.301 13.925 71.947l42.211 32.783c10.59-31.477 39.891-54.251 74.414-54.251"
							></path>
						</svg>
						<span>Login with Google</span>
					</Button>
				</div>
			{/if}
			<!-- <Button type="button" variant="outline" disabled>
					<svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 256 256">
						<path fill="#f1511b" d="M121.666 121.666H0V0h121.666z"></path>
						<path fill="#80cc28" d="M256 121.666H134.335V0H256z"></path>
						<path fill="#00adef" d="M121.663 256.002H0V134.336h121.663z"></path>
						<path fill="#fbbc09" d="M256 256.002H134.335V134.336H256z"></path>
					</svg>
					<span>Microsoft</span>
				</Button> -->
			<!-- </div> -->

			<!-- <hr class="my-4 border-dashed" /> -->

			<!-- <div class="space-y-6">
				<div class="space-y-2">
					<Label for="email" class="block text-sm">Email</Label>
					<Input
						type="email"
						required
						autocomplete="email"
						bind:value={formData.email}
						class="input sz-md variant-mixed"
					/>
				</div>

				<div class="space-y-0.5">
					<div class="flex items-center justify-between">
						<Label for="pwd" class="text-title text-sm">Password</Label>
						<Button
							variant="link"
							href="#"
							size="sm"
							class="link intent-info variant-ghost text-sm"
						>
							Forgot your Password ?
						</Button>
					</div>
					<Input
						type="password"
						required
						name="pwd"
						id="pwd"
						bind:value={formData.pwd}
						class="input sz-md variant-mixed"
					/>
				</div>

				<Button class="w-full" type="submit">Sign In</Button>
			</div>
		</div>
		<div class="p-3">
			<p class="text-accent-foreground text-center text-sm">
				Don't have an account ?
				<Button href="/sign-up" variant="link" class="px-2">Create account</Button>
			</p>
			-->
		</div>
	</div>
</section>
