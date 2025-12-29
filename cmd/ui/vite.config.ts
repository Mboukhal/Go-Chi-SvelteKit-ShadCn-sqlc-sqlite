import devtoolsJson from 'vite-plugin-devtools-json';
import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	// set up env file path and prefix
	envDir: '../',
	envPrefix: ['APP_'],
	server: {
		port: 1337,
		proxy: {
			'/api': 'http://localhost:3000'
		}
	},
	plugins: [tailwindcss(), sveltekit(), devtoolsJson()]
});
