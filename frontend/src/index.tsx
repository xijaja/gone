/* @refresh reload */
import './styles/tailwind.css';
import { render } from 'solid-js/web';
import { Router } from '@solidjs/router';
import App from './Router';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    '找不到根元素。您是否忘记将其添加到您的 index.html 中？或者可能是 id 属性拼写错误？',
  );
}

render(() => (
  <Router>
    <App />
  </Router>
), root!);
