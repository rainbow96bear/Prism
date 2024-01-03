import { Route, Routes } from "react-router-dom";
import styled from "styled-components";

import Profile from "./Components/profile/profile";
import Update from "./Components/profile/update";
import Root from "./Components/root/root";
import Home from "./Components/home/home";
import Project from "./Components/project/project";

const Main = () => {
  return (
    <MainContainer>
      <Routes>
        <Route path={"/*"} element={<Root></Root>}></Route>
        <Route path={"/home"} element={<Home></Home>}></Route>
        <Route path={"/project"} element={<Project></Project>}></Route>
        <Route path={"/profile/:id"} element={<Profile />}></Route>
        <Route path={"/profile/update"} element={<Update />}></Route>
      </Routes>
    </MainContainer>
  );
};

export default Main;

const MainContainer = styled.div`
  height: 100%;
`;
