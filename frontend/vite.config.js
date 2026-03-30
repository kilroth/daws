import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
        host: '0.0.0.0',
        port: 5173,
        strictPort: true, // Prevents Vite from jumping to 5174 if 5173 is "busy"
        hmr: {
            protocol: 'ws',
            host: 'localhost',
            port: 5173
        },
        watch: {
            usePolling: true,
            interval: 100
        },
        fs: {
            allow: ['./src/lib/wailsjs', './src/lib', '..'] 
        }
	},
    css: {
        preprocessorOptions: {
            api: 'modern-compiler',
            scss: {
                additionalData: `
                    @import "$lib/styles/__variables.scss";
                    @import "$lib/styles/__colors.scss";
                    @import "$lib/styles/__common.scss";
                `
            }
        }
    }
});