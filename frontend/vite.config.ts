import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig, searchForWorkspaceRoot } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		fs: {
			// allow: [
			// 	searchForWorkspaceRoot(process.cwd()),
			// 	"./packages",
			// ],
			strict: false, // TODO: remove this after troubleshooting why forbidden
		},
	},
});
