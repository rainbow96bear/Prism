import { Routes, Route } from "react-router-dom";
import styled from "styled-components";

import Tech from "./tech/tech";

const Setting = () => {
  return (
    <>
      <Box>
        <Routes>
          <Route path="/tech" element={<Tech></Tech>}></Route>
        </Routes>
      </Box>
    </>
  );
};

export default Setting;

const Box = styled.div``;
