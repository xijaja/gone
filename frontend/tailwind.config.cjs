/** @type {import('tailwindcss').Config} */
module.exports = {
  // class 表示允许使用 class="dark" 手动切换暗黑模式，media 表示根据系统切换
  // darkMode: ["class", '[data-mode="forest"]'],
  darkMode: "media",
  content: ["./src/**/*.{ts,tsx}", "./src/*.{ts,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  // 配置 daisyui 插件主题
  daisyui: {
    styled: true, // 是否启用 daisyui 的样式
    themes: ["light", "dark"], // 更改它可以使用其他主题
    base: true, //  是否启用基本样式
    utils: true, // 是否启用工具类
    logs: false, // 是否启用日志 (在控制台中显示 daisyui 的日志)
    rtl: false, // 是否启用 rtl 模式 (从右到左)
    prefix: "", // 更改前缀
    darkTheme: "dark", // 指定暗色主题
  },
};
