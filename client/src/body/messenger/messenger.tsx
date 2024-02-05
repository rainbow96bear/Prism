import { Route, Routes } from "react-router-dom";
import styled from "styled-components";

const Messenger = () => {
  return (
    <MessengerContainer>
      <Routes>
        {/* <Route path={"/*"} element={<Root></Root>}></Route>
        <Route path={"/home"} element={<Home></Home>}></Route>
        <Route path={"/project"} element={<Project></Project>}></Route>
        <Route path={"/profile/:id"} element={<Profile />}></Route>
        <Route path={"/profile/update/userinfo"} element={<UserInfo />}></Route>
        <Route path={"/profile/update/techlist"} element={<TechList />}></Route> */}
      </Routes>
    </MessengerContainer>
  );
};

export default Messenger;

const MessengerContainer = styled.div`
  height: 100%;
`;
