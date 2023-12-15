import styled from "styled-components";

const Root = () => {
  return <Box>root입니다.</Box>;
};

export default Root;

const Box = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  position: absolute;
`;
