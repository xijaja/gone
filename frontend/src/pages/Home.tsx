import { A } from "@solidjs/router";

export default function HomePage() {
  return (
    <div>
      <h1>Home</h1>
      <p>This is the Home component.</p>

      <div class="mt-6 flex gap-2">
        <A href="/">首页</A>
        <A href="/about">关于</A>
      </div>
    </div>
  );
}