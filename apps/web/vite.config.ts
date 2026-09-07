import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// The build is embedded into the Go binary and served from the API origin, so the
// emitted asset URLs are relative. index.html carries <base href="/"> so those
// relative URLs still resolve from the root on deep client routes like /p/123.
export default defineConfig({
  base: "./",
  plugins: [react()],
  build: {
    outDir: "dist",
    emptyOutDir: true,
    sourcemap: false,
  },
  server: {
    port: 5173,
    proxy: {
      "/api": {
        target: "http://localhost:8787",
        changeOrigin: false,
      },
    },
  },
});
