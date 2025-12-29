<script lang="ts">
 import SunIcon from "@lucide/svelte/icons/sun";
 import MoonIcon from "@lucide/svelte/icons/moon";
 
 import { toggleMode } from "mode-watcher";
 import { Button } from "$lib/components/ui/button/index.js";
	import * as Tooltip from "$lib/components/ui/tooltip/index.js";
    

</script>

<Tooltip.Provider>
	<Tooltip.Root>
		<Tooltip.Trigger>
			<Button onclick={toggleMode} variant="outline" size="icon">
				<SunIcon
					class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all! dark:-rotate-90 dark:scale-0"
				/>
				<MoonIcon
					class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all! dark:rotate-0 dark:scale-100"
				/>
			</Button>
		</Tooltip.Trigger>
		<Tooltip.Content side="right" align="center">
			<span class="dark:hidden">Dark mode</span>
			<span class="hidden dark:block">Light mode</span>
		</Tooltip.Content>
	</Tooltip.Root>
</Tooltip.Provider>
