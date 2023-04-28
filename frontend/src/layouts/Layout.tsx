import { MetaProvider, Title } from "@solidjs/meta";
import { A } from "@solidjs/router";
import { JSX } from "solid-js";
import DarkBtn from "../components/DarkBtn";

type Props = {
  title?: string; // 页面标题
  children?: JSX.Element; // 子组件
};

export default function Layout(props: Props) {
  return (
    <>
      {/* meta 元数据 */}
      <MetaProvider>
        <Title>{props.title || "Gone App"}</Title>
      </MetaProvider>

      <div class="p-4 mx-auto max-w-screen-sm">
        <div class="mt-6 flex gap-6 items-center justify-center">
          <A class="text-lg" href="/">首页</A>
          <A class="text-lg" href="/about">关于</A>
          <DarkBtn />
        </div>
        <div class="mx-4 my-8">
          {props.children}
        </div>
      </div>
    </>
  )
}