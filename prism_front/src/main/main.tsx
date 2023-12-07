import { Route, Routes } from "react-router-dom";
import styled from "styled-components";

import Profile from "./Components/profile/profile";
import Update from "./Components/profile/update";

const Main = () => {
  return (
    <MainContainer>
      <Routes>
        <Route path={"/profile/:id"} element={<Profile />}></Route>
        <Route path={"/profile/update/id"} element={<Update />}></Route>
      </Routes>
    </MainContainer>
  );
};

export default Main;

const MainContainer = styled.div`
  height: 100%;
`;
