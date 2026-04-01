// vite.config.js
import { sveltekit } from "file:///workspaces/Docker-AWS/daws/frontend/node_modules/.pnpm/@sveltejs+kit@2.55.0_@sveltejs+vite-plugin-svelte@4.0.4_svelte@5.55.0_vite@5.4.21_sass@_6bb46836c5eff400447967f31a896f84/node_modules/@sveltejs/kit/src/exports/vite/index.js";
import { defineConfig } from "file:///workspaces/Docker-AWS/daws/frontend/node_modules/.pnpm/vite@5.4.21_sass@1.98.0/node_modules/vite/dist/node/index.js";
var vite_config_default = defineConfig({
  plugins: [sveltekit()],
  server: {
    host: "0.0.0.0",
    port: 5173,
    strictPort: true,
    // Prevents Vite from jumping to 5174 if 5173 is "busy"
    hmr: {
      protocol: "ws",
      host: "localhost",
      port: 5173
    },
    watch: {
      usePolling: true,
      interval: 100
    },
    fs: {
      allow: ["./src/lib/wailsjs", "./src/lib", ".."]
    }
  },
  css: {
    preprocessorOptions: {
      api: "modern-compiler",
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
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvd29ya3NwYWNlcy9Eb2NrZXItQVdTL2Rhd3MvZnJvbnRlbmRcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIi93b3Jrc3BhY2VzL0RvY2tlci1BV1MvZGF3cy9mcm9udGVuZC92aXRlLmNvbmZpZy5qc1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vd29ya3NwYWNlcy9Eb2NrZXItQVdTL2Rhd3MvZnJvbnRlbmQvdml0ZS5jb25maWcuanNcIjtpbXBvcnQgeyBzdmVsdGVraXQgfSBmcm9tICdAc3ZlbHRlanMva2l0L3ZpdGUnO1xuaW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSc7XG5cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XG5cdHBsdWdpbnM6IFtzdmVsdGVraXQoKV0sXG5cdHNlcnZlcjoge1xuICAgICAgICBob3N0OiAnMC4wLjAuMCcsXG4gICAgICAgIHBvcnQ6IDUxNzMsXG4gICAgICAgIHN0cmljdFBvcnQ6IHRydWUsIC8vIFByZXZlbnRzIFZpdGUgZnJvbSBqdW1waW5nIHRvIDUxNzQgaWYgNTE3MyBpcyBcImJ1c3lcIlxuICAgICAgICBobXI6IHtcbiAgICAgICAgICAgIHByb3RvY29sOiAnd3MnLFxuICAgICAgICAgICAgaG9zdDogJ2xvY2FsaG9zdCcsXG4gICAgICAgICAgICBwb3J0OiA1MTczXG4gICAgICAgIH0sXG4gICAgICAgIHdhdGNoOiB7XG4gICAgICAgICAgICB1c2VQb2xsaW5nOiB0cnVlLFxuICAgICAgICAgICAgaW50ZXJ2YWw6IDEwMFxuICAgICAgICB9LFxuICAgICAgICBmczoge1xuICAgICAgICAgICAgYWxsb3c6IFsnLi9zcmMvbGliL3dhaWxzanMnLCAnLi9zcmMvbGliJywgJy4uJ10gXG4gICAgICAgIH1cblx0fSxcbiAgICBjc3M6IHtcbiAgICAgICAgcHJlcHJvY2Vzc29yT3B0aW9uczoge1xuICAgICAgICAgICAgYXBpOiAnbW9kZXJuLWNvbXBpbGVyJyxcbiAgICAgICAgICAgIHNjc3M6IHtcbiAgICAgICAgICAgICAgICBhZGRpdGlvbmFsRGF0YTogYFxuICAgICAgICAgICAgICAgICAgICBAaW1wb3J0IFwiJGxpYi9zdHlsZXMvX192YXJpYWJsZXMuc2Nzc1wiO1xuICAgICAgICAgICAgICAgICAgICBAaW1wb3J0IFwiJGxpYi9zdHlsZXMvX19jb2xvcnMuc2Nzc1wiO1xuICAgICAgICAgICAgICAgICAgICBAaW1wb3J0IFwiJGxpYi9zdHlsZXMvX19jb21tb24uc2Nzc1wiO1xuICAgICAgICAgICAgICAgIGBcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgIH1cbn0pOyJdLAogICJtYXBwaW5ncyI6ICI7QUFBOFIsU0FBUyxpQkFBaUI7QUFDeFQsU0FBUyxvQkFBb0I7QUFFN0IsSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDM0IsU0FBUyxDQUFDLFVBQVUsQ0FBQztBQUFBLEVBQ3JCLFFBQVE7QUFBQSxJQUNELE1BQU07QUFBQSxJQUNOLE1BQU07QUFBQSxJQUNOLFlBQVk7QUFBQTtBQUFBLElBQ1osS0FBSztBQUFBLE1BQ0QsVUFBVTtBQUFBLE1BQ1YsTUFBTTtBQUFBLE1BQ04sTUFBTTtBQUFBLElBQ1Y7QUFBQSxJQUNBLE9BQU87QUFBQSxNQUNILFlBQVk7QUFBQSxNQUNaLFVBQVU7QUFBQSxJQUNkO0FBQUEsSUFDQSxJQUFJO0FBQUEsTUFDQSxPQUFPLENBQUMscUJBQXFCLGFBQWEsSUFBSTtBQUFBLElBQ2xEO0FBQUEsRUFDUDtBQUFBLEVBQ0csS0FBSztBQUFBLElBQ0QscUJBQXFCO0FBQUEsTUFDakIsS0FBSztBQUFBLE1BQ0wsTUFBTTtBQUFBLFFBQ0YsZ0JBQWdCO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxNQUtwQjtBQUFBLElBQ0o7QUFBQSxFQUNKO0FBQ0osQ0FBQzsiLAogICJuYW1lcyI6IFtdCn0K
