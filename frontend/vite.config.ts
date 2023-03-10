import { defineConfig } from "vite";
import solidPlugin from "vite-plugin-solid";
import autoImport from "unplugin-auto-import/vite";

export default defineConfig({
  plugins: [
    solidPlugin(),
    autoImport({
      dts: "./src/auto-import.d.ts",
      imports: ["solid-js", "@solidjs/router"],
    }),
  ],
  server: {
    port: 3000,
  },
  build: {
    target: "esnext",
  },
});
