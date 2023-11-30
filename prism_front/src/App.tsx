import styled from "styled-components";

import Header from "./header/header";

function App() {
  return (
    <Box>
      <Header></Header>
    </Box>
  );
}

export default App;

const Box = styled.div`
  display: flex;
  flex-direction: column;
`;
