import { MetaProvider, Meta } from "@solidjs/meta";
import { BsSunFill, BsMoonFill } from "solid-icons/bs";
import { themeChange } from "theme-change";

export default function DarkBtn() {
  // 获取当前主题并判断是否为暗色主题
  const themeIsDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const [dark, setDark] = createSignal(themeIsDark);
  onMount(async () => themeChange(dark()));

  return (
    <>
      {/* meta 元数据 */}
      <MetaProvider>
        <Show
          when={dark()}
          fallback={<Meta name="theme-color" content="#ffffff" />}
        >
          <Meta name="theme-color" content="#000000" />
        </Show>
      </MetaProvider>


      <button
        class="text-xl"
        data-toggle-theme="light,dark" // 注意：逗号前后不能有空格，否则将无效；并且主题需要在 daisyui 主题插件中有定义。
        data-act-class="ACTIVE-CLASS"
        onclick={() => setDark(!dark())}
      >
        <Show when={dark()} fallback={<BsSunFill />} keyed>
          <BsMoonFill />
        </Show>
      </button>
    </>
  )
}