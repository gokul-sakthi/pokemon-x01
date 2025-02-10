import Navbar from "@/components/header/Navbar";
import { Container } from "@chakra-ui/react";
import { Outlet } from "react-router";

const HomeLayout = () => {
  return (
    <Container>
      <Navbar />
      <main>
        <Outlet />
      </main>
    </Container>
  );
};

export default HomeLayout;
