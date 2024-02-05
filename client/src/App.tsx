import styled from "styled-components";
import { Route, Routes, BrowserRouter as Router } from "react-router-dom";
import Modal from "react-modal";

import Header from "./header/user/header";
import User from "./body/user/user";
import AdminMain from "./body/admin/main/main";
import Messenger from "./body/messenger/messenger";
import AdminHeader from "./header/admin/header";

Modal.setAppElement("#root");

function App() {
  return (
    <Box>
      <HeaderArea>
        <Routes>
          <Route path="/*" element={<Header />} />
          <Route path="/admin/*" element={<AdminHeader />} />
        </Routes>
      </HeaderArea>
      <BodyArea>
        <Routes>
          <Route path="/*" element={<User />} />
          <Route path="/messenger" element={<Messenger />} />
          <Route path="/admin/*" element={<AdminMain />} />
        </Routes>
      </BodyArea>
    </Box>
  );
}

const Box = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const HeaderArea = styled.div`
  display: flex;
  justify-content: center;
  width: 100%;
  height: 70px;
  border-bottom: 2px solid gray;
`;

const BodyArea = styled.div`
  width: 100%;
  max-width: 70%;
  min-height: calc(100vh - 70px);
  border: 1px solid lightgray;
  border-top: none;
  border-bottom: none;
  position: relative;
`;

export default App;
