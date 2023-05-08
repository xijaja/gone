import HomePage from "./pages/Home";
import AboutPage from "./pages/About";
import NotFoundPage from "./pages/404";

export default function AppRouter() {
  return (
    <>
      <Routes>
        <Route path="/" component={HomePage} />
        <Route path="/about" component={AboutPage} />
        <Route path="*" component={NotFoundPage} />
      </Routes>
    </>
  );
};
