import styled from "styled-components";

import BeforeLogin from "./afterLogin";
import AfterLogin from "./beforeLogin";

const FuncBar: React.FC = () => {
  return (
    <Box>
      {/* <BeforeLogin /> */}
      <AfterLogin />
    </Box>
  );
};

export default FuncBar;

const Box = styled.div`
  display: flex;
  justify-content: right;
  height: 100%;
  align-items: center;
`;
