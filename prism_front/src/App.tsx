import styled from "styled-components";
import { Route, Routes, BrowserRouter as Router } from "react-router-dom";

import Header from "./header/header";
import Main from "./main/main";

function App() {
  return (
    <Box>
      <Routes>
        <Route path="/*" element={<MainComponent />} />
        <Route path="/admin/*" element={<AdminComponent />} />
      </Routes>
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
  height: auto;
  border: 1px solid lightgray;
  border-top: none;
  border-bottom: none;
`;

const AdminHeader = () => {
  return <div>Admin Header</div>;
};

const AdminBody = () => {
  return <div>Admin Body</div>;
};

const AdminComponent = () => {
  return (
    <>
      <HeaderArea>
        <AdminHeader />
      </HeaderArea>
      <BodyArea>
        <AdminBody />
      </BodyArea>
    </>
  );
};

const MainComponent = () => {
  return (
    <>
      <HeaderArea>
        <Header />
      </HeaderArea>
      <BodyArea>
        <Main />
      </BodyArea>
    </>
  );
};

export default App;
