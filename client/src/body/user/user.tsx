import { Route, Routes } from "react-router-dom";
import styled from "styled-components";

import Profile from "./Components/profile/profile";
import Root from "./Components/root/root";
import Home from "./Components/home/home";
import Project from "./Components/project/project";
import UserInfo from "./Components/update/userInfo";
import TechList from "./Components/update/techList";

const User = () => {
  return (
    <MainContainer>
      <Routes>
        <Route path={"/*"} element={<Root></Root>}></Route>
        <Route path={"/home"} element={<Home></Home>}></Route>
        <Route path={"/project"} element={<Project></Project>}></Route>
        <Route path={"/profile/:id"} element={<Profile />}></Route>
        <Route path={"/profile/update/userinfo"} element={<UserInfo />}></Route>
        <Route path={"/profile/update/techlist"} element={<TechList />}></Route>
      </Routes>
    </MainContainer>
  );
};

export default User;

const MainContainer = styled.div`
  height: 100%;
`;
