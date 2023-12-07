import styled from "styled-components";

import Header from "./header/header";
import Main from "./main/main";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <Box>
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
