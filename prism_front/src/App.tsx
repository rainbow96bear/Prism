import styled from "styled-components";

import Header from "./header/header";
import Main from "./main/main";
import Login from "./login/login";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <Box>
      <Routes>
        <Route path="/login" element={<Login />} />
      </Routes>
      <HeaderArea>
        <Header></Header>
      </HeaderArea>
      <BodyArea>
        <Main></Main>
      </BodyArea>
    </Box>
  );
}

export default App;

const Box = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;
const HeaderArea = styled.div`
  width: 100%;
  height: 55px;
  border-bottom: 2px solid gray;
`;

const BodyArea = styled.div`
  width: 100%;
  max-width: 70%;
  height: auto;
  background-color: blue;
`;
