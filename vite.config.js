import { defineConfig } from 'vite';

export default defineConfig({
  root: "assets",
  publicDir: "assets/public",
  build: {
    emptyOutDir: true,
    outDir: "../public",
    manifest: true,
  }
});
