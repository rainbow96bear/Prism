import styled from "styled-components";

const Loading = () => {
  return <LoadingBox>Loading...</LoadingBox>;
};

export default Loading;

const LoadingBox = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 3rem;
  font-weight: bold;
  width: 100%;
  height: 100%;
  position: absolute;
`;
