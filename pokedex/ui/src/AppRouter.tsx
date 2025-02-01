import { Route, Routes } from "react-router";
import HomeLayout from "./layout/HomeLayout";

const AppRouter = () => (
  <Routes>
    <Route index element={<HomeLayout />} />
  </Routes>
);

export default AppRouter;
