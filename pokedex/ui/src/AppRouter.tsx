import { Route, Routes } from "react-router";
import HomeLayout from "./layout/HomeLayout";
import MainPage from "./pages/MainPage";

const AppRouter = () => (
  <Routes>
    <Route path="/" element={<MainPage />}></Route>
    <Route path="/home" element={<HomeLayout />} />
  </Routes>
);

export default AppRouter;
